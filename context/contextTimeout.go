package context

import (
	"context"
	"time"
)

func longRunningMethod() (int, error) {

	time.Sleep(time.Microsecond * 500)

	return 200, nil
}

func MockCaller(ctx context.Context) (int, error) {

	val, err := longRunningMethod()

	return val, err
}
