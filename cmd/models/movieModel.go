package models

import "gorm.io/gorm"

type cast_and_crew struct {
	gorm.Model
	Type      string `json:"type" gorm:"required"` // cast or crew
	Name      string `json:"name" gorm:"required"`
	Character string `json:"character"`
	PhotoURL  string `json:"photo_url"`
}

type SeatMatrix struct {
	gorm.Model
	SeatNumber string `json:"seat_number"`
	IsBooked   bool   `json:"is_booked"`
	Type       string `json:"type"`
	Price      int    `json:"price"`
	Row        int    `json:"row"`
	Column     int    `json:"column"`
}

type Movie_time_slot struct {
	gorm.Model
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Duration     int    `json:"duration"`
	Movie        Movie  `json:"movie"`
	Date         string `json:"date"`
	Movie_Format string `json:"movie_format"`
	Venue        Venue  `json:"venue"`
}

type Movie struct {
	gorm.Model
	Title           string          `json:"title" gorm:"required;unique"`
	Description     string          `json:"description" gorm:"required"`
	Duration        int             `json:"duration" gorm:"required"`
	Language        []string        `json:"language" gorm:"required"`
	Type            []string        `json:"type" gorm:"required"`
	CastCrew        []cast_and_crew `json:"cast_crew"`
	PosterURL       string          `json:"poster_url"`
	TrailerURL      string          `json:"trailer_url"`
	ReleaseDate     string          `json:"release_date" gorm:"required"`
	MovieResolution []string        `json:"movie_resolution" gorm:"required"`
	Venues          []Venue         `json:"venues" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:movie_venues;"` // venues where movie will be played
}

type Venue struct {
	gorm.Model
	Name                   string            `json:"name" gorm:"required;unique"`
	Type                   string            `json:"type" gorm:"required"`
	Address                string            `json:"address" gorm:"required"`
	Rows                   int               `json:"rows" gorm:"required"`
	Columns                int               `json:"columns" gorm:"required"`
	Seats                  []SeatMatrix      `json:"seats"`
	ScreenNumber           int               `json:"screen_number" gorm:"required"`
	Longitude              float64           `json:"longitude" gorm:"required"`
	Latitude               float64           `json:"latitude" gorm:"required"`
	Movie_Format_Supported []string          `json:"movie_format_supported" gorm:"required"`
	Languages_Supported    []string          `json:"languages_supported" gorm:"required"`
	Movie_Time_Slots       []Movie_time_slot `json:"movie_time_slots" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Movies                 []Movie           `json:"movies" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:movie_venues;"`
}
