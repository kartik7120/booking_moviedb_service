package tests

import (
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/kartik7120/booking_moviedb_service/cmd/api"
	"github.com/kartik7120/booking_moviedb_service/cmd/helper"
	"github.com/kartik7120/booking_moviedb_service/cmd/models"
	"github.com/lib/pq"
)

func TestMovieDB(t *testing.T) {
	t.Run("Add movie to database", func(t *testing.T) {
		err := godotenv.Load()

		if err != nil {
			t.Errorf("Could not load .env file")
		}

		m := api.NewMovieDB()

		// connect to database

		conn, err := helper.ConnectToDB()

		if err != nil {
			t.Errorf("unable to connect to database")
		}

		m.DB.Conn = conn

		releaseDate, err := time.Parse("2006-01-02", "2022-03-04")

		if err != nil {
			t.Errorf("error parsing release date")
			return
		}

		movieTimeSlotDate, err := time.Parse("2006-01-02", "2025-10-10")

		if err != nil {
			t.Errorf("error parsing movie time slot date")
			return
		}

		// add movie to database

		// movie := models.Movie{
		// 	Title:       "The Batman",
		// 	Description: "The Batman is an upcoming American superhero film based on the DC Comics character Batman.",
		// 	ReleaseDate: "2022-03-04",
		// 	PosterURL:   "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4q2hz2m8hubgvijz8Ez0T2Os2Yv.jpg",
		// 	Duration:    10560,               // should be in seconds,
		// 	Language:    []string{"English"}, // Correctly formatted array of strings
		// 	Type:        []string{"Action", "Crime", "Drama"},
		// 	CastCrew: []models.CastAndCrew{
		// 		{
		// 			Name:      "Robert Pattinson",
		// 			Type:      "Actor",
		// 			Character: "Bruce Wayne / Batman",
		// 			PhotoURL:  "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4q2hz2m8hubgvijz8Ez0T2Os2Yv.jpg",
		// 		},
		// 		{
		// 			Name:      "Zoë Kravitz",
		// 			Type:      "Actor",
		// 			Character: "Selina Kyle / Catwoman",
		// 			PhotoURL:  "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4q2hz2m8hubgvijz8Ez0T2Os2Yv.jpg",
		// 		},
		// 		{
		// 			Name:      "Paul Dano",
		// 			Type:      "Actor",
		// 			Character: "Edward Nashton / The Riddler",
		// 			PhotoURL:  "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4q2hz2m8hubgvijz8Ez0T2Os2Yv.jpg",
		// 		},
		// 		{
		// 			Name:      "Jeffrey Wright",
		// 			Type:      "Actor",
		// 			Character: "James Gordon",
		// 			PhotoURL:  "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4q2hz2m8hubgvijz8Ez0T2Os2Yv.jpg",
		// 		},
		// 		{
		// 			Name:      "Andy Serkis",
		// 			Type:      "Actor",
		// 			Character: "Alfred Pennyworth",
		// 			PhotoURL:  "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4q2hz2m8hubgvijz8Ez0T2Os2Yv.jpg",
		// 		},
		// 		{
		// 			Name:      "Colin Farrell",
		// 			Type:      "Actor",
		// 			Character: "Oswald Cobblepot / The Penguin",
		// 			PhotoURL:  "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4q2hz2m8hubgvijz8Ez0T2Os2Yv.jpg",
		// 		},
		// 	},
		// 	TrailerURL:      "https://www.youtube.com/watch?v=IhVf_3TeTQE",
		// 	MovieResolution: []string{"4K", "2K", "HD"},
		// 	Venues: []models.Venue{
		// 		{
		// 			Name:      "PVR Cinemas",
		// 			Type:      "Multiplex",
		// 			Address:   "PVR Plaza, Connaught Place, New Delhi, Delhi 110001",
		// 			Latitude:  28.6315,
		// 			Longitude: 77.2167,
		// 			Rows:      10,
		// 			Columns:   10,
		// 			Seats: []models.SeatMatrix{
		// 				{
		// 					Row:        1,
		// 					Column:     1,
		// 					Price:      200,
		// 					SeatNumber: "A1",
		// 					IsBooked:   false,
		// 					Type:       "Regular",
		// 				},
		// 			},
		// 		},
		// 	},
		// }

		movie := models.Movie{
			Title:           "Blade Runner 2049",
			Description:     "A young blade runner's discovery of a long-buried secret leads him to track down former blade runner Rick Deckard, who's been missing for thirty years.",
			ReleaseDate:     releaseDate, // Release Date: October 6, 2017
			PosterURL:       "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/gajva2L0rPYkEWjzgFlBXCAVBE5.jpg",
			Duration:        164, // 2 hours 44 minutes
			Language:        pq.StringArray([]string{"English", "Spanish", "French"}),
			Type:            pq.StringArray([]string{"Science Fiction", "Drama", "Thriller"}),
			MovieResolution: pq.StringArray([]string{"4K", "1080p", "720p"}),
			Venues: []models.Venue{
				{
					Name:      "IMAX - Grand Cinema",
					Type:      "IMAX",
					Address:   "Grand Mall, Downtown, Los Angeles, CA 90012",
					Latitude:  34.0522,
					Longitude: -118.2437,
					Rows:      15,
					Columns:   20,
					// Seats: []models.SeatMatrix{
					// 	{Row: 1, Column: 1, Price: 1200, SeatNumber: "A1", IsBooked: false, Type: "Platinum"},
					// 	{Row: 1, Column: 2, Price: 1200, SeatNumber: "A2", IsBooked: false, Type: "Platinum"},
					// 	{Row: 2, Column: 1, Price: 1000, SeatNumber: "B1", IsBooked: true, Type: "Gold"},
					// 	{Row: 2, Column: 2, Price: 1000, SeatNumber: "B2", IsBooked: false, Type: "Gold"},
					// },
					LanguagesSupported:   pq.StringArray([]string{"English", "Spanish"}),
					ScreenNumber:         1,
					MovieFormatSupported: pq.StringArray([]string{"IMAX", "3D", "2D"}),
					// MovieTimeSlots: []models.MovieTimeSlot{
					// 	{
					// 		StartTime:   "18:00", // 6:00 PM
					// 		EndTime:     "21:00", // 9:00 PM
					// 		Duration:    10800,   // 3 hours
					// 		Date:        time.Date(2025, 3, 20, 0, 0, 0, 0, time.UTC),
					// 		MovieFormat: "IMAX",
					// 	},
					// 	{
					// 		StartTime:   "21:30", // 9:30 PM
					// 		EndTime:     "00:30", // 12:30 AM
					// 		Duration:    10800,   // 3 hours
					// 		Date:        time.Date(2025, 3, 20, 0, 0, 0, 0, time.UTC),
					// 		MovieFormat: "3D",
					// 	},
					// },
				},
			},
		}

		_, status, err := m.AddMovie(movie, []models.MovieTimeSlot{
			{
				StartTime:   "1742155000",
				EndTime:     "1742165000",
				Duration:    7200, // 2 hours
				Date:        movieTimeSlotDate,
				MovieFormat: "4DX",
			},
		}, []models.SeatMatrix{
			{Row: 1, Column: 1, Price: 700, SeatNumber: "B1", IsBooked: false, Type: "Gold"},
			{Row: 1, Column: 2, Price: 700, SeatNumber: "B2", IsBooked: true, Type: "Gold"},
		})

		if err != nil {
			t.Error(err.Error())
			return
		}

		if status != 200 {
			t.Errorf("Status should be 200 after succesful addition of movies")
			return
		}

	})
}
