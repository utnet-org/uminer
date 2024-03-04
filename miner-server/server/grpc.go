package server

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	grpc2 "google.golang.org/grpc"
	"io/ioutil"
	"log"
	http2 "net/http"
	"uminer/common/middleware/ctxcopy"
	"uminer/common/middleware/logging"
	"uminer/common/middleware/validate"
	chainApi "uminer/miner-server/api/chainApi/rpc"
	chipApi "uminer/miner-server/api/chipApi/rpc"
	"uminer/miner-server/api/containerApi"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service"
)

// NewMinerGRPCServer new a gRPC server.
func NewMinerGRPCServer(c *serverConf.Server, s *service.Service) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				ctxcopy.Server(),
				//status.Server(status.WithHandler(errors.ErrorEncode)),
				tracing.Server(),
				logging.Server(),
				validate.Server(),
				MiddlewareCors(),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	gs := grpc.NewServer(opts...)
	chainApi.RegisterChainServiceServer(gs, s.ChainService)
	containerApi.RegisterImageServiceServer(gs, s.ImageService)
	containerApi.RegisterNotebookServiceServer(gs, s.NotebookService)
	return gs
}

// NewWorkerGRPCServer new a gRPC server.
func NewWorkerGRPCServer(c *serverConf.Server, s *service.Service) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				ctxcopy.Server(),
				//status.Server(status.WithHandler(errors.ErrorEncode)),
				tracing.Server(),
				logging.Server(),
				validate.Server(),
				MiddlewareCors(),
			),
		),
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc2.UnaryServerInfo, handler grpc2.UnaryHandler) (resp interface{}, err error) {
			return handler(ctx, req)
		}),
		grpc.Options(grpc2.InitialConnWindowSize(0)),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	opts = append(opts, grpc.Middleware(MiddlewareCors()))
	gs := grpc.NewServer(opts...)
	chipApi.RegisterChipServiceServer(gs, s.ChipServiceG)
	return gs
}

// MiddlewareCors kratos框架跨域中间件
func MiddlewareCors() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if ts, ok := transport.FromServerContext(ctx); ok {
				log.Println("logging: rpc", ts.Operation())
				if ts.ReplyHeader() != nil {
					ts.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
					ts.ReplyHeader().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
					ts.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
					ts.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type,"+
						"X-Requested-With,Access-Control-Allow-Credentials,User-Agent,Content-Length,Authorization")
				}
			}
			return handler(ctx, req)
		}
	}
}

// HandleJSONRPCRequest json gprc
func HandleJSONRPCRequest(srv *service.Service, w http.ResponseWriter, r *http2.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusBadRequest)
		return
	}
	// JSON-RPCc request
	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		http2.Error(w, err.Error(), http2.StatusBadRequest)
		return
	}
	// get method name
	method, ok := jsonData["method"].(string)
	if !ok {
		http2.Error(w, "Method not found in request", http2.StatusBadRequest)
		return
	}
	params, ok := jsonData["params"].(map[string]interface{})
	if !ok {
		http2.Error(w, "Params not found in request", http2.StatusBadRequest)
		return
	}
	// method
	switch method {

	// to worker chips operation
	case "startchips":
		if !ok {
			http2.Error(w, "Params not found in request", http2.StatusBadRequest)
			return
		}
		devId, ok := params["dev_id"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		request := &chipApi.ChipsRequest{
			Url:       "",
			SerialNum: "",
			BusId:     "",
			DevId:     devId,
		}
		response, err := srv.ChipServiceG.StartChipCPU(r.Context(), request)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}
	case "burnchips":
		devId, ok := params["dev_id"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		request := &chipApi.ChipsRequest{
			Url:       "",
			SerialNum: "",
			BusId:     "",
			DevId:     devId,
		}
		response, err := srv.ChipServiceG.BurnChipEfuse(r.Context(), request)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}
	case "keysgen":
		devId, ok := params["dev_id"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		request := &chipApi.ChipsRequest{
			Url:       "",
			SerialNum: "",
			BusId:     "",
			DevId:     devId,
		}
		response, err := srv.ChipServiceG.GenerateChipKeyPairs(r.Context(), request)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}
	case "keysread":
		serial, ok := params["serial"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		busId, ok := params["bus_id"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		devId, ok := params["dev_id"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		request := &chipApi.ChipsRequest{
			Url:       "",
			SerialNum: serial,
			BusId:     busId,
			DevId:     devId,
		}
		response, err := srv.ChipServiceG.ObtainChipKeyPairs(r.Context(), request)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}

	// to chain nodes
	case "getminerkeys":
		accessKey, ok := params["access_key"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		request := &chainApi.GetMinerKeysRequest{AccessKeys: accessKey}
		response, err := srv.ChainService.GetMinerKeys(r.Context(), request)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}
	case "claimstake":
		if !ok {
			http2.Error(w, "Params not found in request", http2.StatusBadRequest)
			return
		}
		accountId, ok := params["account_id"].(string)
		if !ok {
			http2.Error(w, "Account ID not found in params", http2.StatusBadRequest)
			return
		}
		amount, ok := params["amount"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		nearPath, ok := params["near_path"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		keyPath, ok := params["key_path"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		request := &chainApi.ClaimStakeRequest{
			AccountId: accountId,
			Amount:    amount,
			NearPath:  nearPath,
			KeyPath:   keyPath,
		}
		response, err := srv.ChainService.ClaimStake(r.Context(), request)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}
	case "claimcomputation":
		if !ok {
			http2.Error(w, "Params not found in request", http2.StatusBadRequest)
			return
		}
		accountId, ok := params["account_id"].(string)
		if !ok {
			http2.Error(w, "Account ID not found in params", http2.StatusBadRequest)
			return
		}
		nearPath, ok := params["near_path"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		keyPath, ok := params["key_path"].(string)
		if !ok {
			http2.Error(w, "Access key not found in params", http2.StatusBadRequest)
			return
		}
		request := &chainApi.ClaimChipComputationRequest{
			AccountId: accountId,
			ChipPubK:  "",
			ChipP2K:   "",
			NearPath:  nearPath,
			KeyPath:   keyPath,
		}
		response, err := srv.ChainService.ClaimChipComputation(r.Context(), request)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}

	// container cloud server
	case "getNotebookList":
		if !ok {
			http2.Error(w, "Params not found in request", http2.StatusBadRequest)
			return
		}
		token, ok := params["token"].(string)
		if !ok {
			http2.Error(w, "token not found in params", http2.StatusBadRequest)
			return
		}
		notebookId, ok := params["notebookId"].(string)
		if !ok {
			http2.Error(w, "notebookId not found in params", http2.StatusBadRequest)
			return
		}
		request := &containerApi.QueryNotebookByConditionRequest{
			Token:     token,
			Id:        notebookId,
			PageSize:  10,
			PageIndex: 1,
		}
		response, err := srv.NotebookService.QueryNotebookByCondition(r.Context(), request)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}

	default:
		http2.Error(w, "Method not supported", http2.StatusMethodNotAllowed)
		return
	}
}
