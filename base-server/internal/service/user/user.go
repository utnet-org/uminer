package user

import (
	"context"
	"uminer/base-server/api/v1"
	"uminer/base-server/internal/common"
	"uminer/base-server/internal/conf"
	"uminer/base-server/internal/data"
	"uminer/base-server/internal/data/model/fun"
	"uminer/common/log"
)

type UserService struct {
	v1.UnimplementedUserServiceServer
	conf       *conf.Bootstrap
	log        *log.Helper
	data       *data.Data
	defaultPVS common.PersistentVolumeSourceExtender
}

func NewUserService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) v1.UserServiceServer {
	pvs, err := common.BuildStorageSource(conf.Storage)
	if err != nil {
		panic(err)
	}
	if pvs.Size() == 0 {
		panic("mod init failed, missing config [module.storage.source]")
	}
	return &UserService{
		conf:       conf,
		log:        log.NewHelper("UserService", logger),
		data:       data,
		defaultPVS: *pvs,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	usersTbl, err := s.data.UserM.List(ctx, &fun.UserList{
		SortBy:    req.SortBy,
		OrderBy:   req.OrderBy,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		FullName:  req.FullName,
		Email:     req.Email,
		SearchKey: req.SearchKey,
		Phone:     req.Phone,
		Status:    int32(req.Status),
		Desc:      req.Desc,
	})
	if err != nil {
		return nil, err
	}

	usersCount, err := s.data.UserM.Count(ctx, &fun.UserList{
		FullName:  req.FullName,
		Email:     req.Email,
		SearchKey: req.SearchKey,
		Status:    int32(req.Status),
	})
	if err != nil {
		return nil, err
	}

	users := make([]*v1.UserItem, len(usersTbl))
	for idx, user := range usersTbl {
		bindInfo := make([]*v1.Bind, 0)
		if user.Bind != nil {
			for _, a := range user.Bind {
				replyBind := new(v1.Bind)
				replyBind.Platform = a.Platform
				replyBind.UserId = a.UserId
				replyBind.UserName = a.UserName
				bindInfo = append(bindInfo, replyBind)
			}
		}
		item := &v1.UserItem{
			Id:            user.Id,
			FullName:      user.FullName,
			Email:         user.Email,
			Phone:         user.Phone,
			Gender:        v1.GenderType(user.Gender),
			Status:        v1.UserStatus(user.Status),
			Password:      user.Password,
			CreatedAt:     user.CreatedAt.Unix(),
			UpdatedAt:     user.UpdatedAt.Unix(),
			Bind:          bindInfo,
			ResourcePools: user.ResourcePools,
			Desc:          user.Desc,
			//Permission:    (*v1.UserPermission)(user.Permission),
			MinioUserName: user.MinioUserName,
			Buckets:       user.Buckets,
		}
		users[idx] = item
	}

	return &v1.ListUserReply{
		TotalSize: usersCount,
		Users:     users,
	}, nil
}
