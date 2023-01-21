package data

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

var (
	ErrEditConflict = errors.New("edit conflict")

	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Movies interface {
		Insert(movie *Movie, r *http.Request) error
		Get(id int64, r *http.Request) (*Movie, error)
		Update(movie *Movie, r *http.Request) error
		Delete(id int64, r *http.Request) error
		GetAll(title string, genres []string, filters Filters, r *http.Request) ([]*Movie, Metadata, error)
	}
	Users interface {
		Insert(user *User) error
		GetByEmail(email string) (*User, error)
		Update(user *User) error
	}
}

func NewModels(pool *pgxpool.Pool) Models {
	return Models{
		Movies: MovieModel{pool: pool},
		Users:  UserModel{pool: pool},
	}
}
