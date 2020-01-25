package app

import "github.com/gorilla/mux"

import "net/http"

import "time"

import "github.com/waytkheming/bookstore-items-api/client/elasticsearch"

var (
	router = mux.NewRouter()
)

func StartApp() {
	elasticsearch.Init()

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
