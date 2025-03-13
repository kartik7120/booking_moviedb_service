package api

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/kartik7120/booking_moviedb_service/cmd/helper"
	"github.com/kartik7120/booking_moviedb_service/cmd/models"
)

type MovieDB struct {
	DB helper.DBConfig
}

var validate *validator.Validate

func NewMovieDB() *MovieDB {
	validate = validator.New()
	return &MovieDB{}
}

func (m *MovieDB) GetCurrentMovies(
	latitude float64,
	longitude float64,
) ([]models.Movie, int, error) {
	var movies []models.Movie
	var venues []models.Venue

	// Fetch all venues from the database
	result := m.DB.Conn.Table("moviedb").Find(&venues)

	if result.Error != nil {
		return movies, 500, result.Error
	}

	// Filter venues within 30km radius
	var nearbyVenues []models.Venue
	for _, venue := range venues {
		distance := helper.Haversine(latitude, longitude, venue.Latitude, venue.Longitude)
		if distance <= 30 {
			nearbyVenues = append(nearbyVenues, venue)
		}
	}

	// Fetch movies for the nearby venues
	for _, venue := range nearbyVenues {
		var venueMovies []models.Movie
		result := m.DB.Conn.Table("moviedb").Model(&venue).Association("Movies").Find(&venueMovies)

		if result.Error != nil {
			return movies, 500, errors.New("Error fetching movies for venue")
		}

		movies = append(movies, venueMovies...)
	}

	return movies, 200, nil
}

func (m *MovieDB) GetMovieDetails(movieID uint) (models.Movie, int, error) {
	var movie models.Movie
	result := m.DB.Conn.Preload("CastCrew").First(&movie, movieID)

	if result.Error != nil {
		return movie, 500, result.Error
	}
	return movie, 200, nil
}

func (m *MovieDB) GetMovieShowtimes(movieID uint, venueID uint, movie_format string, date string) ([]models.Movie_time_slot, int, error) {
	var movie_time_slots []models.Movie_time_slot
	result := m.DB.Conn.Table("moviedb").Where("movie_id = ? AND venue_id = ? AND movie_format = ? AND date = ?", movieID, venueID, movie_format, date).Find(&movie_time_slots)

	if result.Error != nil {
		return movie_time_slots, 500, result.Error
	}

	return movie_time_slots, 200, nil
}

func (m *MovieDB) GetMovieSeatLayout(movieID uint, venueID uint, movie_format string, date string, start_time string) (models.Venue, int, error) {
	var venue models.Venue
	result := m.DB.Conn.Table("moviedb").Where("id = ?", venueID).Find(&venue)

	if result.Error != nil {
		return venue, 500, result.Error
	}

	return venue, 200, nil
}

func (m *MovieDB) AddVenue(venue models.Venue) (models.Venue, int, error) {
	err := validate.Struct(venue)
	if err != nil {
		return venue, 400, err
	}

	result := m.DB.Conn.Table("moviedb").Create(&venue)
	if result.Error != nil {
		return venue, 500, result.Error
	}

	return venue, 200, nil
}

func (m *MovieDB) AddMovie(movie models.Movie) (models.Movie, int, error) {
	err := validate.Struct(movie)
	if err != nil {
		return movie, 400, err
	}

	result := m.DB.Conn.Table("moviedb").Create(&movie)

	if result.Error != nil {
		return movie, 500, result.Error
	}

	return movie, 200, nil
}

func (m *MovieDB) UpdateMovie(movieID uint, movie models.Movie) (models.Movie, int, error) {
	err := validate.Struct(movie)
	if err != nil {
		return movie, 400, err
	}

	var existingMovie models.Movie
	result := m.DB.Conn.Table("moviedb").First(&existingMovie, movieID)

	if result.Error != nil {
		return movie, 500, result.Error
	}

	result = m.DB.Conn.Table("moviedb").Model(&existingMovie).Updates(&movie)

	if result.Error != nil {
		return movie, 500, result.Error
	}

	return movie, 200, nil
}

func (m *MovieDB) DeleteMovie(movieID uint) (int, error) {
	var movie models.Movie
	result := m.DB.Conn.Table("moviedb").First(&movie, movieID)

	if result.Error != nil {
		return 500, result.Error
	}

	result = m.DB.Conn.Table("moviedb").Delete(&movie)

	if result.Error != nil {
		return 500, result.Error
	}

	return 200, nil
}

func (m *MovieDB) DeleteVenue(venueId uint) (int, error) {
	var venue models.Venue
	result := m.DB.Conn.Table("moviedb").First(&venue, venueId)

	if result.Error != nil {
		return 500, result.Error
	}

	result = m.DB.Conn.Table("moviedb").Delete(&venue)

	if result.Error != nil {
		return 500, result.Error
	}

	return 200, nil
}

func (m *MovieDB) UpdateVenue(venueId uint, venue models.Venue) (models.Venue, int, error) {
	err := validate.Struct(venue)
	if err != nil {
		return venue, 400, err
	}

	var existingVenue models.Venue
	result := m.DB.Conn.Table("moviedb").First(&existingVenue, venueId)

	if result.Error != nil {
		return venue, 500, result.Error
	}

	result = m.DB.Conn.Table("moviedb").Model(&existingVenue).Updates(&venue)

	if result.Error != nil {
		return venue, 500, result.Error
	}

	return venue, 200, nil
}

func (m *MovieDB) GetVenue(venueId uint) (models.Venue, int, error) {
	var venue models.Venue
	result := m.DB.Conn.Table("moviedb").First(&venue, venueId)

	if result.Error != nil {
		return venue, 500, result.Error
	}

	return venue, 200, nil
}
