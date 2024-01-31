package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	grpc2 "google.golang.org/grpc"
	"log"
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
				log.Println("logging: rpc call")
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
