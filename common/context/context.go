package commctx

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type CtxKey string

const (
	CTX_USERID_KEY     CtxKey = "userId"
	CTX_REQUEST_ID_KEY CtxKey = "requestId" //请求id 存在于请求链路的各个服务
	CTX_DB_TX_KEY      CtxKey = "dbTx"      //数据库事务 存在于common.Transaction()
	CTX_CREATED_AT_KEY CtxKey = "createdAt"
	CTX_PLATFORM_ID    CtxKey = "platformId" //第三方平台Id
	CTX_SPACEID_KEY    CtxKey = "spaceId"    //群组id
)

func RequestIdKey() CtxKey {
	return CTX_REQUEST_ID_KEY
}

func RequestIdToContext(ctx context.Context, val interface{}) context.Context {
	return context.WithValue(ctx, RequestIdKey(), val)
}

func RequestIdFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	id, ok := ctx.Value(RequestIdKey()).(string)
	if ok {
		return id
	}
	return ""
}

func UserIdKey() CtxKey {
	return CTX_USERID_KEY
}

func UserIdToContext(ctx context.Context, val interface{}) context.Context {
	return context.WithValue(ctx, UserIdKey(), val)
}

func UserIdFromContext(ctx context.Context) string {
	id, ok := ctx.Value(UserIdKey()).(string)
	if ok {
		return id
	}
	return ""
}

func SpaceIdKey() CtxKey {
	return CTX_SPACEID_KEY
}

func SpaceIdToContext(ctx context.Context, val interface{}) context.Context {
	return context.WithValue(ctx, SpaceIdKey(), val)
}

func SpaceIdFromContext(ctx context.Context) string {
	id, ok := ctx.Value(SpaceIdKey()).(string)
	if ok {
		return id
	}
	return ""
}

func UserIdAndSpaceIdFromContext(ctx context.Context) (string, string) {
	return UserIdFromContext(ctx), SpaceIdFromContext(ctx)
}

func PlatformIdKey() CtxKey {
	return CTX_PLATFORM_ID
}

func PlatformIdToContext(ctx context.Context, val interface{}) context.Context {
	return context.WithValue(ctx, PlatformIdKey(), val)
}

func PlatformIdFromContext(ctx context.Context) string {
	id, ok := ctx.Value(PlatformIdKey()).(string)
	if ok {
		return id
	}
	return ""
}

func DbTxKey() CtxKey {
	return CTX_DB_TX_KEY
}

func DbTxToContext(ctx context.Context, val interface{}) context.Context {
	return context.WithValue(ctx, DbTxKey(), val)
}

func DbTxFromContext(ctx context.Context) *gorm.DB {
	id, ok := ctx.Value(DbTxKey()).(*gorm.DB)
	if ok {
		return id
	}
	return nil
}

func CreatedAtKey() CtxKey {
	return CTX_CREATED_AT_KEY
}

func CreatedAtToContext(ctx context.Context, val interface{}) context.Context {
	return context.WithValue(ctx, CreatedAtKey(), val)
}

func CreatedAtFromContext(ctx context.Context) int64 {
	createdAt, ok := ctx.Value(CreatedAtKey()).(int64)
	if ok {
		return createdAt
	}
	return 0
}

type noCancel struct {
	ctx context.Context
}

func (c noCancel) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (c noCancel) Done() <-chan struct{}             { return nil }
func (c noCancel) Err() error                        { return nil }
func (c noCancel) Value(key interface{}) interface{} { return c.ctx.Value(key) }

// http请求结束后会cancel ctx，在请求里创建go routine可以调用此方法创建不会取消的ctx，一些下游方法会判断ctx是否已经cancel如gorm的WithContext这种情况需要特别注意
func WithoutCancel(ctx context.Context) context.Context {
	return noCancel{ctx: ctx}
}
