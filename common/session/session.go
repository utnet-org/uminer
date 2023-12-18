package session

//import (
//	"context"
//	"server/common/constant"
//)
//
//const (
//	SESSION_WORKSPACE = "workspace"
//)
//
//type Session struct {
//	ctx        context.Context   `json:"-"`
//	store      SessionStore      `json:"-"`
//	Id         string            `json:"id"`
//	UserId     string            `json:"userId"`
//	Status     int32             `json:"status"`
//	CreatedAt  int64             `json:"createdAt"`
//	Attributes map[string]string `json:"attributes"`
//}
//
//func (s *Session) GetWorkspace() string {
//	if workspaceId, ok := s.Attributes[SESSION_WORKSPACE]; !ok {
//		return constant.SYSTEM_WORKSPACE_DEFAULT
//	} else {
//		return workspaceId
//	}
//}
//
//func (s *Session) SetWorkspace(workspaceId string) error {
//	if workspaceId == constant.SYSTEM_WORKSPACE_DEFAULT {
//		delete(s.Attributes, SESSION_WORKSPACE)
//	} else {
//		s.Attributes[SESSION_WORKSPACE] = workspaceId
//	}
//	return s.sync()
//}
//
//func (s *Session) IsDefaultWorkspace() bool {
//	if workspaceId, ok := s.Attributes[SESSION_WORKSPACE]; !ok {
//		return true
//	} else {
//		return workspaceId == constant.SYSTEM_WORKSPACE_DEFAULT
//	}
//}
//
//func (s *Session) sync() error {
//	if s.store == nil {
//		return nil
//	}
//	if err := s.store.Update(s.ctx, s); err != nil {
//		return err
//	}
//	return nil
//}
