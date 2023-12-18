package session

//import (
//	"context"
//
//	"net/url"
//	"server/common/errors"
//	"server/common/redis"
//
//	"github.com/go-kratos/kratos/v2/log"
//	redisLib "github.com/go-redis/redis/v8"
//	jsoniter "github.com/json-iterator/go"
//)
//
//type SessionStore interface {
//	Create(ctx context.Context, session *Session) error
//	Get(ctx context.Context, sessionId string) (*Session, error)
//	Update(ctx context.Context, session *Session) error
//	Delete(ctx context.Context, sessionId string) error
//}
//
//type SessionStoreConfig struct {
//	RedisAddr     string
//	RedisUsername string
//	RedisPassword string
//	RedisDBIndex  string
//}
//
//func NewSessionStore(key string, config SessionStoreConfig, logger log.Logger) SessionStore {
//	logHelper := log.NewHelper("Session", logger)
//
//	redisUrl := url.URL{
//		Scheme: "redis",
//		Host:   config.RedisAddr,
//		Path:   config.RedisDBIndex,
//		User:   url.UserPassword(config.RedisUsername, config.RedisPassword),
//	}
//	rdb, err := redis.GetRedisInstance(redisUrl.String())
//	if err != nil {
//		panic(err)
//	}
//
//	return &RemoteSessionStore{
//		StoreKey: key,
//		config:   config,
//		Instance: rdb,
//		logger:   logHelper,
//	}
//}
//
//type RemoteSessionStore struct {
//	StoreKey string
//	config   SessionStoreConfig
//	Instance *redis.RedisInstance
//	logger   *log.Helper
//}
//
//func (s *RemoteSessionStore) Create(ctx context.Context, session *Session) error {
//	if session == nil || session.Id == "" {
//		return errors.Errorf(nil, errors.ErrorSessionIdNotFound)
//	}
//	sbytes, err := jsoniter.Marshal(&session)
//	if err != nil {
//		return errors.Errorf(err, errors.ErrorJsonMarshal)
//	}
//
//	_, err = s.Instance.Redis.HSet(ctx, s.StoreKey, session.Id, sbytes).Result()
//	if err != nil {
//		return errors.Errorf(err, errors.ErroRedisHSetFailed)
//	}
//	return nil
//}
//
//func (s *RemoteSessionStore) Get(ctx context.Context, sessionId string) (*Session, error) {
//	if sessionId == "" {
//		return nil, errors.Errorf(nil, errors.ErrorSessionIdNotFound)
//	}
//	resultStr, err := s.Instance.Redis.HGet(ctx, s.StoreKey, sessionId).Result()
//	if err != nil {
//		if redisLib.Nil == err {
//			return nil, nil
//		}
//		return nil, errors.Errorf(err, errors.ErroRedisHGetFailed)
//	}
//
//	var session Session
//	err = jsoniter.Unmarshal([]byte(resultStr), &session)
//	if err != nil {
//		return nil, errors.Errorf(err, errors.ErrorJsonUnmarshal)
//	}
//	session.ctx = ctx
//	session.store = s
//	return &session, nil
//}
//
//func (s *RemoteSessionStore) Update(ctx context.Context, session *Session) error {
//	return s.Create(ctx, session)
//}
//
//func (s *RemoteSessionStore) Delete(ctx context.Context, sessionId string) error {
//	if sessionId == "" {
//		return errors.Errorf(nil, errors.ErrorSessionIdNotFound)
//	}
//
//	_, err := s.Instance.Redis.HDel(ctx, s.StoreKey, sessionId).Result()
//	if err != nil {
//		return errors.Errorf(err, errors.ErroRedisHDelFailed)
//	}
//	return nil
//}
