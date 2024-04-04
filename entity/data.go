package entity

import "time"

type LargeData struct {
	ID          int
	Name        string
	Description string
	Age         int
	Address     string
	Email       string
	Phone       string
	City        string
	State       string
	Country     string
	ZipCode     string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Notes       []string
	Tags        []string
	Metadata    map[string]interface{}
	ImageURL    string
	Latitude    float64
	Longitude   float64
}
