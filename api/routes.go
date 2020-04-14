package api

import (
	"fmt"
	"net/http"
	service "pigLatin/services"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type API struct {
	piglatinService *service.PiglatinService
	router          *mux.Router
}

//New creates new API instance
func New(piglatinService *service.PiglatinService, router *mux.Router) *API {
	return &API{
		piglatinService: piglatinService,
		router:          router,
	}
}

//init routing for exposed API end-points.
func (s *API) init() {
	s.router.HandleFunc("/piglatin", s.handlePigLatinTranslation()).Methods("POST")
}

//Start setup server configurations and initialization.
func (s *API) Start() {

	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	s.init()

	server := &http.Server{
		Addr:        fmt.Sprintf(":%d", 8006),
		Handler:     cors(s.router),
		ReadTimeout: 60 * time.Second,
	}

	logrus.Println(fmt.Sprintf("serving api at http://127.0.0.1:%d", 8006))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.Error(err)
	}

}
