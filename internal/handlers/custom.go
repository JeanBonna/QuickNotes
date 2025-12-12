package handlers

import "net/http"

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

func (f HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
