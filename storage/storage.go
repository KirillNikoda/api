package storage

import (
	"database/sql"
	"log"

	// Just to invoke Init() function of pq package
	_ "github.com/lib/pq"
)

//Instance of storage
type Storage struct {
	config *Config
	//Database FileDescriptor
	db *sql.DB
	//Subfield for repo interfacing (model User)
	userRepository *UserRepository
	//Subfield for repo interfacing (model Article)
	articleRepository *ArticleRepository
}

//Storage Constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

//Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	storage.db = db
	log.Println("database connection created successfully")
	return nil
}

//Close connection method
func (storage *Storage) Close() {
	storage.db.Close()
}

//Public Repo for User
func (storage *Storage) User() *UserRepository {
	if storage.userRepository != nil {
		return storage.userRepository
	}
	storage.userRepository = &UserRepository{
		storage: storage,
	}
	return nil
}

//Public Repo for ARticle
func (storage *Storage) Article() *ArticleRepository {
	if storage.articleRepository != nil {
		return storage.articleRepository
	}
	storage.articleRepository = &ArticleRepository{
		storage: storage,
	}
	return nil
}
