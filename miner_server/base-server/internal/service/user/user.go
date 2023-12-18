package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"uminer/common/errors"
	"uminer/common/log"
	"uminer/common/utils"
	api "uminer/miner_server/base-server/api/v1"
	"uminer/miner_server/base-server/internal/common"
	"uminer/miner_server/base-server/internal/conf"
	"uminer/miner_server/base-server/internal/data"
	"uminer/miner_server/base-server/internal/data/model/fun"
)

type UserService struct {
	api.UnimplementedUserServiceServer
	conf       *conf.Bootstrap
	log        *log.Helper
	data       *data.Data
	defaultPVS common.PersistentVolumeSourceExtender
}

func NewUserService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.UserServiceServer {
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

func (s *UserService) ListUser(ctx context.Context, req *api.ListUserRequest) (*api.ListUserReply, error) {
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

	users := make([]*api.UserItem, len(usersTbl))
	for idx, user := range usersTbl {
		bindInfo := make([]*api.Bind, 0)
		if user.Bind != nil {
			for _, a := range user.Bind {
				replyBind := new(api.Bind)
				replyBind.Platform = a.Platform
				replyBind.UserId = a.UserId
				replyBind.UserName = a.UserName
				bindInfo = append(bindInfo, replyBind)
			}
		}
		item := &api.UserItem{
			Id:            user.Id,
			FullName:      user.FullName,
			Email:         user.Email,
			Phone:         user.Phone,
			Gender:        api.GenderType(user.Gender),
			Status:        api.UserStatus(user.Status),
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

	return &api.ListUserReply{
		TotalSize: usersCount,
		Users:     users,
	}, nil
}

func (s *UserService) CheckOrInitUser(ctx context.Context, req *api.CheckOrInitUserRequest) (*api.CheckOrInitUserReply, error) {
	// to check or init user home storage bucket
	return &api.CheckOrInitUserReply{}, nil
}

func (s *UserService) FindUser(ctx context.Context, req *api.FindUserRequest) (*api.FindUserReply, error) {
	a := fun.UserQuery{
		Id:    req.Id,
		Email: req.Email,
		Phone: req.Phone,
	}
	if req.Bind != nil {
		a.Bind = &fun.Bind{
			Platform: req.Bind.Platform,
			UserId:   req.Bind.UserId,
			UserName: req.Bind.UserName,
		}
	}
	user, err := s.data.UserM.Find(ctx, &a)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &api.FindUserReply{
			User: nil,
		}, nil
	}
	bindInfo := make([]*api.Bind, 0)
	if user.Bind != nil {
		for _, a := range user.Bind {
			replyBind := new(api.Bind)
			replyBind.Platform = a.Platform
			replyBind.UserId = a.UserId
			replyBind.UserName = a.UserName
			bindInfo = append(bindInfo, replyBind)
		}
	}
	reply := &api.FindUserReply{
		User: &api.UserItem{
			Id:            user.Id,
			FullName:      user.FullName,
			Email:         user.Email,
			Phone:         user.Phone,
			Gender:        api.GenderType(user.Gender),
			Status:        api.UserStatus(user.Status),
			Password:      user.Password,
			CreatedAt:     user.CreatedAt.Unix(),
			UpdatedAt:     user.UpdatedAt.Unix(),
			Bind:          bindInfo,
			ResourcePools: user.ResourcePools,
			Desc:          user.Desc,
			//Permission:    (*v1.UserPermission)(user.Permission),
			MinioUserName: user.MinioUserName,
			Buckets:       user.Buckets,
		},
	}

	return reply, nil
}

func (s *UserService) AddUser(ctx context.Context, req *api.AddUserRequest) (*api.AddUserReply, error) {
	cond := fun.UserQuery{
		Email: req.Email,
		Phone: req.Phone,
	}
	if req.Bind != nil {
		cond.Bind = &fun.Bind{
			Platform: req.Bind.Platform,
			UserId:   req.Bind.UserId,
			UserName: req.Bind.UserName,
		}
	}
	existed, err := s.data.UserM.Find(ctx, &cond)
	if err != nil {
		return nil, err
	}
	if existed != nil {
		return nil, errors.Errorf(nil, errors.ErrorUserAccountExisted)
	}

	password, err := utils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := fun.UserAdd{}
	user.Email = req.Email
	user.Phone = req.Phone
	user.Password = password
	user.Id = utils.GetUUIDWithoutSeparator()
	user.FullName = req.FullName
	user.Gender = int32(req.Gender)
	user.Status = int32(api.UserStatus_ACTIVITY)
	user.Bind = cond.Bind
	user.Desc = req.Desc
	u, err := s.data.UserM.Add(ctx, &user)
	if err != nil {
		return nil, err
	}

	bindInfo := make([]*api.Bind, 0)
	if u.Bind != nil {
		for _, a := range u.Bind {
			replyBind := new(api.Bind)
			replyBind.Platform = a.Platform
			replyBind.UserId = a.UserId
			replyBind.UserName = a.UserName
			bindInfo = append(bindInfo, replyBind)
		}
	}
	reply := &api.AddUserReply{
		User: &api.UserItem{
			Id:       u.Id,
			FullName: u.FullName,
			Email:    u.Email,
			Phone:    u.Phone,
			Gender:   api.GenderType(u.Gender),
			Status:   api.UserStatus(u.Status),
			Password: u.Password,
			Bind:     bindInfo,
		},
	}

	reply.User.CreatedAt = u.CreatedAt.Unix()
	reply.User.UpdatedAt = u.UpdatedAt.Unix()

	return reply, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *api.UpdateUserRequest) (*api.UpdateUserReply, error) {
	userId := req.Id
	bindInfo := make([]*fun.Bind, 0)
	if req.Bind != nil {
		for _, a := range req.Bind {
			reqBind := new(fun.Bind)
			reqBind.Platform = a.Platform
			reqBind.UserId = a.UserId
			reqBind.UserName = a.UserName
			bindInfo = append(bindInfo, reqBind)
		}
	}
	user := fun.UserUpdate{
		FullName:      req.FullName,
		Email:         req.Email,
		Phone:         req.Phone,
		Gender:        int32(req.Gender),
		Status:        int32(req.Status),
		ResourcePools: req.ResourcePools,
		Desc:          req.Desc,
	}
	if len(bindInfo) > 0 {
		user.Bind = bindInfo
	}
	if req.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, err
		}
		user.Password = string(password)
	}
	result, err := s.data.UserM.Update(ctx, &fun.UserUpdateCond{Id: userId}, &user)
	if err != nil {
		return nil, err
	}

	bindInfo2 := make([]*api.Bind, 0)
	if result.Bind != nil {
		for _, a := range result.Bind {
			replyBind := new(api.Bind)
			replyBind.Platform = a.Platform
			replyBind.UserId = a.UserId
			replyBind.UserName = a.UserName
			bindInfo2 = append(bindInfo2, replyBind)
		}
	}
	return &api.UpdateUserReply{
		User: &api.UserItem{
			Id:       result.Id,
			FullName: result.FullName,
			Email:    result.Email,
			Phone:    result.Phone,
			Gender:   api.GenderType(result.Gender),
			Status:   api.UserStatus(result.Status),
			Password: result.Password,
			Bind:     bindInfo2,
			Desc:     result.Desc,
		},
	}, nil
}

func (s *UserService) ListUserInCond(ctx context.Context, req *api.ListUserInCondRequest) (*api.ListUserInCondReply, error) {
	users, err := s.data.UserM.ListIn(ctx, &fun.UserListIn{Ids: req.Ids})
	if err != nil {
		return nil, err
	}

	userItems := make([]*api.UserItem, len(users))
	for idx, user := range users {
		bindInfo := make([]*api.Bind, 0)
		if user.Bind != nil {
			for _, a := range user.Bind {
				replyBind := new(api.Bind)
				replyBind.Platform = a.Platform
				replyBind.UserId = a.UserId
				replyBind.UserName = a.UserName
				bindInfo = append(bindInfo, replyBind)
			}
		}
		item := &api.UserItem{
			Id:        user.Id,
			FullName:  user.FullName,
			Email:     user.Email,
			Phone:     user.Phone,
			Gender:    api.GenderType(user.Gender),
			Status:    api.UserStatus(user.Status),
			Password:  user.Password,
			CreatedAt: user.CreatedAt.Unix(),
			UpdatedAt: user.UpdatedAt.Unix(),
			Bind:      bindInfo,
		}
		userItems[idx] = item
	}
	return &api.ListUserInCondReply{
		Users: userItems,
	}, nil
}
