package session

//import (
//	"context"
//	"net/http"
//	commctx "server/common/context"
//	"server/common/errors"
//	ss "server/common/session"
//	"strings"
//
//	"github.com/go-kratos/kratos/v2/middleware"
//	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"
//)
//
//// Option is HTTP logging option.
//type Option func(*Options)
//
//type Options struct {
//	NoAuthUris   []string
//	Store        ss.SessionStore
//	CheckSession func(ctx context.Context, s *ss.Session) error
//}
//
//// Server is an server logging middleware.
//func Server(opts ...Option) middleware.Middleware {
//	options := Options{}
//	for _, o := range opts {
//		o(&options)
//	}
//	return func(handler middleware.Handler) middleware.Handler {
//		return func(ctx context.Context, req interface{}) (interface{}, error) {
//			var request *http.Request
//			if info, ok := kratosHttp.FromServerContext(ctx); ok {
//				request = info.Request
//			} else {
//				return handler(ctx, req)
//			}
//
//			needAuth := true
//			for _, i := range options.NoAuthUris {
//				if strings.Contains(request.RequestURI, i) {
//					needAuth = false
//				}
//			}
//
//			if needAuth {
//				userId := commctx.UserIdFromContext(ctx)
//				store := options.Store
//				session, err := store.Get(ctx, userId)
//				if err != nil {
//					return nil, errors.Errorf(err, errors.ErrorUserGetAuthSessionFailed)
//				}
//				if session == nil {
//					return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
//				}
//				if options.CheckSession != nil {
//					if err := options.CheckSession(ctx, session); err != nil {
//						return nil, err
//					}
//				}
//
//				ctx = ss.SessionToContext(ctx, session)
//			}
//
//			return handler(ctx, req)
//		}
//	}
//}
