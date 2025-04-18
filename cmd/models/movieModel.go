package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type CastAndCrew struct {
	gorm.Model
	Type      string `json:"type" gorm:"not null"` // cast or crew
	Name      string `json:"name" gorm:"not null"`
	Character string `json:"character"`
	PhotoURL  string `json:"photo_url"`
	MovieID   uint
}

type SeatMatrix struct {
	gorm.Model
	SeatNumber string `json:"seat_number" gorm:"not null"`
	IsBooked   bool   `json:"is_booked"`
	Type       string `json:"type"` // seat type (e.g., 2D, 3D, 4DX)
	Price      int    `json:"price"`
	Row        int    `json:"row"`
	Column     int    `json:"column"`
	VenueID    uint   `json:"venue_id"`
}

type MovieTimeSlot struct {
	gorm.Model
	StartTime   string    `json:"start_time" gorm:"not null"`
	EndTime     string    `json:"end_time" gorm:"not null"`
	Duration    int       `json:"duration" gorm:"not null"` // in minutes
	MovieID     uint      `json:"movie_id"`
	Date        time.Time `json:"date" gorm:"not null"`
	MovieFormat string    `json:"movie_format" gorm:"not null"` // movie format (e.g., 2D, 3D)
	VenueID     uint      `json:"venue_id"`
}

// Movie model
type Movie struct {
	ID              uint `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Title           string         `json:"title" gorm:"not null;unique"`
	Description     string         `json:"description" gorm:"not null"`
	Duration        int            `json:"duration" gorm:"not null"`
	Language        pq.StringArray `json:"language" gorm:"type:text[];not null"`
	Type            pq.StringArray `json:"type" gorm:"type:text[];not null"`
	CastCrew        []CastAndCrew  `json:"cast_crew"`
	PosterURL       string         `json:"poster_url"`
	TrailerURL      string         `json:"trailer_url"`
	ReleaseDate     time.Time      `json:"release_date" gorm:"not null"`
	MovieResolution pq.StringArray `json:"movie_resolution" gorm:"type:text[];not null"`
	Venues          []Venue        `json:"venues" gorm:"many2many:movie_venues;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Ranking         uint           `json:"ranking"`
	Votes           uint           `json:"votes"`
}

// Venue model
type Venue struct {
	gorm.Model
	Name                 string         `json:"name" gorm:"not null"`
	Type                 string         `json:"type" gorm:"not null"`
	Address              string         `json:"address" gorm:"not null"`
	Rows                 int            `json:"rows" gorm:"not null"`
	Columns              int            `json:"columns" gorm:"not null"`
	ScreenNumber         int            `json:"screen_number" gorm:"not null"`
	Longitude            float64        `json:"longitude" gorm:"not null"`
	Latitude             float64        `json:"latitude" gorm:"not null"`
	MovieFormatSupported pq.StringArray `json:"movie_format_supported" gorm:"type:text[];not null"`
	LanguagesSupported   pq.StringArray `json:"languages_supported" gorm:"type:text[];not null"`

	// Relationships
	Seats          []SeatMatrix    `json:"seats" gorm:"foreignKey:VenueID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MovieTimeSlots []MovieTimeSlot `json:"movie_time_slots" gorm:"foreignKey:VenueID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Movies         []Movie         `json:"movies" gorm:"many2many:movie_venues;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
