package attendence

import "time"

//外出
type Egress struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	UserID      uint      `json:"user_id"`
	StartTime   time.Time `json:"start_time"`   //开始时间
	EndTime     time.Time `json:"end_time"`     //开始时间
	WorkPlace   string    `json:"work_place"`   //工作地点
	WorkContent string    `json:"work_content"` //工作的内容
}
