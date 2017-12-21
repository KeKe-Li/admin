package attendence

import "time"

//交通
type Transportation struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	UserID           uint      `json:"user_id"`
	Mobile           string    `json:"mobile" sql:"size:20;unique_index"` // 手机
	AttendanceNumber string    `json:"attendance_number"` //考勤序列号
	RidingTime       time.Time `json:"riding_time"` //乘车时间
	Origin           string    `json:"origin"`      //出发地
	Destination      string    `json:"destination"` //目的地
	TotalPrice       float32   `json:"total_price"` //总金额
	Distance         float32   `json:"distance"`    //距离
	Duration         int       `json:"duration"`    //时长
	City             string    `json:"city"`
	Note             string    `json:"note"`       //备注
	Additional       string    `json:"additional"` //补充说明
}
