package data

import (
	"time"
)

// Movie defines the structure of our struct
type Movie struct {
	ID int64 `json:"id"`
	// CreatedAt time.Time `json:"created_at"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	RunTime   int32     `json:"run_time,omitempty,string"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
