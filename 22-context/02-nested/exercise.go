package subtask

import (
	"context"
	"time"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func StartTask(ctx context.Context) (result string, err error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()

	type subResult struct {
		res string
		err error
	}

	resultChan := make(chan subResult, 1)

	go func() {
		res, er := SubTask(timeoutCtx)
		resultChan <- subResult{res: res, err: er}
		close(resultChan)
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()

	case result := <-resultChan:
		if result.err != nil {
			return "", result.err
		}

		return "Main task status: " + result.res, nil
	}
}

func SubTask(ctx context.Context) (result string, err error) {
	select {
	case <-time.After(200 * time.Millisecond):
		return "Subtask completed successfully", nil

	case <-ctx.Done():
		return "", ctx.Err()
	}
}
