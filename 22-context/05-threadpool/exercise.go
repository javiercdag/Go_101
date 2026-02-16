package threadpool

import (
	"context"
	"fmt"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

type Runnable interface {
	Run(context.Context) error
}

type ThreadPool interface {
	Run(Runnable)
	Close()
}

type threadPool struct {
	tasks     chan Runnable
	errChan   chan error
	wg        sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	closeOnce sync.Once
}

func NewThreadPool(n int) (ThreadPool, chan error) {
	ctx, cancel := context.WithCancel(context.Background())

	tp := &threadPool{
		tasks:   make(chan Runnable, 100),
		errChan: make(chan error, 100),
		ctx:     ctx,
		cancel:  cancel,
	}

	tp.wg.Add(n)
	for i := 0; i < n; i++ {
		go tp.worker()
	}

	return tp, tp.errChan
}

func (tp *threadPool) worker() {
	defer tp.wg.Done()

	for {
		select {
		case <-tp.ctx.Done():
			return

		case task := <-tp.tasks:
			err := task.Run(tp.ctx)

			if err != nil {
				select {
				case tp.errChan <- err:
					// res
				default:
					fmt.Println("Error channel full, dropping error:", err)
				}
			}
		}
	}
}

func (tp *threadPool) Run(r Runnable) {
	select {
	case tp.tasks <- r:
		// res
	case <-tp.ctx.Done():
		// res
	}
}

func (tp *threadPool) Close() {
	tp.closeOnce.Do(func() {
		tp.cancel()
		tp.wg.Wait()

		close(tp.errChan)
	})
}