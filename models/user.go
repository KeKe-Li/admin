package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	Email             string    `json:"email" sql:"size:255"`
	Telephone         string    `json:"telephone" sql:"size:20;unique_index"`
	EncryptedPassword string    `json:"encrypted_password" sql:"size:60,not null"`
	VerifyState       string    `json:"verify_state" sql:"default:'0'"`
	SecretKey         string    `json:"secret_key" sql:"default:null"`
	OutTime           time.Time `json:"out_time"`
	GithubLoginId     string    `json:"github_login_id" sql:"unique_index;default:null"` // github唯一标识
	GithubUrl         string    `json:"github_url"`                                      //github地址
	IsAdmin           bool      `json:"is_admin"`                                        //是否是管理员
	AvatarUrl         string    `json:"avatar_url"`                                      // 头像链接
	NickName          string    `json:"nick_name"`                                       // 昵称
	LockState         bool      `json:"lock_state" sql:"default:'0'"`                    //锁定状态
}
