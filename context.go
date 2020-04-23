package webdav

import (
	"context"
	"net/http"
)

type CtxKey int

const (
	CtxError = iota + 1
	CtxValue
)

func setError(r *http.Request, errStr string) {
	ctx := context.WithValue(r.Context(), CtxError, errStr)
	*r = *r.WithContext(ctx)
}

func setValue(r *http.Request, value interface{}) {
	ctx := context.WithValue(r.Context(), CtxValue, value)
	*r = *r.WithContext(ctx)
}

type pathValue struct {
	Path string `json:"path"`
}

type srcDstValue struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type idValue struct {
	ID string `json:"id"`
}
