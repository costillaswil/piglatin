package main

import (
	"fmt"
	"net/http"
	"test/todos/api"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func serve(api *api.API) {
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	router := mux.NewRouter()
	api.Init(router.PathPrefix("/api").Subrouter())

	s := http.Server{
		Addr:        fmt.Sprintf(":%d", 8006),
		Handler:     cors(router),
		ReadTimeout: 60 * time.Second,
	}

	logrus.Println(fmt.Sprintf("serving api at http://127.0.0.1:%d", 8006))
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}
}
