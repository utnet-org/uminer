package model

import (
	"context"
	"errors"
	stderrors "errors"
	"gorm.io/gorm/clause"
	"uminer/common/log"
	"uminer/common/utils/collections/set"
	"uminer/miner_server/base-server/internal/data/model/fun"

	commerrors "uminer/common/errors"

	"gorm.io/gorm"
)

type UserMethod interface {
	List(ctx context.Context, condition *fun.UserList) ([]*fun.User, error)
	Count(ctx context.Context, condition *fun.UserList) (int64, error)
	Find(ctx context.Context, condition *fun.UserQuery) (*fun.User, error)
	Add(ctx context.Context, user *fun.UserAdd) (*fun.User, error)
	Update(ctx context.Context, condition *fun.UserUpdateCond, user *fun.UserUpdate) (*fun.User, error)
	ListIn(ctx context.Context, condition *fun.UserListIn) ([]*fun.User, error)
	UpdateConfig(ctx context.Context, userId string, config map[string]string) error
	GetConfig(ctx context.Context, userId string) (map[string]string, error)
}

type userMethod struct {
	log *log.Helper
	db  *gorm.DB
}

func NewUserMethod(db *gorm.DB, logger log.Logger) UserMethod {
	return &userMethod{
		log: log.NewHelper("UserMethod", logger),
		db:  db,
	}
}

func (d *userMethod) List(ctx context.Context, condition *fun.UserList) ([]*fun.User, error) {
	db := d.db
	users := make([]*fun.User, 0)

	db = condition.Pagination(db)
	db = condition.Order(db)
	db = condition.Where(db)
	db = condition.Or(db)
	db.Find(&users)

	return users, nil
}

func (d *userMethod) Count(ctx context.Context, condition *fun.UserList) (int64, error) {
	db := d.db
	var count int64

	db = condition.Where(db)
	db = condition.Or(db)

	db.Model(&fun.User{}).Count(&count)
	return count, nil
}

func (d *userMethod) Find(ctx context.Context, condition *fun.UserQuery) (*fun.User, error) {
	db := d.db

	var user fun.User
	var result *gorm.DB
	if condition.Bind == nil {
		result = db.Where(&fun.User{
			Id:    condition.Id,
			Email: condition.Email,
			Phone: condition.Phone,
		}).First(&user)
	} else {
		querySql := "1 = 1"
		params := make([]interface{}, 0)
		if condition.Email != "" {
			querySql += " and email = ? "
			params = append(params, condition.Email)
			if condition.Bind.UserId != "" {
				querySql += " or (JSON_CONTAINS(bind,JSON_OBJECT('platform', ?))"
				params = append(params, condition.Bind.Platform)
				querySql += " and JSON_CONTAINS(bind,JSON_OBJECT('userId', ?)))"
				params = append(params, condition.Bind.UserId)
			} else {
				querySql += " and JSON_CONTAINS(bind,JSON_OBJECT('platform', ?))"
				params = append(params, condition.Bind.Platform)
			}
		} else {
			querySql += " and JSON_CONTAINS(bind,JSON_OBJECT('platform', ?))"
			params = append(params, condition.Bind.Platform)
			querySql += " and JSON_CONTAINS(bind,JSON_OBJECT('userId', ?))"
			params = append(params, condition.Bind.UserId)
		}
		result = db.Where(querySql, params...).First(&user)
	}
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (d *userMethod) Add(ctx context.Context, user *fun.UserAdd) (*fun.User, error) {
	db := d.db
	bindInfo := make([]*fun.Bind, 0)
	if user.Bind != nil {
		bindInfo = append(bindInfo, user.Bind)
	}
	u := fun.User{
		Id:            user.Id,
		FullName:      user.FullName,
		Gender:        user.Gender,
		Email:         user.Email,
		Phone:         user.Phone,
		Password:      user.Password,
		Status:        user.Status,
		Bind:          bindInfo,
		ResourcePools: user.ResourcePools,
		Desc:          user.Desc,
	}

	result := db.Omit("ftp_user_name", "minio_user_name").Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func (d *userMethod) Update(ctx context.Context, cond *fun.UserUpdateCond, user *fun.UserUpdate) (*fun.User, error) {
	if cond.Id == "" {
		return nil, gorm.ErrPrimaryKeyRequired
	}

	condition := fun.User{
		Id:    cond.Id,
		Email: cond.Email,
		Phone: cond.Phone,
	}

	result := d.db.Model(&condition).Updates(fun.User{
		FullName:      user.FullName,
		Email:         user.Email,
		Phone:         user.Phone,
		Gender:        user.Gender,
		Password:      user.Password,
		Status:        user.Status,
		Bind:          user.Bind,
		ResourcePools: user.ResourcePools,
		Desc:          user.Desc,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return d.Find(ctx, &fun.UserQuery{
		Id:    cond.Id,
		Email: cond.Email,
		Phone: cond.Phone,
	})
}

func (d *userMethod) ListIn(ctx context.Context, condition *fun.UserListIn) ([]*fun.User, error) {
	if len(condition.Ids) < 1 {
		return nil, gorm.ErrMissingWhereClause
	}
	idsSet := set.NewStrings(condition.Ids...)
	var users []*fun.User
	result := d.db.Find(&users, idsSet.Values())
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (d *userMethod) UpdateConfig(ctx context.Context, userId string, config map[string]string) error {
	db := d.db
	if userId == "" || len(config) == 0 {
		return commerrors.Errorf(nil, commerrors.ErrorInvalidRequestParameter)
	}

	configUp := make(map[string]string)
	for k, v := range config {
		if v != "" {
			configUp[k] = v
		}
	}

	c := &fun.UserConfig{UserId: userId, Config: configUp}
	res := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(c)
	if res.Error != nil {
		return commerrors.Errorf(nil, commerrors.ErrorDBUpdateFailed)
	}

	return nil
}

func (d *userMethod) GetConfig(ctx context.Context, userId string) (map[string]string, error) {
	db := d.db
	c := &fun.UserConfig{}

	res := db.First(c, "user_id = ?", userId)
	if res.Error != nil && !stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, commerrors.Errorf(res.Error, commerrors.ErrorDBFindFailed)
	}

	return c.Config, nil
}
