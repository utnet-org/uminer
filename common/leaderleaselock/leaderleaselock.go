package leaderleaselock

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"uminer/common/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

type OnStartedLeading func(ctx context.Context)
type LeaderLeaseLock struct {
	leaselock   *resourcelock.LeaseLock
	candidateid string
}

func NewLeaderLeaselock(namespace string, lockName string, kubeConfig *rest.Config) *LeaderLeaseLock {

	if lockName == "" {
		panic("unable to get lease lock resource name (missing lease-lock-name flag).")
	}

	if namespace == "" {
		namespace = "default"
	}

	k8sClient := clientset.NewForConfigOrDie(kubeConfig)
	candidateid := utils.GetUUIDWithoutSeparator()

	lock := &LeaderLeaseLock{
		leaselock: &resourcelock.LeaseLock{
			LeaseMeta: metav1.ObjectMeta{
				Name:      lockName,
				Namespace: namespace,
			},
			Client: k8sClient.CoordinationV1(),
			LockConfig: resourcelock.ResourceLockConfig{
				Identity: candidateid,
			},
		},
		candidateid: candidateid,
	}

	return lock
}
func (lock *LeaderLeaseLock) RunOrRetryLeaderElection(ctx context.Context, onStartedLeading OnStartedLeading) {
	go func() {
		//when network is down, LeaderElection will Die,
		//so we should use for loop to retry LeaderElection
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx Stop RunOrRetryLeaderElection!!!")
				break
			default:
			}

			lock.runOrDieLeaderElection(ctx, onStartedLeading)

			fmt.Println("leaderelection Die!!!")
		}
	}()
}

func (lock *LeaderLeaseLock) runOrDieLeaderElection(ctx context.Context, onStartedLeading OnStartedLeading) {
	// use a Go context so we can tell the leaderelection code when we
	// want to step down
	// we can not know ctx is WithCancelContext, so create a WithCancelContext
	lockCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// listen for interrupts or the Linux SIGTERM signal and cancel
	// our context, which the leader election code will observe and
	// step down
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		cancel()
	}()

	// start the leader election code loop
	leaderelection.RunOrDie(lockCtx, leaderelection.LeaderElectionConfig{
		Lock: lock.leaselock,
		// IMPORTANT: you MUST ensure that any code you have that
		// is protected by the lease must terminate **before**
		// you call cancel. Otherwise, you could have a background
		// loop still running and another process could
		// get elected before your background loop finished, violating
		// the stated goal of the lease.
		ReleaseOnCancel: true,
		LeaseDuration:   60 * time.Second,
		RenewDeadline:   15 * time.Second,
		RetryPeriod:     5 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				// we're notified when we start - this is where you would
				// usually put your code
				onStartedLeading(ctx)
			},
			OnStoppedLeading: func() {
				// we can do cleanup here
				//os.Exit(0)
				fmt.Println("leaderelection OnStoppedLeading!")
			},
			OnNewLeader: func(identity string) {
				// we're notified when new leader elected
				if identity == lock.candidateid {
					// I just got the lock
					return
				}
			},
		},
	})
}
