package attendence


import "time"

//考勤记录表
type Attendance struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	Name         string  `json:"name"`
	AttendanceNumber   string    `json:"attendance_number"` //考勤机序号
	RegisterTime time.Time `json:"register_time"` //打卡时间
}