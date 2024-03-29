package main

import (
	"fmt"
	"net/http"
)

// RecoverPanic is a middleware that recovers panics during runtime and sends server-side error
// to inform client.
// *NOTE: Chi provides a Recovers middlware which is more comprehensive. But either works fine.
func (svc *service) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				svc.Send500Error(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
