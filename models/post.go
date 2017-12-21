package models

import "time"

type Post struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	Title       string     `json:"title"`
	Body        string     `json:"body"`
	View        int        `json:"view"`
	IsPublished bool       `json:"is_published"`
	Tags        []*Tag     `gorm:"-"`
	Comments    []*Comment `gorm:"-"`
}
