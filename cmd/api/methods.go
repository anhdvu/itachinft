package main

import "net/http"

// logError method logs a specific error along with the ongoing request HTTP method and URL path.
func (svc *service) logError(r *http.Request, err error) {
	svc.logger.Error(err.Error(), "method", r.Method, "url", r.URL.Path)
}

// SendError method is a helper function to send server-side or client-side errors to clients.
func (svc *service) SendError(w http.ResponseWriter, r *http.Request, status int, message any) {
	payload := capsule{"error": message}

	err := SendJSON(w, status, nil, payload)
	if err != nil {
		svc.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// SendServerError method logs an error and sends HTTP status code 500 to clients.
func (svc *service) Send500Error(w http.ResponseWriter, r *http.Request, err error) {
	svc.logError(r, err)

	message := "the server encountered some technical problem"
	svc.SendError(w, r, http.StatusInternalServerError, message)
}

func (svc *service) Send400Error(w http.ResponseWriter, r *http.Request, err error) {
	svc.SendError(w, r, http.StatusBadRequest, err)
}

func (svc *service) Send422Error(w http.ResponseWriter, r *http.Request, errs map[string]any) {
	svc.SendError(w, r, http.StatusUnprocessableEntity, errs)
}
