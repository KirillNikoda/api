package api

import (
	"net/http"

	"github.com/KirillNikoda/api/api/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// API Base API server instance description
type API struct {
	//UNEXPORTED FIELD!!!
	config *Config
	logger *logrus.Logger
	router *mux.Router
	//Adding field in order to work with storage
	storage *storage.Storage
}

// New API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config:  config,
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		storage: &storage.Storage{},
	}
}

//Start http server/configure loggers, router, database connection and e.t.c. ...
func (api *API) Start() error {
	//Trying to configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	// Confirmation that logger was configured successfully
	api.logger.Info("starting api server at port:", api.config.BindArr)

	//Configuring Router
	api.configureRouterField()

	//Configuring Storage
	if err := api.configureStorageField(); err != nil {
		return err
	}

	// At the stage of valid completion, we are gonna start http - server
	return http.ListenAndServe(api.config.BindArr, api.router)
}
