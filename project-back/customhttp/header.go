package customhttp

import (
	"github.com/gorilla/mux"
	"net/http"
)

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, session")
}

func HandleFunc(router *mux.Router, path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		setHeader(w)
		f(w, r)
	})
}
