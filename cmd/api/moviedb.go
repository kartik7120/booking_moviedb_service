package api

import (
	"errors"
	"fmt"
	"time"

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
	result := m.DB.Conn.Find(&venues)

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
		result := m.DB.Conn.Model(&venue).Association("Movies").Find(&venueMovies)

		if result.Error() != "" {
			return movies, 500, errors.New("error fetching movies for venue")
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

func (m *MovieDB) GetMovieShowtimes(movieID uint, venueID uint, movie_format string, date string) ([]models.MovieTimeSlot, int, error) {
	var movie_time_slots []models.MovieTimeSlot

	result := m.DB.Conn.Preload("Venue").Preload("MovieTimeSlots").Preload("MovieTimeSlots.SeatLayout").Preload("MovieTimeSlots.SeatLayout.Seats").Find(&models.Movie{}, movieID)

	if result.Error != nil {
		return movie_time_slots, 500, result.Error
	}

	return movie_time_slots, 200, nil
}

func (m *MovieDB) GetMovieSeatLayout(movieID uint, venueID uint, movie_format string, date string, start_time string) (models.Venue, int, error) {
	var venue models.Venue
	result := m.DB.Conn.Where("id = ?", venueID).Find(&venue)

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

	result := m.DB.Conn.Create(&venue)
	if result.Error != nil {
		return venue, 500, result.Error
	}

	return venue, 200, nil
}

func (m *MovieDB) AddMovie(movie models.Movie, movieTimeSlots []models.MovieTimeSlot, seats []models.SeatMatrix) (models.Movie, int, error) {
	// Validate movie struct
	if err := validate.Struct(movie); err != nil {
		return movie, 400, err
	}

	// Start transaction
	tx := m.DB.Conn.Begin()
	if tx.Error != nil {
		return movie, 500, tx.Error
	}

	// Ensure rollback on panic
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Step 1: Insert Movie
	result := tx.Create(&movie)
	if result.Error != nil {
		tx.Rollback()
		return movie, 500, result.Error
	}

	if result.RowsAffected == 0 {
		tx.Rollback()
		return movie, 500, errors.New("failed to insert movie, no rows affected")
	}

	// Step 2: Insert Venues
	for i := range movie.Venues {
		venue := &movie.Venues[i]

		// Create Venue

		fmt.Println("Inserted Venue ID:", venue.ID) // Debugging

		// Step 3: Insert MovieTimeSlots (from function parameter)
		for j := range movieTimeSlots {
			movieTimeSlots[j].MovieID = movie.ID
			movieTimeSlots[j].VenueID = venue.ID
		}

		if len(movieTimeSlots) > 0 {
			if err := tx.Create(&movieTimeSlots).Error; err != nil {
				tx.Rollback()
				return movie, 500, fmt.Errorf("error inserting time slots: %v", err)
			}
		}

		// Step 4: Insert Seat Matrices (from function parameter)
		for k := range seats {
			seats[k].VenueID = venue.ID
		}

		if len(seats) > 0 {
			if err := tx.Create(&seats).Error; err != nil {
				tx.Rollback()
				return movie, 500, fmt.Errorf("error inserting seat matrix: %v", err)
			}
		}
	}

	// Step 5: Commit transaction
	if err := tx.Commit().Error; err != nil {
		return movie, 500, fmt.Errorf("commit error: %v", err)
	}

	return movie, 200, nil
}

func (m *MovieDB) UpdateMovie(movieID uint, movie models.Movie) (models.Movie, int, error) {
	err := validate.Struct(movie)
	if err != nil {
		return movie, 400, err
	}

	var existingMovie models.Movie
	result := m.DB.Conn.First(&existingMovie, movieID)

	if result.Error != nil {
		return movie, 500, result.Error
	}

	result = m.DB.Conn.Model(&existingMovie).Updates(&movie)

	if result.Error != nil {
		return movie, 500, result.Error
	}

	return movie, 200, nil
}

func (m *MovieDB) DeleteMovie(movieID uint) (int, error) {

	result := m.DB.Conn.Unscoped().Delete(&models.Movie{}, movieID)

	if result.Error != nil {
		return 500, result.Error
	}

	return 200, nil
}

func (m *MovieDB) DeleteVenue(venueId uint) (int, error) {
	var venue models.Venue
	var seats models.SeatMatrix

	result := m.DB.Conn.Unscoped().Where("venue_id = ?", venueId).Delete(&seats)

	if result.Error != nil {
		return 500, result.Error
	}

	result = m.DB.Conn.Unscoped().Delete(&venue, venueId)

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
	result := m.DB.Conn.First(&existingVenue, venueId)

	if result.Error != nil {
		return venue, 500, result.Error
	}

	result = m.DB.Conn.Model(&existingVenue).Updates(&venue)

	if result.Error != nil {
		return venue, 500, result.Error
	}

	return venue, 200, nil
}

func (m *MovieDB) GetVenue(venueId uint) (models.Venue, int, error) {
	var venue models.Venue
	result := m.DB.Conn.First(&venue, venueId)

	if result.Error != nil {
		return venue, 500, result.Error
	}

	return venue, 200, nil
}

// Used to fetch upcoming movies based on the range date given by user,starting from date + 2 weeks to date + 2 weeks + 1 month
func (m *MovieDB) GetUpcomingMovies(date string) ([]models.Movie, int, error) {
	// Parse the input date
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, 400, err
	}

	// Calculate start and end dates
	startDate := d.AddDate(0, 0, 14)
	endDate := d.AddDate(0, 1, 14)

	// Query the database
	var movies []models.Movie
	result := m.DB.Conn.Table("movies").Where("release_date BETWEEN ? AND ?", startDate, endDate).Find(&movies)

	if result.Error != nil {
		return nil, 500, result.Error
	}

	// Return the movies
	return movies, 200, nil
}

func (m *MovieDB) GetNowPlayingMovies() ([]models.Movie, int, error) {
	today := time.Now().Truncate(24 * time.Hour)

	var movies []models.Movie
	err := m.DB.Conn.
		Joins("JOIN movie_time_slots mts ON mts.movie_id = movies.id").
		Where("movies.release_date <= ?", today).
		Where("DATE(mts.date) = ?", today).
		Group("movies.id").
		Find(&movies).Error

	if err != nil {
		return nil, 500, err
	}
	return movies, 200, nil
}
