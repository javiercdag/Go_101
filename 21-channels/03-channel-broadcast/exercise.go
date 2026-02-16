package channelbroadcaster

import (
	"context"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func channelBroadcast(ctx context.Context, input <-chan any, outputs []chan<- any) {
	go safeBridge(ctx, input, outputs)
}

func safeBridge(ctx context.Context, in <-chan any, out []chan<- any) {

	defer func() {
		for _, output := range out {
			close(output)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return

		case msg, ok := <-in:
			if !ok {
				return
			}

			for _, output := range out {
				select {
				case output <- msg:
					// res
				case <-ctx.Done():
					return
				}
			}
		}
	}
}
