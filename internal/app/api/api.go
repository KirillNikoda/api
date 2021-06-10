package api

// API Base API server instance description
type API struct {
	//UNEXPORTED FIELD!!!
	config *Config
}

// New API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

//Start http server/configure loggers, router, database connection and e.t.c. ...
func (api *API) Start() error {
	return nil
}