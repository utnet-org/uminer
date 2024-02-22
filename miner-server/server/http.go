package server

import (
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"log"
	http2 "net/http"
	"uminer/common/middleware/logging"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service"
)

// http router for different methods
func deployRouters(s *http.Server, service *service.Service) {
	// miner login to get token by username and password
	s.HandleFunc("/chainService/Login", func(w http.ResponseWriter, r *http.Request) {
		service.MinerUIServiceH.LoginHandler(w, r)
	})
	// get userid by token
	s.HandleFunc("/chainService/GetMinerInfo", func(w http.ResponseWriter, r *http.Request) {
		service.MinerUIServiceH.GetMinerInfoHandler(w, r)
	})
	// list all workers' bm-chip information
	s.HandleFunc("/chipService/ListAllChips", func(w http.ResponseWriter, r *http.Request) {
		service.MinerUIServiceH.ListAllChipsHTTPHandler(w, r)
	})
	// activate the chip manually before using it as cryptological tool
	s.HandleFunc("/chipService/StartChipCPU", func(w http.ResponseWriter, r *http.Request) {
		service.MinerUIServiceH.StartChipCPUHandler(w, r)
	})
	// update the latest information of the node
	s.HandleFunc("/chainService/GetNodesStatus", func(w http.ResponseWriter, r *http.Request) {
		service.MinerUIServiceH.GetNodesStatusHandler(w, r)
	})
	// view account
	s.HandleFunc("/chainService/ViewAccount", func(w http.ResponseWriter, r *http.Request) {
		service.MinerUIServiceH.ViewAccountHandler(w, r)
	})
	// list all rental orders related to the miner with their details
	s.HandleFunc("/chainService/GetRentalOrderList", func(w http.ResponseWriter, r *http.Request) {
		service.MinerUIServiceH.GetRentalOrderListHandler(w, r)
	})
	// list all notebooks related to the miner with their details
	s.HandleFunc("/notebookService/GetNotebookList", func(w http.ResponseWriter, r *http.Request) {
		service.MinerUIServiceH.GetNotebookListHandler(w, r)
	})

}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *serverConf.Server, service *service.Service) *http.Server {
	var opts = []http.ServerOption{}

	middlewareChain :=
		middleware.Chain(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(),
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
	opts = append(opts, http.Middleware(middlewareChain))
	opts = append(opts, http.Filter(corsFilter, loggingFilter))
	srv := http.NewServer(opts...)
	deployRouters(srv, service)
	return srv
}

// kratos框架跨域中间件
func corsFilter(next http2.Handler) http2.Handler {
	return http2.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http2.MethodOptions {
			log.Println("cors:", r.Method, r.RequestURI)
			w.WriteHeader(http2.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func loggingFilter(next http2.Handler) http2.Handler {
	return http2.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("logging:", r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
