package attendence

import "time"

//工作时间
type WorkTime struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
