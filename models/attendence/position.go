package attendence


import "time"


//职位
type Position struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"index" json:"deleted_at,omitempty"`

	Name      string      `json:"name"`
	TreePath  string      `json:"tree_path"`
	ParentID  uint        `json:"parent_id"`
	Orders    int         `json:"orders"`   //排序
}
