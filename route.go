package main

import (
	"net/http"

	"github.com/0xma12k/graylog-line-notify-gateway/controllers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var router = mux.NewRouter().StrictSlash(true)

// Router ...
func Router() *mux.Router {
	router.Path("/line").Methods("POST").HandlerFunc(controllers.Webhook)
	router.Use(loggingMiddleware)
	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logrus.Debug(r.Method + " " + r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
