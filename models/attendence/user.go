package attendence

import "time"

//员工
type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	Name             string    `json:"name"`              // 昵称
	AttendanceNumber string    `json:"attendance_number"` //考勤机序号
	TeambitionSign   string    `json:"teambition_sign"`
	TeambitionToken  string    `json:"teambition_token"`
	Avatar           string    `json:"avatar"` //头像
	Email            string    `json:"email" sql:"size:255"`
	Telephone        string    `json:"telephone" sql:"size:20;"`
	EntryDate        time.Time `json:"entry_date"`    //入职时间
	LeaveDate        time.Time `json:"leave_date"`    //离职时间
	Birthday         time.Time `json:"birthday"`      //生日
	IsAdmin          bool      `json:"is_admin"`      //是否是管理员
	Status           string    `json:"status"`        //员工状态
	DepartmentID     uint      `json:"department_id"` //部门
	PositionID       uint      `json:"position_id"`   //职位
	UserTypeID       uint      `json:"user_type_id"`  //员工类型
}
