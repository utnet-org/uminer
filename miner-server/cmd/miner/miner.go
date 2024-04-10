package miner

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	minerGprc "github.com/go-kratos/kratos/v2/transport/grpc"
	minerHttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"net/http"
	"path/filepath"
	"time"
	"uminer/common/graceful"
	"uminer/common/log"
	chainApi "uminer/miner-server/api/chainApi/rpc"
	"uminer/miner-server/cmd"
	"uminer/miner-server/cmd/utlog"
	"uminer/miner-server/server"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service"
	"uminer/miner-server/util"
)

func StartMinerServer(c *cli.Context) error {

	cmd.MinerServerIP = c.String("serverip")
	if c.IsSet("node") {
		cmd.NodeURL = c.String("node")
	}
	cmd.WorkerLists = c.StringSlice("workerip")

	utlog.Mainlog.Info("Initializing utility miner")

	// activate miner server
	httpServer := &serverConf.Server_HTTP{
		Network: "tcp",
		Addr:    cmd.MinerServerIP + ":8001",
		Timeout: &duration.Duration{Seconds: 60},
	}

	grpcServer := &serverConf.Server_GRPC{
		Network: "tcp",
		Addr:    cmd.MinerServerIP + ":9001",
		Timeout: &duration.Duration{Seconds: 60},
	}

	// build Bootstrap configuration
	bootstrap := &serverConf.Bootstrap{
		App: &serverConf.App{
			// set App
		},
		Server: &serverConf.Server{
			// set Server
			Http: httpServer,
			Grpc: grpcServer,
		},
		Data: &serverConf.Data{
			// set Data
		},
		Storage: []byte("my_storage_data"), // 设置 Storage 字段
	}

	app, close, err := initMinerApp(context.Background(), bootstrap, log.DefaultLogger)
	if err != nil {
		panic(err)
	}
	defer close()

	// start and wait for stop signal
	utlog.Mainlog.Info("Miner started successfully")
	if err := app.Run(); err != nil {
		panic(err)
	}

	// exit gently
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		graceful.Shutdown(ctx)
	}()

	return nil
}

// initMinerApp init kratos application.
func initMinerApp(ctx context.Context, bc *serverConf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	//data, close, err := data.NewData(bc, logger)
	//if err != nil {
	//	return nil, nil, err
	//}
	newService, err := service.NewMinerService(ctx, bc, logger, nil)
	if err != nil {
		return nil, nil, err
	}

	// new miner grpc
	grpcServer := server.NewMinerGRPCServer(bc.Server, newService)

	httpServer := server.NewHTTPServer(bc.Server, newService)

	// listen to the nodes for being chosen to mine and report a block
	go listenMining(ctx, bc.Server.Grpc.Addr)

	app := newMiner(ctx, logger, httpServer, grpcServer)

	return app, nil, nil
}

func newMiner(ctx context.Context, logger log.Logger, hs *minerHttp.Server, gs *minerGprc.Server) *kratos.App {
	return kratos.New(
		kratos.Context(ctx),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

// listenMining listen to the node, preparing for burst
func listenMining(ctx context.Context, address string) {

	// dial local grpc for challenge computation
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		utlog.Mainlog.Error("fail to dial miner RPC")
		return
	}
	chaincli := chainApi.NewChainServiceClient(conn)
	// get miner keys
	file, fileErr := filepath.Glob("*.json")
	if fileErr != nil || len(file) == 0 {
		utlog.Mainlog.Error("no pubkey file is found")
		return
	}
	keys, err := chaincli.GetMinerAccountKeys(ctx, &chainApi.GetMinerAccountKeysRequest{Mnemonic: ""})
	if err != nil {
		utlog.Mainlog.Error("fail to get miner address RPC:", err.Error())
		return
	}
	// get all chips and workers address ready
	list, err := chaincli.GetMinerChipsList(ctx, &chainApi.GetMinerChipsListRequest{AccountId: keys.Address})
	if err != nil {
		utlog.Mainlog.Error("fail to get miner chip lists RPC:", err.Error())
		return
	}
	fmt.Println("my total chip power:", list.TotalPower)
	workers := make([]string, 0)
	for _, item := range cmd.WorkerLists {
		workers = append(workers, item)
	}

	// start listening loop
	request := &chainApi.ChallengeComputationRequest{
		ChallengeKey: keys.PubKey,
		Url:          workers,
		Message:      "utility",
		Chips:        list.Chips,
	}
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}

		/* listening to the node to be informed if being chosen as miner to burst the block */
		// get the latest hash
		jsonData := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      "dontcare",
			"method":  "status",
			"params":  make([]interface{}, 0),
		}
		jsonStr, _ := json.Marshal(jsonData)
		clientDeadline := time.Now().Add(time.Duration(4 * time.Second))
		ctx2, cancel := context.WithDeadline(context.Background(), clientDeadline)
		defer cancel()
		r, err := http.NewRequestWithContext(ctx2, http.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
		if err != nil {
			utlog.Mainlog.Error("Error connecting to query latest blockHash RPC:", err.Error())
			continue
		}
		r.Header.Add("Content-Type", "application/json; charset=utf-8")
		r.Header.Add("accept-encoding", "gzip,deflate")

		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(r)
		if err != nil {
			utlog.Mainlog.Error("fail to get query latest blockHeight RPC response:", err.Error())
			continue
		}
		defer resp.Body.Close()
		gzipBytes := util.GzipApi(resp)
		res := gjson.Get(string(gzipBytes), "result").String()
		sync := gjson.Get(res, "sync_info").String()
		latestBlockH := gjson.Get(sync, "latest_block_height").Int()
		if cmd.LatestBlockH == latestBlockH {
			continue
		}

		// get the next mining provider
		utlog.Mainlog.Info("querying the next block candidate...")
		cmd.LatestBlockH = latestBlockH
		jsonData = map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      "dontcare",
			"method":  "provider",
			"params": map[string]interface{}{
				"block_height": latestBlockH,
			},
		}
		jsonStr, _ = json.Marshal(jsonData)
		r, err = http.NewRequestWithContext(ctx2, http.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
		if err != nil {
			utlog.Mainlog.Error("Error connecting to query miner provider RPC:", err.Error())
			continue
		}
		r.Header.Add("Content-Type", "application/json; charset=utf-8")
		r.Header.Add("accept-encoding", "gzip,deflate")

		cli := &http.Client{Timeout: 5 * time.Second}
		resp, err = cli.Do(r)
		if err != nil {
			utlog.Mainlog.Error("fail to get query miner provider RPC response:", err.Error())
			continue
		}
		defer resp.Body.Close()
		gzipBytes = util.GzipApi(resp)
		res = gjson.Get(string(gzipBytes), "result").String()
		provider := gjson.Get(res, "provider_account").String()
		// check if the provider candidate is yourself
		if provider != request.ChallengeKey {
			utlog.Mainlog.Warn("chosen : ", provider, ", my account is", request.ChallengeKey)
			continue
		}

		// wait for the timing to be challenged and asked to sign chips for computation proof
		utlog.Mainlog.Info("you are selected as the next block candidate!")
		waitingChallengeLoop(ctx, chaincli, request, time.Now().Unix()+10)

	}

}
func waitingChallengeLoop(ctx context.Context, cli chainApi.ChainServiceClient, request *chainApi.ChallengeComputationRequest, challengeTime int64) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
		// get the right challenge timing
		if time.Now().Unix() >= challengeTime {
			response, err := cli.ChallengeComputation(ctx, request)
			if err != nil {
				utlog.Mainlog.Error("Error calling ChallengeComputation:", err.Error())
			} else {
				utlog.Mainlog.Info("ChallengeComputation response: ", response)
			}
			utlog.Mainlog.Info("block is burst and broadcast !")
			return
		}
	}
}
