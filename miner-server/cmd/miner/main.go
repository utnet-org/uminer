package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	minerGprc "github.com/go-kratos/kratos/v2/transport/grpc"
	minerHttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang/protobuf/ptypes/duration"
	"time"
	"uminer/common/graceful"
	"uminer/common/log"
	"uminer/miner-server/cmd"
	"uminer/miner-server/server"
	"uminer/miner-server/serverConf"
	"uminer/miner-server/service"
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

	httpServer := server.NewHTTPServer(bc.Server, newService)
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
