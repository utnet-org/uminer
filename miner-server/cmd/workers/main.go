package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/golang/protobuf/ptypes/duration"
	"time"
	"uminer/common/graceful"

	"uminer/common/log"
	"uminer/miner-server/serverConf"

	"uminer/miner-server/server"
	"uminer/miner-server/service"

	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func main() {

	httpServer := &serverConf.Server_HTTP{
		Network: "tcp",
		Addr:    "0.0.0.0:8001",
		Timeout: &duration.Duration{Seconds: 60},
	}

	// 创建 Server_GRPC 对象并设置相关字段
	grpcServer := &serverConf.Server_GRPC{
		Network: "tcp",
		Addr:    "0.0.0.0:9001",
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

	app, close, err := initApp(context.Background(), bootstrap, log.DefaultLogger)
	if err != nil {
		panic(err)
	}
	defer close()

	// start and wait for stop signal
	fmt.Println("Worker started successfully.")
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
func initApp(ctx context.Context, bc *serverConf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	//data, close, err := data.NewData(bc, logger)
	//if err != nil {
	//	return nil, nil, err
	//}
	service, err := service.NewWorkerService(ctx, bc, logger, nil)
	if err != nil {
		return nil, nil, err
	}

	grpcServer := server.NewWorkerGRPCServer(bc.Server, service)
	//reflection.Register(grpcServer.Server)
	httpServer := server.NewHTTPServer(bc.Server, service)
	app := newApp(ctx, logger, httpServer, grpcServer)

	return app, nil, nil
}

func newApp(ctx context.Context, logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
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
