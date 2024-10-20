package detachutil

import (
	"context"
	"time"
)

func Detach(ctx context.Context) context.Context {
	return detachedContext{ctx}
}

type detachedContext struct {
	parent context.Context
}

func (dc detachedContext) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (dc detachedContext) Done() <-chan struct{} {
	return nil
}

func (dc detachedContext) Err() error {
	return nil
}

func (dc detachedContext) Value(key interface{}) interface{} {
	return dc.parent.Value(key)
}
