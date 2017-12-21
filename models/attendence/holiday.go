package attendence


import "time"

//请假
type Holiday struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	AttendanceNumber   string    `json:"attendance_number"` //考勤机序号
	Years        int     `json:"years"`
	Day          int     `json:"day"`
}