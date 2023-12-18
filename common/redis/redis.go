package redis

import (
	"context"
	"fmt"
	"time"
	"uminer/common/constant"
	"uminer/common/errors"
	"uminer/common/log"

	"github.com/bsm/redislock"
	redisLib "github.com/go-redis/redis/v8"
)

type RedisInstance struct {
	Redis     *redisLib.Client
	RedisLock *redislock.Client
}

const (
	INSTANCE_FMT   = "%v_%v"
	CONNECTION_FMT = "redis://%v:%v@%v/%v"
)

var instances = make(map[string]*RedisInstance)

// redis://<user>:<pass>@localhost:6379/<db>
func GetRedisInstance(connUrl string) (*RedisInstance, error) {
	opt, err := redisLib.ParseURL(connUrl)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErroRedisParseUrlFailed)
	}

	instanceKey := fmt.Sprintf(INSTANCE_FMT, opt.Addr, opt.DB)
	if inst, existed := instances[instanceKey]; existed {
		return inst, nil
	}

	instance := connect(opt)
	instances[instanceKey] = &RedisInstance{
		Redis:     instance,
		RedisLock: redislock.New(instance),
	}

	return instances[instanceKey], nil
}

func connect(options *redisLib.Options) *redisLib.Client {
	rdb := redisLib.NewClient(options)
	return rdb
}

func (r *RedisInstance) LockAndCall(ctx context.Context, key string, ttl time.Duration, f func() error) error {
	lock, err := r.RedisLock.Obtain(ctx, key, ttl, nil)
	if err != nil {
		return errors.Errorf(err, errors.ErrorRedisLockObtainFailed)
	}
	defer func() {
		err := lock.Release(ctx)
		if err != nil {
			log.Errorf(ctx, "release lock err:%v", err)
		}
	}()

	err = f()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisInstance) SMembersMinioRemovingObject() ([]string, error) {
	ctx := context.Background()
	log.Info(ctx, "set:", constant.REDIS_MINIO_REMOVING_OBJECT_SET, " members minio removing object")
	objects, err := r.Redis.SMembers(ctx, constant.REDIS_MINIO_REMOVING_OBJECT_SET).Result()
	if err != nil {
		log.Error(ctx, "set:", constant.REDIS_MINIO_REMOVING_OBJECT_SET,
			" members minio removing object: failed. err:", err.Error())
		return objects, err
	}
	return objects, nil
}

func (r *RedisInstance) SAddMinioRemovingObject(object string) error {
	ctx := context.Background()
	log.Info(ctx, "set:", constant.REDIS_MINIO_REMOVING_OBJECT_SET,
		" add minio removing object:", object)
	_, err := r.Redis.SAdd(ctx, constant.REDIS_MINIO_REMOVING_OBJECT_SET, object).Result()
	if err != nil {
		log.Error(ctx, "set:", constant.REDIS_MINIO_REMOVING_OBJECT_SET,
			" add minio removing object:", object, " failed. err:", err.Error())
		return err
	}
	return nil
}

func (r *RedisInstance) SRemMinioRemovingObject(object string) error {
	ctx := context.Background()
	log.Info(ctx, "set:", constant.REDIS_MINIO_REMOVING_OBJECT_SET,
		" rem minio removing object:", object)
	_, err := r.Redis.SRem(ctx, constant.REDIS_MINIO_REMOVING_OBJECT_SET, object).Result()
	if err != nil {
		log.Error(ctx, "set:", constant.REDIS_MINIO_REMOVING_OBJECT_SET,
			" rem minio removing object:", object, " failed. err:", err.Error())
		return err
	}
	return nil
}
