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

// http router for different handlers func
func deployRouters(s *http.Server, service *service.Service) {
	// JSON-RPC handler function
	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HandleJSONRPCRequest(service, w, r)
	})
	// miner login to get address and container cloud server token
	s.HandleFunc("/chainService/Login", func(w http.ResponseWriter, r *http.Request) {
		service.MinerLoginServiceH.LoginHandler(w, r)
	})
	// get container cloud userid by the token
	s.HandleFunc("/chainService/GetMinerInfo", func(w http.ResponseWriter, r *http.Request) {
		service.MinerLoginServiceH.GetMinerInfoHandler(w, r)
	})
	// list all workers' bm-chip information
	s.HandleFunc("/chipService/ListAllChips", func(w http.ResponseWriter, r *http.Request) {
		service.MinerStatusServiceH.ListAllChipsHTTPHandler(w, r)
	})
	// activate the chip manually before using it as cryptological tool
	s.HandleFunc("/chipService/StartChipCPU", func(w http.ResponseWriter, r *http.Request) {
		service.MinerStatusServiceH.StartChipCPUHandler(w, r)
	})
	// update the latest status and information of the utility node
	s.HandleFunc("/chainService/GetNodesStatus", func(w http.ResponseWriter, r *http.Request) {
		service.MinerStatusServiceH.GetNodesStatusHandler(w, r)
	})
	// view miner account information by the rpc method on utility node
	s.HandleFunc("/chainService/ViewAccount", func(w http.ResponseWriter, r *http.Request) {
		service.MinerStatusServiceH.ViewAccountHandler(w, r)
	})
	// list all chip rental orders related to the miner with their details
	s.HandleFunc("/chainService/GetRentalOrderList", func(w http.ResponseWriter, r *http.Request) {
		service.MinerContainerServiceH.GetRentalOrderListHandler(w, r)
	})

}

// NewHTTPServer initialize an HTTP server.
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

// corsFilter kratos framework cross-origin middleware
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
