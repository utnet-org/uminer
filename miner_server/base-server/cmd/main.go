package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"gopkg.in/yaml.v2"
	"time"
	"uminer/miner_server/base-server/internal/conf"
	"uminer/miner_server/base-server/internal/data"
	server2 "uminer/miner_server/base-server/internal/server"
	"uminer/miner_server/base-server/internal/service"

	"uminer/common/errors"
	"uminer/common/graceful"
	"uminer/common/log"
	"uminer/common/third_party/kratos/config"
	"uminer/common/third_party/kratos/config/file"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	// Name is the name of the compiled software.
	// go build -ldflags "-X main.Name=xyz"
	Name string
	// Version is the version of the compiled software.
	// go build -ldflags "-X main.Version=x.y.z"
	v       bool
	Version string
	// Built is the build time of the compiled software.
	// go build -ldflags "-X main.Built=2021-06-02 17:29:36"
	Built string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.BoolVar(&v, "v", false, "software version, eg: -v")
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf config.yaml")
}

// marshalJson error, when values contains map[interface{}]interface{}. 临时修改代码后放入third_party, 后续升级kratos解决
func main() {
	flag.Parse()
	//if v {
	//	fmt.Printf("Version: %s\nBuilt: %s\n", Version, Built)
	//	return
	//}
	if flagconf == "" {
		fmt.Printf("Miss param: -conf, use -h to get help\n")
		return
	}
	conf, c, err := initConf()
	if err != nil {
		panic(err)
	}

	l := log.ConvertFromString(conf.App.LogLevel)
	log.DefaultLogger.ResetLevel(l)
	log.DefaultGormLogger.LogMode(log.ConvertToGorm(l))
	err = c.Watch("app.logLevel", func(k string, v config.Value) {
		ls, _ := v.String()
		l := log.ConvertFromString(ls)
		log.DefaultLogger.ResetLevel(l)
		log.DefaultGormLogger.LogMode(log.ConvertToGorm(l))
	})
	if err != nil {
		log.Infof(context.TODO(), "watch app.logLevel err:%v", err)
	}

	app, close, err := initApp(context.Background(), conf, log.DefaultLogger)
	if err != nil {
		panic(err)
	}
	defer close()

	// start and wait for stop signal
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

func newApp(ctx context.Context, logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.Context(ctx),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

// initApp init kratos application.
func initApp(ctx context.Context, bc *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	data, close, err := data.NewData(bc, logger)
	if err != nil {
		return nil, nil, err
	}
	service, err := service.NewService(ctx, bc, logger, data)
	if err != nil {
		return nil, nil, err
	}

	grpcServer := server2.NewGRPCServer(bc.Server, service)
	httpServer := server2.NewHTTPServer(bc.Server, service)
	app := newApp(ctx, logger, httpServer, grpcServer)

	return app, close, nil
}

func initStorageConf(c config.Config) ([]byte, error) {
	m := make(map[string]interface{})
	value := c.Value("module.storage.source")
	err := value.Scan(&m)
	if err != nil {
		return nil, err
	}

	storageConf, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Errorf(nil, errors.ErrorJsonMarshal)
	}
	return storageConf, nil
}

func initConf() (*conf.Bootstrap, config.Config, error) {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := c.Load(); err != nil {
		return nil, nil, err
	}
	var conf conf.Bootstrap
	if err := c.Scan(&conf); err != nil {
		return nil, nil, err
	}
	if Name != "" {
		conf.App.Name = Name
	} else {
		Name = conf.App.Name
	}
	if Version != "" {
		conf.App.Version = Version
	} else {
		Version = conf.App.Version
	}

	// json Marshal []byte
	//storageConf, err := initStorageConf(c)
	//if err != nil {
	//	return nil, nil, err
	//}
	//conf.Storage = storageConf

	return &conf, c, nil
}
