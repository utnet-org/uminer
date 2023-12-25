package common

import (
	"context"
	"time"
)

type PipelineCallbackReq struct {
	Id           string    `json:"id"`
	UserID       string    `json:"userID"`
	Namespace    string    `json:"namespace"`
	CurrentState string    `json:"currentState"`
	CurrentTime  time.Time `json:"currentTime"`
}

type PipelineCallback interface {
	PipelineCallback(ctx context.Context, req *PipelineCallbackReq) string //返回值为OK|RE|EX OK:成功，RE:重试,EX:异常
}

const (
	PipeLineCallbackOK = "OK"
	PipeLineCallbackRE = "RE"
)
