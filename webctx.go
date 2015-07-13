package webctx

import (
	"io"
	"net/http"
)

type WebCtx struct {
	orig io.ReadCloser
	data map[interface{}]interface{}
}

func GetCtx(r *http.Request) *WebCtx {
	ctx, ok := r.Body.(*WebCtx)
	if !ok {
		ctx = &WebCtx{orig: r.Body, data: make(map[interface{}]interface{})}
		r.Body = ctx
	}
	return ctx
}

func Get(r *http.Request, key interface{}) interface{} {
	return GetCtx(r).Get(key)
}
func Set(r *http.Request, key interface{}, val interface{}) {
	GetCtx(r).Set(key, val)
}

func (w *WebCtx) Get(key interface{}) interface{} {
	return w.data[key]
}

func (w *WebCtx) Set(key, val interface{}) {
	w.data[key] = val
}

func (w *WebCtx) Close() error {
	return w.orig.Close()
}

func (w *WebCtx) Read(p []byte) (int, error) {
	return w.orig.Read(p)
}
