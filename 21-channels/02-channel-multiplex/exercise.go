package channelmultiplexer

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func channelMultiplex(ctx context.Context, inputs []chan any) chan any {
	returnChan := make(chan any)

	var eg errgroup.Group

	for w := 0; w < len(inputs); w++ {
		eg.Go(func() error {
			return safeBridge(ctx, inputs[w], returnChan)
		})
	}

	go func() {
		_ = eg.Wait()
		close(returnChan)
	}()

	return returnChan
}

func safeBridge(ctx context.Context, in <-chan any, out chan<- any) error {
	for {
		select {
		case msg, ok := <-in:
			if !ok {
				return nil
			}

			select {
			case out <- msg:
			case <-ctx.Done():
				return nil
			}

		case <-ctx.Done():
			return nil
		}
	}
}
