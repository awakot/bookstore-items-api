package app

import "github.com/gorilla/mux"

import "net/http"

import "time"

var (
	router = mux.NewRouter()
)

func StartApp() {
	mapUrls()
	srv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
