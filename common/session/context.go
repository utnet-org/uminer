package session

//import "context"
//
//type CtxSession string
//
//const (
//	CTX_SESSION_KEY CtxSession = "CtxSessionKey"
//)
//
//func CtxSessionKey() CtxSession {
//	return CTX_SESSION_KEY
//}
//
//func SessionToContext(ctx context.Context, val interface{}) context.Context {
//	return context.WithValue(ctx, CtxSessionKey(), val)
//}
//
//func SessionFromContext(ctx context.Context) *Session {
//	session, ok := ctx.Value(CtxSessionKey()).(*Session)
//	if ok {
//		return session
//	}
//	return nil
//}
