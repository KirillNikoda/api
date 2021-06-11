package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// API Base API server instance description
type API struct {
	//UNEXPORTED FIELD!!!
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
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

	// At the stage of valid completion, we are gonna start http - server
	return nil
}
