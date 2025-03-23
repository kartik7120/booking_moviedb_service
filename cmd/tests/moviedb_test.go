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

		if testing.Short() {
			t.Skip("Skipping this test in short mode")
		}

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
			Title:           "The Lord of the Rings: The Fellowship of the Ring",
			Description:     "A young hobbit, Frodo Baggins, embarks on a journey to destroy the One Ring and defeat the Dark Lord Sauron.",
			ReleaseDate:     releaseDate,
			PosterURL:       "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/6oom5QYQ2yQTMJIbnvbkBL9cHo6.jpg",
			Duration:        178, // 2 hours 58 minutes
			Language:        pq.StringArray([]string{"English", "Elvish", "Dwarvish"}),
			Type:            pq.StringArray([]string{"Fantasy", "Adventure", "Drama"}),
			MovieResolution: pq.StringArray([]string{"4K", "1080p", "720p"}),
			CastCrew: []models.CastAndCrew{
				{Type: "Cast", Name: "Elijah Wood", Character: "Frodo Baggins", PhotoURL: "https://example.com/elijah_wood.jpg"},
				{Type: "Cast", Name: "Ian McKellen", Character: "Gandalf", PhotoURL: "https://example.com/ian_mckellen.jpg"},
				{Type: "Cast", Name: "Viggo Mortensen", Character: "Aragorn", PhotoURL: "https://example.com/viggo_mortensen.jpg"},
				{Type: "Cast", Name: "Sean Astin", Character: "Samwise Gamgee", PhotoURL: "https://example.com/sean_astin.jpg"},
				{Type: "Crew", Name: "Peter Jackson", Character: "Director", PhotoURL: "https://example.com/peter_jackson.jpg"},
			},
			Venues: []models.Venue{
				{
					Name:                 "Rivendell Grand Theater",
					Type:                 "IMAX",
					Address:              "123 Elven Road, Middle-earth",
					Latitude:             40.7128,
					Longitude:            -74.0060,
					Rows:                 20,
					Columns:              30,
					ScreenNumber:         1,
					MovieFormatSupported: pq.StringArray([]string{"IMAX", "3D", "2D"}),
					LanguagesSupported:   pq.StringArray([]string{"English", "Elvish"}),

					// Seats: []models.SeatMatrix{
					// 	{Row: 1, Column: 1, Price: 1500, SeatNumber: "A1", IsBooked: false, Type: "Platinum"},
					// 	{Row: 1, Column: 2, Price: 1500, SeatNumber: "A2", IsBooked: true, Type: "Platinum"},
					// 	{Row: 2, Column: 1, Price: 1200, SeatNumber: "B1", IsBooked: false, Type: "Gold"},
					// 	{Row: 2, Column: 2, Price: 1200, SeatNumber: "B2", IsBooked: true, Type: "Gold"},
					// },

					// MovieTimeSlots: []models.MovieTimeSlot{
					// 	{
					// 		StartTime:   "16:00", // 4:00 PM
					// 		EndTime:     "19:00", // 7:00 PM
					// 		Duration:    180,     // 3 hours
					// 		Date:        time.Date(2025, 3, 22, 0, 0, 0, 0, time.UTC),
					// 		MovieFormat: "IMAX",
					// 	},
					// 	{
					// 		StartTime:   "20:00", // 8:00 PM
					// 		EndTime:     "23:00", // 11:00 PM
					// 		Duration:    180,     // 3 hours
					// 		Date:        time.Date(2025, 3, 22, 0, 0, 0, 0, time.UTC),
					// 		MovieFormat: "3D",
					// 	},
					// },
				},
			},
		}

		_, status, err := m.AddMovie(movie, []models.MovieTimeSlot{
			{
				StartTime:   "16:00", // 4:00 PM
				EndTime:     "19:00", // 7:00 PM
				Duration:    180,     // 3 hours
				Date:        movieTimeSlotDate,
				MovieFormat: "IMAX",
			},
			{
				StartTime:   "20:00", // 8:00 PM
				EndTime:     "23:00", // 11:00 PM
				Duration:    180,     // 3 hours
				Date:        movieTimeSlotDate,
				MovieFormat: "3D",
			},
		}, []models.SeatMatrix{
			{Row: 1, Column: 1, Price: 1500, SeatNumber: "A1", IsBooked: false, Type: "Platinum"},
			{Row: 1, Column: 2, Price: 1500, SeatNumber: "A2", IsBooked: true, Type: "Platinum"},
			{Row: 2, Column: 1, Price: 1200, SeatNumber: "B1", IsBooked: false, Type: "Gold"},
			{Row: 2, Column: 2, Price: 1200, SeatNumber: "B2", IsBooked: true, Type: "Gold"},
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

	t.Run("Update a movie in database", func(t *testing.T) {

		if testing.Short() {
			t.Skip("Skipping this test in short mode")
		}

		err := godotenv.Load()

		if err != nil {
			t.Errorf("Error loading in .env file")
			return
		}

		m := api.NewMovieDB()

		conn, err := helper.ConnectToDB()

		if err != nil {
			t.Error("Error connecting to the database", err)
			return
		}

		m.DB.Conn = conn

		movieID := 23

		updateMovieObj := models.Movie{
			Title: "Blade Runner 2050",
		}

		_, status, err := m.UpdateMovie(uint(movieID), updateMovieObj)

		if status != 200 {
			t.Errorf("Movie should have been updated")
			return
		}

		if err != nil {
			t.Error("Error updating movies", err)
			return
		}

	})

	t.Run("Delete movie in database", func(t *testing.T) {

		err := godotenv.Load()

		if err != nil {
			t.Error("Failed to load .env file")
			return
		}

		m := api.NewMovieDB()

		conn, err := helper.ConnectToDB()

		if err != nil {
			t.Error("Failed to connect to the database")
			return
		}

		m.DB.Conn = conn

		movieID := 23

		status, err := m.DeleteMovie(uint(movieID))

		if status != 200 {
			t.Error("Movie should have been deleted with status 200")
			return
		}

		if err != nil {
			t.Error("Error delete movie from database", err)
			return
		}
	})

	t.Run("Delete venue in database", func(t *testing.T) {

		err := godotenv.Load()

		if err != nil {
			t.Error("error occured when loading .env file", err)
			return
		}

		m := api.NewMovieDB()

		conn, err := helper.ConnectToDB()

		if err != nil {
			t.Error("error connecting to the database", err)
			return
		}

		m.DB.Conn = conn

		venueID := 21

		m.DB.Conn.AutoMigrate(&models.SeatMatrix{})

		status, err := m.DeleteVenue(uint(venueID))

		if status != 200 {
			t.Error("status should have been", err)
			return
		}

		if err != nil {
			t.Error("error should have been nil", err)
			return
		}
	})

	t.Run("Update a venue in database", func(t *testing.T) {

		err := godotenv.Load()

		if err != nil {
			t.Error("error loading .env file", err)
			return
		}

		m := api.NewMovieDB()

		conn, err := helper.ConnectToDB()

		if err != nil {
			t.Error("error connecting to the database", err)
			return
		}

		m.DB.Conn = conn

		venueID := 4

		venue := models.Venue{
			ScreenNumber: 3,
		}

		_, status, err := m.UpdateVenue(uint(venueID), venue)

		if status != 200 {
			t.Error("status should be 200 when updating venue", err)
			return
		}

		if err != nil {
			t.Error("error should be nil", err)
			return
		}
	})

	t.Run("Add a venue to the database", func(t *testing.T) {

		err := godotenv.Load()

		if err != nil {
			t.Fatal("error loading .env file")
			return
		}

		m := api.NewMovieDB()

		conn, err := helper.ConnectToDB()
		if err != nil {
			t.Fatalf("Error connecting to the database: %v", err) // Use t.Fatalf for fatal errors
		}

		m.DB.Conn = conn

		venue := models.Venue{
			Name:                 "IMAX Theater",
			Type:                 "Multiplex",
			Address:              "123 Movie Street, City",
			Rows:                 10,
			Columns:              20,
			ScreenNumber:         1,
			Longitude:            12.34,
			Latitude:             56.78,
			MovieFormatSupported: pq.StringArray{"2D", "3D", "IMAX"},
			LanguagesSupported:   pq.StringArray{"English", "Spanish"},
		}

		// Insert into DB
		result := m.DB.Conn.Create(&venue)
		if result.Error != nil {
			t.Errorf("Failed to add venue: %v", result.Error)
			return
		}

		// Verify venue exists
		var savedVenue models.Venue
		if err := m.DB.Conn.First(&savedVenue, venue.ID).Error; err != nil {
			t.Errorf("Venue was not saved in the database: %v", err)
		} else {
			t.Logf("Venue successfully added: %v", savedVenue)
		}
	})

	t.Run("Add venue along side movies in database", func(t *testing.T) {

		err := godotenv.Load()

		if err != nil {
			t.Fatal("error loading .env file", err)
			return
		}

		m := api.NewMovieDB()

		conn, err := helper.ConnectToDB()

		if err != nil {
			t.Fatal("error connecting to database", err)
			return
		}

		m.DB.Conn = conn

		// m.DB.Conn.Migrator().DropTable("movie_venues")
		m.DB.Conn.AutoMigrate(&models.Venue{}, &models.Venue{})

		venue := models.Venue{
			Name:                 "IMAX Theater",
			Type:                 "Multiplex",
			Address:              "123 Movie Street, City",
			Rows:                 10,
			Columns:              20,
			ScreenNumber:         1,
			Longitude:            12.34,
			Latitude:             56.78,
			MovieFormatSupported: pq.StringArray{"2D", "3D"},
			LanguagesSupported:   pq.StringArray{"English", "Spanish"},
			Movies: []models.Movie{
				{
					Title:           "Inception",
					Description:     "A mind-bending thriller",
					Duration:        148,
					Language:        pq.StringArray{"English"},
					Type:            pq.StringArray{"Sci-Fi", "Thriller"},
					ReleaseDate:     time.Now(),
					MovieResolution: pq.StringArray{"1080p", "4K"},
				},
			},
		}

		_, status, err := m.AddVenue(venue)

		if status != 200 {
			t.Error("status should be 200", err)
			return
		}

		if err != nil {
			t.Error("error should be nil", err)
			return
		}

	})
}
