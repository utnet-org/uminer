package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"uminer/common/middleware/ctxcopy"
	"uminer/common/middleware/logging"
	"uminer/common/middleware/validate"
	chainApi "uminer/miner-server/api/chainApi/rpc"
	chipApi "uminer/miner-server/api/chipApi/rpc"
	containerApi "uminer/miner-server/api/containerApi"
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
	chipApi.RegisterChipServiceServer(gs, s.ChipService)
	return gs
}
