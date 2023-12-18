package ctxcopy

import (
	"context"
	commctx "uminer/common/context"

	"github.com/go-kratos/kratos/v2/middleware"
	"google.golang.org/grpc/metadata"
)

// Option is HTTP logging option.
type Option func(*options)

type options struct {
}

// Server is an server logging middleware.
func Server(opts ...Option) middleware.Middleware {
	options := options{}
	for _, o := range opts {
		o(&options)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				md = metadata.Pairs()
			}

			reqId := md.Get(string(commctx.RequestIdKey()))
			if len(reqId) > 0 {
				ctx = commctx.RequestIdToContext(ctx, reqId[0])
			}
			return handler(ctx, req)
		}
	}
}

func Client(opts ...Option) middleware.Middleware {
	options := options{}
	for _, o := range opts {
		o(&options)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			md, ok := metadata.FromOutgoingContext(ctx)
			if !ok {
				md = metadata.Pairs()
			}
			md.Set(string(commctx.RequestIdKey()), commctx.RequestIdFromContext(ctx))
			ctx = metadata.NewOutgoingContext(ctx, md)
			return handler(ctx, req)
		}
	}
}
