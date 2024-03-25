package main

import (
	"fmt"
	"net/http"
)

// These are custom default handlers for common status codes such as 404 and 405.
// They are used to customize Chi router behavior.

func (svc *service) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	message := "the requested resources could not be found"
	svc.SendError(w, r, http.StatusNotFound, message)
}

func (svc *service) MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	svc.SendError(w, r, http.StatusMethodNotAllowed, message)
}
