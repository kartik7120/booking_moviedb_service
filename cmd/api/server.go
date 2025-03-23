package api

import (
	"context"
	"strconv"
	"time"

	moviedb "github.com/kartik7120/booking_moviedb_service/cmd/grpcServer"
	"github.com/kartik7120/booking_moviedb_service/cmd/models"
	log "github.com/sirupsen/logrus"
)

// var validate *validator.Validate

type MoviedbService struct {
	moviedb.UnimplementedMovieDBServiceServer
	MovieDB *MovieDB
}

func NewMoviedbService() *MoviedbService {
	// validate = validator.New()
	return &MoviedbService{}
}

func (m *MoviedbService) AddMovie(ctx context.Context, in *moviedb.Movie) (*moviedb.MovieResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// if ctx.Done() != nil {
	// 	return &moviedb.MovieResponse{
	// 		Status:  408,
	// 		Message: "Context was cancelled",
	// 		Error:   "",
	// 	}, ctx.Err()
	// }

	castAndCrew := make([]models.CastAndCrew, 0)

	for _, cc := range in.CastCrew {
		c := models.CastAndCrew{
			Type:      string(cc.Type),
			Name:      cc.Name,
			Character: cc.CharacterName,
			PhotoURL:  cc.Photourl,
		}
		castAndCrew = append(castAndCrew, c)
	}

	venues := make([]models.Venue, 0)

	releaseDate, err := time.Parse("2006-01-02", in.ReleaseDate)

	if err != nil {
		return &moviedb.MovieResponse{
			Status:  400,
			Message: "error parsing release date",
			Error:   err.Error(),
		}, err
	}

	movieTimeSlots := make([]models.MovieTimeSlot, 0)
	seats := make([]models.SeatMatrix, 0)

	for _, v := range in.Venues {

		for _, slot := range v.MovieTimeSlots {
			timestring, err := time.Parse("2006-01-02", slot.Date)

			if err != nil {
				return &moviedb.MovieResponse{
					Status:  400,
					Message: "error parsing slot date",
					Error:   err.Error(),
				}, err
			}

			timeSlot := models.MovieTimeSlot{
				StartTime:   slot.StartTime,
				EndTime:     slot.EndTime,
				Duration:    int(slot.Duration),
				Date:        timestring,
				MovieFormat: slot.MovieFormat.String(),
			}

			movieTimeSlots = append(movieTimeSlots, timeSlot)
		}

		for _, seat := range v.Seats {

			seat := models.SeatMatrix{
				SeatNumber: seat.SeatNumber,
				IsBooked:   seat.IsBooked,
				Type:       seat.Type.String(),
				Price:      int(seat.Price),
				Row:        int(seat.Row),
				Column:     int(seat.Column),
			}

			seats = append(seats, seat)
		}

		venue := models.Venue{
			Name:         v.Name,
			Type:         string(v.Type),
			Address:      v.Address,
			Rows:         int(v.Rows),
			Columns:      int(v.Columns),
			ScreenNumber: int(v.ScreenNumber),
			Longitude:    float64(v.Longitude),
			Latitude:     float64(v.Latitude),
		}
		venues = append(venues, venue)
	}

	movie := models.Movie{
		Title:           in.Title,
		Description:     in.Description,
		Duration:        int(in.Duration),
		Language:        in.Language,
		Type:            in.Type,
		CastCrew:        castAndCrew,
		PosterURL:       in.PosterUrl,
		TrailerURL:      in.TrailerUrl,
		ReleaseDate:     releaseDate,
		MovieResolution: in.MovieResolution,
		Venues:          venues,
	}

	m.MovieDB.AddMovie(movie, movieTimeSlots, seats)

	return &moviedb.MovieResponse{
		Status:  200,
		Message: "Movie added successfully",
		Movie:   in,
		Error:   "",
	}, nil
}

func (m *MoviedbService) GetMovie(ctx context.Context, in *moviedb.MovieRequest) (*moviedb.MovieResponse, error) {

	log.Info("Starting gRPC GetMovie function call")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// if ctx.Done() != nil {
	// 	return &moviedb.MovieResponse{
	// 		Status:  408,
	// 		Message: "Context was cancelled",
	// 		Error:   "",
	// 	}, ctx.Err()
	// }

	movieID, err := strconv.ParseUint(in.Movieid, 10, 32)

	if err != nil {
		log.Error("Invalid movie id", err)
		return &moviedb.MovieResponse{
			Status:  400,
			Message: "Invalid movie ID",
			Error:   err.Error(),
		}, nil
	}

	movie, status, err := m.MovieDB.GetMovieDetails(uint(movieID))

	if err != nil {
		log.Info("error calling get movie details function", err)
		return &moviedb.MovieResponse{
			Status:  int32(status),
			Message: "Movie not found",
			Error:   err.Error(),
		}, nil
	}

	castCrew := make([]*moviedb.CastAndCrew, 0)

	for _, cc := range movie.CastCrew {
		c := &moviedb.CastAndCrew{
			Type:          moviedb.CastAndCrewType(moviedb.CastAndCrewType_value[cc.Type]),
			Name:          cc.Name,
			CharacterName: cc.Character,
			Photourl:      cc.PhotoURL,
		}
		castCrew = append(castCrew, c)
	}

	return &moviedb.MovieResponse{
		Status:  200,
		Message: "Sucess",
		Movie: &moviedb.Movie{
			Title:           movie.Title,
			Description:     movie.Description,
			Duration:        int32(movie.Duration),
			Language:        movie.Language,
			Type:            movie.Type,
			CastCrew:        castCrew,
			PosterUrl:       movie.PosterURL,
			ReleaseDate:     movie.ReleaseDate.GoString(),
			TrailerUrl:      movie.TrailerURL,
			MovieResolution: movie.MovieResolution,
			Movieid:         string(movie.ID),
		},
	}, nil
}

func (m *MoviedbService) UpdateMovie(ctx context.Context, in *moviedb.Movie) (*moviedb.MovieResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// if ctx.Done() != nil {
	// 	return &moviedb.MovieResponse{
	// 		Status:  408,
	// 		Message: "Context was cancelled",
	// 		Error:   "",
	// 	}, ctx.Err()
	// }

	castAndCrew := make([]models.CastAndCrew, 0)

	for _, cc := range in.CastCrew {
		c := models.CastAndCrew{
			Type:      string(cc.Type),
			Name:      cc.Name,
			Character: cc.CharacterName,
			PhotoURL:  cc.Photourl,
		}
		castAndCrew = append(castAndCrew, c)
	}

	releaseDate, err := time.Parse("2006-01-02", in.ReleaseDate)

	if err != nil {
		return &moviedb.MovieResponse{
			Status:  400,
			Message: "error parsing release date",
			Error:   err.Error(),
		}, err
	}

	movie := models.Movie{
		Title:           in.Title,
		Description:     in.Description,
		Duration:        int(in.Duration),
		Language:        in.Language,
		Type:            in.Type,
		CastCrew:        castAndCrew,
		PosterURL:       in.PosterUrl,
		TrailerURL:      in.TrailerUrl,
		ReleaseDate:     releaseDate,
		MovieResolution: in.MovieResolution,
	}

	movieID, err := strconv.ParseUint(in.Movieid, 10, 32)

	if err != nil {
		return &moviedb.MovieResponse{
			Status:  400,
			Message: "Invalid movie ID",
			Error:   err.Error(),
		}, nil
	}

	m.MovieDB.UpdateMovie(uint(movieID), movie)

	return &moviedb.MovieResponse{
		Status:  200,
		Message: "Movie updated successfully",
		Movie:   in,
		Error:   "",
	}, nil

}

func (m *MoviedbService) DeleteMovie(ctx context.Context, in *moviedb.MovieRequest) (*moviedb.MovieResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// if ctx.Done() != nil {
	// 	return &moviedb.MovieResponse{
	// 		Status:  408,
	// 		Message: "Context was cancelled",
	// 		Error:   "",
	// 	}, ctx.Err()
	// }

	movieID, err := strconv.ParseUint(in.Movieid, 10, 32)

	if err != nil {
		return &moviedb.MovieResponse{
			Status:  400,
			Message: "Invalid movie ID",
			Error:   err.Error(),
		}, nil
	}

	status, err := m.MovieDB.DeleteMovie(uint(movieID))

	if err != nil {
		return &moviedb.MovieResponse{
			Status:  int32(status),
			Message: "Movie not found",
			Error:   "",
		}, nil
	}

	return &moviedb.MovieResponse{
		Status:  200,
		Message: "Movie deleted successfully",
		Error:   "",
	}, nil
}

func (m *MoviedbService) DeleteVenue(ctx context.Context, in *moviedb.MovieRequest) (*moviedb.MovieResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// if ctx.Done() != nil {
	// 	return &moviedb.MovieResponse{
	// 		Status:  408,
	// 		Message: "Context was cancelled",
	// 		Error:   ctx.Err().Error(),
	// 	}, ctx.Err()
	// }

	Venueid, err := strconv.ParseUint(in.Venueid, 10, 32)

	if err != nil {
		return &moviedb.MovieResponse{
			Status:  400,
			Message: "Invalid venue ID",
			Error:   err.Error(),
		}, err
	}

	status, err := m.MovieDB.DeleteVenue(uint(Venueid))

	if status != 200 || err != nil {
		return &moviedb.MovieResponse{
			Status:  int32(status),
			Message: "error deleting venue",
			Error:   err.Error(),
		}, err
	}

	return &moviedb.MovieResponse{
		Status:  int32(status),
		Message: "Venue deleted",
	}, nil
}

func (m *MoviedbService) UpdateVenue(ctx context.Context, in *moviedb.Venue) (*moviedb.VenueResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// if ctx.Done() != nil {
	// 	return &moviedb.VenueResponse{
	// 		Status:  500,
	// 		Message: "context is already cancelled",
	// 	}, ctx.Err()
	// }

	v := models.Venue{
		Name:         in.Name,
		Type:         in.Type.String(),
		Address:      in.Address,
		Rows:         int(in.Rows),
		Columns:      int(in.Columns),
		ScreenNumber: int(in.ScreenNumber),
		Longitude:    float64(in.Longitude),
		Latitude:     float64(in.Latitude),
	}

	movieFormatSupported := make([]string, 0)

	for _, val := range v.MovieFormatSupported {
		movieFormatSupported = append(movieFormatSupported, val)
	}

	languageSupported := make([]string, 0)

	for _, val := range v.LanguagesSupported {
		languageSupported = append(languageSupported, val)
	}

	if len(movieFormatSupported) > 0 {
		v.MovieFormatSupported = movieFormatSupported
	}

	if len(languageSupported) > 0 {
		v.LanguagesSupported = languageSupported
	}

	_, status, err := m.MovieDB.UpdateVenue(uint(in.Id), v)

	if status != 200 || err != nil {
		return &moviedb.VenueResponse{
			Status:  int32(status),
			Message: "error updating venue",
			Error:   err.Error(),
		}, err
	}

	return &moviedb.VenueResponse{
		Status:  200,
		Message: "Updated venue",
	}, nil
}

func (m *MoviedbService) AddVenue(ctx context.Context, in *moviedb.Venue) (*moviedb.VenueResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// if ctx.Done() != nil {
	// 	return &moviedb.VenueResponse{
	// 		Status:  500,
	// 		Message: "context is already cancelled",
	// 	}, nil
	// }

	venue := models.Venue{
		Name:         in.Name,
		Type:         in.Type.String(),
		Address:      in.Address,
		Rows:         int(in.Rows),
		Columns:      int(in.Columns),
		ScreenNumber: int(in.ScreenNumber),
		Longitude:    float64(in.Longitude),
		Latitude:     float64(in.Latitude),
	}

	movieFormatSupported := make([]string, 0)

	movieFormatSupported = append(movieFormatSupported, in.MovieFormatSupported...)

	if len(movieFormatSupported) > 0 {
		venue.MovieFormatSupported = movieFormatSupported
	}

	languageSupported := make([]string, 0)

	languageSupported = append(languageSupported, in.LanguageSupported...)

	if len(languageSupported) > 0 {
		venue.LanguagesSupported = languageSupported
	}

	seats := make([]models.SeatMatrix, 0)

	for _, val := range in.Seats {
		seat := models.SeatMatrix{
			SeatNumber: val.SeatNumber,
			IsBooked:   val.IsBooked,
			Type:       val.Type.String(),
			Price:      int(val.Price),
			Row:        int(val.Row),
			Column:     int(val.Column),
		}

		seats = append(seats, seat)
	}

	if len(seats) > 0 {
		venue.Seats = seats
	}

	_, status, err := m.MovieDB.AddVenue(venue)

	if status != 200 || err != nil {
		return &moviedb.VenueResponse{
			Status:  int32(status),
			Message: "error adding a new venue",
			Error:   err.Error(),
		}, err
	}

	return &moviedb.VenueResponse{
		Status:  int32(status),
		Message: "added a new venue",
	}, nil
}

func (m *MoviedbService) GetVenue(ctx context.Context, in *moviedb.MovieRequest) (*moviedb.VenueResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	// if ctx.Done() != nil {
	// 	return &moviedb.VenueResponse{
	// 		Status:  500,
	// 		Message: "context is already cancelled",
	// 	}, ctx.Err()
	// }

	Venueid, err := strconv.ParseUint(in.Venueid, 10, 32)

	if err != nil {
		return &moviedb.VenueResponse{
			Status: 500,
			Error:  err.Error(),
		}, err
	}

	venue, status, err := m.MovieDB.GetVenue(uint(Venueid))

	if status != 200 || err != nil {
		return &moviedb.VenueResponse{
			Status:  int32(status),
			Message: "error getting venue",
		}, err
	}

	return &moviedb.VenueResponse{
		Status:  200,
		Message: "success",
		Venue: &moviedb.Venue{
			Name:    venue.Name,
			Address: venue.Address,
			Rows:    int32(venue.Rows),
			Columns: int32(venue.Columns),
		},
	}, nil
}
