package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/kartik7120/booking_moviedb_service/cmd/helper"
)

type MovieDB struct {
	DB helper.DBConfig
}

var validate *validator.Validate

func NewMovieDB() *MovieDB {
	validate = validator.New()
	return &MovieDB{}
}

func (m MovieDB) GetCurrentMovies() {

}
