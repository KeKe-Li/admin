package models

import "time"

type Subscriber struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	Email          string `json:"email" sql:"unique_index"`
	VerifyState    bool   `json:"verify_state" sql:"default:'0'"`
	SubscribeState bool   `json:"sub_scribe_state" sql:"default:'1'"`
	OutTime    time.Time   `json:"out_time"`
	Secretkey  string      `json:"secret_key"`
	Signature  string     `json:"signature"`
}
