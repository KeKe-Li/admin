package attendence

import "time"

//请假
type Absence struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	UserID      uint      `json:"user_id"`
	AbsenceDate time.Time `json:"absence_date"` //请假日期
	StartTime   time.Time `json:"start_time"`   //开始时间
	EndTime     time.Time `json:"end_time"`     //开始时间
	AbsenceType string   `json:"absence_type"` //请假类型
	AbsenceNote string   `json:"absence_note"` //请假的理由
}
