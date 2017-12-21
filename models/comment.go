package models

import "time"

type Comment struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	UserID    uint   `json:"user_id"`
	Content   string `json:"content" `
	PostID    uint   `json:"post_id" `
	ReadState bool   `json:"read_state" sql:"default:'0'"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	GithubUrl string `json:"github_url"`
}
