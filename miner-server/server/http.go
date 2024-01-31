package server

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"log"
	http2 "net/http"
	"uminer/common/middleware/logging"
	"uminer/miner-server/api/chipApi/HTTP"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service"
)

// router
func deployRouters(s *http.Server, service *service.Service) {
	s.HandleFunc("/chipApi.ChipService/ListAllChips", func(w http.ResponseWriter, r *http.Request) {
		// 确认HTTP方法是GET
		if r.Method != http2.MethodGet {
			http2.Error(w, "Method Not Allowed", http2.StatusMethodNotAllowed)
			return
		}
		// 解析参数
		query := r.URL.Query()
		req := &HTTP.ChipsRequest{
			Url:       query.Get("url"),
			SerialNum: query.Get("serialNum"),
			BusId:     query.Get("busId"),
		}
		resp, err := service.ChipService.ListAllChipsHTTP(r.Context(), req)
		if err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http2.Error(w, err.Error(), http2.StatusInternalServerError)
		}
	})
	s.HandleFunc("/chipApi.ChipService/StartChipCPU", func(w http.ResponseWriter, r *http.Request) {

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
