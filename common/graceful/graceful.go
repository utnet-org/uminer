package graceful

import (
	"context"
	"sync"
	"sync/atomic"
	"uminer/common/log"
)

/*
main里增加shutdown，超时时间根据实际调整

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		graceful.Shutdown(ctx)
	}()

在开协程处理时:
graceful.AddOne()

	go func(){
		defer graceful.Done()

....
}()

批处理时查询是否正在关闭，或者用Shutting()返回的chan
graceful.AddOne()

	go func() {
		defer graceful.Done()
		for i := 0; i < 10; i++ {
			if graceful.IsShutting() {
				break
			}
		....
		}
	}()
*/
type graceful struct {
	wg         *sync.WaitGroup
	once       *sync.Once
	allDone    chan struct{}
	shutting   chan struct{}
	isShutting int32
}

var (
	defaultGraceful = newGraceful()
)

func newGraceful() *graceful {
	return &graceful{
		wg:       new(sync.WaitGroup),
		once:     new(sync.Once),
		allDone:  make(chan struct{}),
		shutting: make(chan struct{}),
	}
}

const inShutting = 1

func AddOne() {
	defaultGraceful.AddOne()
}

func Done() {
	defaultGraceful.Done()
}

func Shutdown(ctx context.Context) {
	defaultGraceful.Shutdown(ctx)
}

func IsShutting() bool {
	return defaultGraceful.IsShutting()
}

func Shutting() <-chan struct{} {
	return defaultGraceful.shutting
}

func (g *graceful) AddOne() {
	g.wg.Add(1)
}

func (g *graceful) Done() {
	g.wg.Done()
}

func (g *graceful) Shutdown(ctx context.Context) {
	g.once.Do(
		func() {
			go func() {
				close(g.shutting)
				atomic.StoreInt32(&g.isShutting, inShutting)
				defer close(g.allDone)
				g.wg.Wait()
			}()
		})
	for {
		select {
		case <-g.allDone:
			log.Info(ctx, "all done")
			return
		case <-ctx.Done():
			log.Info(ctx, "ctx done")
			return
		}
	}
}

func (g *graceful) IsShutting() bool {
	return atomic.LoadInt32(&g.isShutting) == inShutting
}

func (g *graceful) Shutting() <-chan struct{} {
	return g.shutting
}
