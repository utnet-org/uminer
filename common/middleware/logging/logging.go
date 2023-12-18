package logging

import (
	"context"
	"fmt"
	"uminer/common/log"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// Option is HTTP logging option.
type Option func(*options)

type options struct {
	logger log.Logger
}

// WithLogger with middleware logger.
func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// Server is an server logging middleware.
func Server(opts ...Option) middleware.Middleware {
	options := options{
		logger: log.DefaultLogger,
	}
	for _, o := range opts {
		o(&options)
	}
	log := log.NewHelper("middleware/logging", options.logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				path    string
				request string
			)

			tr, ok := transport.FromServerContext(ctx)
			if ok {
				operation := tr.Operation()
				switch tr.Kind() {
				case transport.KindHTTP:
					if ht, ok := tr.(http.Transporter); ok {
						path = ht.Request().URL.Path
						request = ht.Request().Form.Encode()
					}
				case transport.KindGRPC:
					path = operation
					request = req.(fmt.Stringer).String()
				}
				reply, err := handler(ctx, req)
				if err != nil {
					log.Errorw(ctx,
						"interface", path,
						"request", request,
						"error", err.Error(),
					)
					return nil, err
				}
				log.Infow(ctx,
					"interface", path,
					"request", request,
				)
				log.Debugw(ctx,
					"interface", path,
					"reply", reply)
				return reply, nil
			} else {
				log.Errorw(ctx,
					"interface", "",
					"request", "",
					"error", "transport.FromServerContext err",
				)
			}

			return nil, nil
		}
	}
}

// Client is an client logging middleware.
func Client(opts ...Option) middleware.Middleware {
	options := options{
		logger: log.DefaultLogger,
	}
	for _, o := range opts {
		o(&options)
	}
	log := log.NewHelper("middleware/logging", options.logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				path    string
				request string
			)

			tr, ok := transport.FromClientContext(ctx)
			if ok {
				operation := tr.Operation()
				switch tr.Kind() {
				case transport.KindHTTP:
					if ht, ok := tr.(http.Transporter); ok {
						request = ht.Request().Form.Encode()
						path = ht.Request().URL.Path
					}
				case transport.KindGRPC:
					path = operation
					request = req.(fmt.Stringer).String()
				}
			}
			reply, err := handler(ctx, req)
			if err != nil {
				log.Errorw(ctx,
					"path", path,
					"request", request,
					"error", err.Error(),
				)
				return nil, err
			}
			log.Infow(ctx,
				"interface", path,
				"request", request,
			)
			log.Debugw(ctx,
				"interface", path,
				"reply", reply)
			return reply, nil
		}
	}
}
