package transaction

import (
	"context"
	commctx "uminer/common/context"

	"gorm.io/gorm"
)

// tx放到ctx里
func Transaction(ctx context.Context, db *gorm.DB, fc func(ctx context.Context) error) error {
	return db.Transaction(func(tx *gorm.DB) error {
		return fc(commctx.DbTxToContext(ctx, tx))
	})
}

// ctx里有tx，则从ctx里取，否则返回db
func GetDBFromCtx(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx := commctx.DbTxFromContext(ctx)
	if tx != nil {
		return tx
	}

	return db.WithContext(ctx)
}

type GetDB func(ctx context.Context) *gorm.DB
