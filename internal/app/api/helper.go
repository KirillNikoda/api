package api

import (
	"net/http"

	"github.com/KirillNikoda/api/api/storage"
	"github.com/sirupsen/logrus"
)

//Trying to configure our API instance (logger field of API instance)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// Trying to configure Router (router field of API instance)
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello! This is rest api!"))
	})
}

//Trying to configure our storage (storage field of API instance)
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	//Trying to set connection, if not possible -> returns error
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
