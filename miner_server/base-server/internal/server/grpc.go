package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"uminer/common/middleware/ctxcopy"
	"uminer/common/middleware/logging"
	"uminer/common/middleware/validate"
	//api "uminer/miner_server/base-server/api/v1"
	"uminer/miner_server/base-server/internal/conf"
	"uminer/miner_server/base-server/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, s *service.Service) *grpc.Server {
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
	//api.RegisterUserServiceServer(gs, s.UserService)
	return gs
}