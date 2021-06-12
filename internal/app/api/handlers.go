package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KirillNikoda/api/api/internal/app/models"
	"github.com/gorilla/mux"
)

//Auxiliary struct to form messages
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(writer http.ResponseWriter, req *http.Request) {
	//Initialize headers
	initHeaders(writer)
	api.logger.Info("Get All Articles GET /api/v1/articles")
	//Trying to get something from DB
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while Articles.SelectAll : ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Unable to access the database",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) PostArticle(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Article POST /api/v1/articles")
	var article models.Article
	if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	a, err := api.storage.Article().Create(&article)

	if err != nil {
		api.logger.Info("Error while creating new article", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Unable to access the database",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) GetArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Article by ID /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Error while getting article by ID (parsing query parameter)")
		msg := Message{
			StatusCode: 400,
			Message:    "Invalid id",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	article, ok, err := api.storage.Article().FindArticleById(id)

	if err != nil {
		api.logger.Info("Couldn't access the database")
		msg := Message{
			StatusCode: 501,
			Message:    "error while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("Article not found")
		msg := Message{
			StatusCode: 404,
			Message:    "article does not exist",
			IsError:    true,
		}
		writer.WriteHeader(404)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)

}

func (api *API) DeleteArticleById(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Delete Article by ID /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		api.logger.Info("Error while deleting article by ID (parsing query parameter)")
		msg := Message{
			StatusCode: 400,
			Message:    "Invalid id",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	article, err := api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Error while accessing the database")
		msg := Message{
			StatusCode: 501,
			Message:    "database access error",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(article)
}

func (api *API) PostUserRegister(writer http.ResponseWriter, req *http.Request) {
	var user models.User

	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		api.logger.Info("Error while creating new user")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is  invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if _, err := api.storage.User().Create(&user); err != nil {
		api.logger.Info("Error while creating new user")
		msg := Message{
			StatusCode: 501,
			Message:    "error while accessing database",
			IsError:    true,
		}
		writer.WriteHeader(501)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	json.NewEncoder(writer).Encode(&user)
}
