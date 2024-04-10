package worker

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/urfave/cli/v2"
	"time"
	"uminer/common/graceful"
	"uminer/miner-server/cmd"
	"uminer/miner-server/cmd/utlog"

	"uminer/common/log"
	"uminer/miner-server/serverConf"

	"uminer/miner-server/server"
	"uminer/miner-server/service"

	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func StartWorkerServer(c *cli.Context) error {

	cmd.WorkerServerIP = c.String("serverip")

	utlog.Mainlog.Info("Initializing utility worker")

	// activate worker server
	httpServer := &serverConf.Server_HTTP{
		Network: "tcp",
		Addr:    cmd.WorkerServerIP + ":6001",
		Timeout: &duration.Duration{Seconds: 60},
	}

	grpcServer := &serverConf.Server_GRPC{
		Network: "tcp",
		Addr:    cmd.WorkerServerIP + ":7001",
		Timeout: &duration.Duration{Seconds: 60},
	}

	// build Bootstrap configuration
	bootstrap := &serverConf.Bootstrap{
		App: &serverConf.App{
			// set App
		},
		Server: &serverConf.Server{
			// set Server
			Http: httpServer,
			Grpc: grpcServer,
		},
		Data: &serverConf.Data{
			// set Data
		},
		Storage: []byte("my_storage_data"), // 设置 Storage 字段
	}

	app, close, err := initWorkerApp(context.Background(), bootstrap, log.DefaultLogger)
	if err != nil {
		panic(err)
	}
	defer close()

	// start and wait for stop signal
	utlog.Mainlog.Info("Worker started successfully")
	if err := app.Run(); err != nil {
		panic(err)
	}

	// exit gently
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		graceful.Shutdown(ctx)
	}()

	return nil
}

// initWorkerApp init kratos application.
func initWorkerApp(ctx context.Context, bc *serverConf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	//data, close, err := data.NewData(bc, logger)
	//if err != nil {
	//	return nil, nil, err
	//}
	newService, err := service.NewWorkerService(ctx, bc, logger, nil)
	if err != nil {
		return nil, nil, err
	}

	grpcServer := server.NewWorkerGRPCServer(bc.Server, newService)

	httpServer := server.NewHTTPServer(bc.Server, newService)

	app := newWorker(ctx, logger, httpServer, grpcServer)

	return app, nil, nil
}

func newWorker(ctx context.Context, logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
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
