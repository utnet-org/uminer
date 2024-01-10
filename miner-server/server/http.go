package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"uminer/common/middleware/logging"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *serverConf.Server, service *service.Service) *http.Server {
	var opts = []http.ServerOption{}

	http.Middleware(
		middleware.Chain(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(),
		),
	)

	//http.WithTimeout(time.Minute *2)

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	return srv
}