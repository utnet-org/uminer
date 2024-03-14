package main

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
	"google.golang.org/grpc"
	"net/http"
	"os"
	"time"
	"uminer/common/graceful"
	"uminer/common/log"
	chainApi "uminer/miner-server/api/chainApi/rpc"
	"uminer/miner-server/cmd"
	"uminer/miner-server/server"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service"
	"uminer/miner-server/util"
)

type ServerAddr struct {
	HttpServer string
	GrpcServer string
}

func main() {

	// worker server address
	workerAddresses := []ServerAddr{
		{HttpServer: "192.168.10.19:8001", GrpcServer: "192.168.10.19:9001"},
		{HttpServer: "192.168.10.47:8001", GrpcServer: "192.168.10.47:9001"},
	}

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

	// 创建 Bootstrap 配置对象
	bootstrap := &serverConf.Bootstrap{
		App: &serverConf.App{
			// 设置 App 相关字段
		},
		Server: &serverConf.Server{
			Http: httpServer,
			Grpc: grpcServer,
		},
		Data: &serverConf.Data{
			// 设置 Data 相关字段
		},
		Storage: []byte("my_storage_data"), // 设置 Storage 字段
	}

	app, close, err := initApp(context.Background(), bootstrap, log.DefaultLogger, workerAddresses)
	if err != nil {
		panic(err)
	}
	defer close()

	// start and wait for stop signal
	fmt.Println("Miner started successfully.")
	if err := app.Run(); err != nil {
		panic(err)
	}

	// 协程优雅退出
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		graceful.Shutdown(ctx)
	}()
}

// initApp init kratos application.
func initApp(ctx context.Context, bc *serverConf.Bootstrap, logger log.Logger, workerAddresses []ServerAddr) (*kratos.App, func(), error) {
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
	// connect worker grpc
	//var workerGrpcConnArr []rpc.ChipServiceServer
	//for _, addr := range workerAddresses {
	//	workerHServer := &serverConf.Server_HTTP{
	//		Network: "tcp",
	//		Addr:    addr.HttpServer,
	//		Timeout: &duration.Duration{Seconds: 60},
	//	}
	//	workerGServer := &serverConf.Server_GRPC{
	//		Network: "tcp",
	//		Addr:    addr.GrpcServer,
	//		Timeout: &duration.Duration{Seconds: 60},
	//	}
	//	bs := &serverConf.Bootstrap{
	//		App: &serverConf.App{},
	//		Server: &serverConf.Server{
	//			Http: workerHServer,
	//			Grpc: workerGServer,
	//		},
	//		Data:    &serverConf.Data{},
	//		Storage: []byte("my_storage_data"), // 设置 Storage 字段
	//	}
	//	client := types.NewChipService(bs, logger, nil)
	//	workerGrpcConnArr = append(workerGrpcConnArr, client)
	//
	//}

	httpServer := server.NewHTTPServer(bc.Server, newService)

	// listen to the nodes for bursting a block
	go listenBurst(ctx, bc.Server.Grpc.Addr)

	app := newApp(ctx, logger, httpServer, grpcServer)

	return app, nil, nil
}

func newApp(ctx context.Context, logger log.Logger, hs *minerHttp.Server, gs *minerGprc.Server) *kratos.App {
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

// listenBurst listen to the node, preparing for burst
func listenBurst(ctx context.Context, address string) {

	// dial local grpc for challenge computation
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("fail to dial miner RPC ", address)
		return
	}
	chaincli := chainApi.NewChainServiceClient(conn)
	// get keys
	_, pubErr := os.Stat("public.pem")
	if pubErr != nil {
		fmt.Println("no pubkey file is found")
		return
	}
	keys, err := chaincli.GetMinerKeys(ctx, &chainApi.GetMinerKeysRequest{PrivateKey: ""})
	if err != nil {
		fmt.Println("fail to get miner address RPC ", err)
		return
	}
	// get all chips and workers address ready
	list, err := chaincli.GetMinerChipsList(ctx, &chainApi.GetMinerChipsListRequest{AccountId: keys.Address})
	if err != nil {
		fmt.Println("fail to get miner chip lists RPC ", err)
		return
	}
	fmt.Println("chip list:", list)
	workers := make([]string, 0)
	for _, item := range cmd.WorkerLists {
		workers = append(workers, item)
	}

	// start loop
	request := &chainApi.ChallengeComputationRequest{
		ChallengeKey: "ed25519:" + keys.PubKey,
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
		ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
		defer cancel()
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
		if err != nil {
			fmt.Println("Error connecting to query latest blockHash RPC: ", err.Error())
			continue
		}
		r.Header.Add("Content-Type", "application/json; charset=utf-8")
		r.Header.Add("accept-encoding", "gzip,deflate")

		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(r)
		if err != nil {
			fmt.Println("fail to get query latest blockHeight RPC response: ", err.Error())
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

		// get the mining provider
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
		r, err = http.NewRequestWithContext(ctx, http.MethodPost, cmd.NodeURL, bytes.NewReader(jsonStr))
		if err != nil {
			fmt.Println("Error connecting to query miner provider RPC: ", err.Error())
			continue
		}
		r.Header.Add("Content-Type", "application/json; charset=utf-8")
		r.Header.Add("accept-encoding", "gzip,deflate")

		cli := &http.Client{Timeout: 5 * time.Second}
		resp, err = cli.Do(r)
		if err != nil {
			fmt.Println("fail to get query miner provider RPC response: ", err.Error())
			continue
		}
		defer resp.Body.Close()
		gzipBytes = util.GzipApi(resp)
		res = gjson.Get(string(gzipBytes), "result").String()
		provider := gjson.Get(res, "provider_account").String()
		//fmt.Println(string(gzipBytes))
		// check if the provider candidate is yourself
		if provider != request.ChallengeKey {
			fmt.Println("chosen : ", provider, ", my account is", request.ChallengeKey)
			continue
		}

		// wait for real burst
		fmt.Println("block candidate is selected!")
		waitingBurstLoop(ctx, chaincli, request, time.Now().Unix()+10)

	}

}
func waitingBurstLoop(ctx context.Context, cli chainApi.ChainServiceClient, request *chainApi.ChallengeComputationRequest, burstTime int64) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
		// get the right burst timing
		if time.Now().Unix() >= burstTime {
			response, err := cli.ChallengeComputation(ctx, request)
			if err != nil {
				fmt.Println("Error calling ChallengeComputation:", err)
			} else {
				fmt.Println("ChallengeComputation response: ", response)
			}
			fmt.Println("block is burst and broadcast !")
			return
		}
	}
}
