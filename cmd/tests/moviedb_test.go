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
			Title:           "Doctor Strange in the Multiverse of Madness",
			Description:     "Dr. Stephen Strange casts a forbidden spell that opens the doorway to the multiverse, including alternate versions of himself, whose threat to humanity is too great for the combined forces of Strange, Wong, and Wanda Maximoff.",
			ReleaseDate:     releaseDate,
			PosterURL:       "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/9Gtg2DzBhmYamXBS1hKAhiwbBKS.jpg",
			Duration:        126,
			Language:        pq.StringArray([]string{"English", "Hindi", "Tamil"}),
			Type:            pq.StringArray([]string{"Action", "Adventure", "Fantasy"}),
			MovieResolution: pq.StringArray([]string{"4K", "1080p", "720p"}),
			Venues: []models.Venue{
				{
					Name:      "PVR Cinemas",
					Type:      "4DX",
					Address:   "PVR Mall, Anna Nagar, Chennai, Tamil Nadu 600040",
					Latitude:  13.0827,
					Longitude: 80.2707,
					Rows:      10,
					Columns:   12,
					// Seats: []models.SeatMatrix{
					// 	{Row: 1, Column: 1, Price: 700, SeatNumber: "B1", IsBooked: false, Type: "Gold"},
					// 	{Row: 1, Column: 2, Price: 700, SeatNumber: "B2", IsBooked: true, Type: "Gold"},
					// },
					LanguagesSupported:   []string{"English", "Tamil", "Telugu"},
					ScreenNumber:         1,
					MovieFormatSupported: []string{"2D", "4DX", "3D"},
					// MovieTimeSlots: []models.MovieTimeSlot{
					// 	{
					// 		StartTime:   "1742155000",
					// 		EndTime:     "1742165000",
					// 		Duration:    7200, // 2 hours
					// 		Date:        movieTimeSlotDate,
					// 		MovieFormat: "4DX",
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
