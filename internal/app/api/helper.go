package api

import (
	"github.com/KirillNikoda/api/api/storage"
	"github.com/sirupsen/logrus"
)

var (
	prefix string = "/api/v1"
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
	a.router.HandleFunc(prefix+"/articles", a.GetAllArticles).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.GetArticleById).Methods("GET")
	a.router.HandleFunc(prefix+"/articles/{id}", a.DeleteArticleById).Methods("DELETE")
	a.router.HandleFunc(prefix+"/articles", a.PostArticle).Methods("POST")
	a.router.HandleFunc(prefix+"/user/register", PostUserRegister).Methods("POST")
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
