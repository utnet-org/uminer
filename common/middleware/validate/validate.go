package validate

import (
	"context"
	"uminer/common/errors"

	"github.com/go-kratos/kratos/v2/middleware"
)

// Option is HTTP logging option.
type Option func(*options)

type options struct {
}

type validate interface {
	Validate() error
}

// Server is an server logging middleware.
func Server(opts ...Option) middleware.Middleware {
	options := options{}
	for _, o := range opts {
		o(&options)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if v, ok := req.(validate); ok {
				err := v.Validate()
				if err != nil {
					return nil, errors.Errorf(err, errors.ErrorInvalidRequestParameter)
				}
			}

			reply, err := handler(ctx, req)
			if err != nil {
				return nil, err
			}

			return reply, nil
		}
	}
}
