package entitys

type TopNav struct {
	Id int64       `json:"id"  gorm:"column:id"`
	Name string    `json:"name"  gorm:"column:name"`
	Action string  `json:"action"  gorm:"column:action"`

}

func (t TopNav) TableName() string {
	return "bili_top_nav"
}

type HeadNav struct {
	Id int64       `json:"id"  gorm:"column:id"`
	Name string    `json:"name"  gorm:"column:name"`
	Action string  `json:"action"  gorm:"column:action"`
	Icon string    `json:"icon"  gorm:"column:icon"`
	NavId int64    `json:"nav_id"  gorm:"column:nav_id"`
}


func (h HeadNav) TableName() string {
	return "bili_head_nav"
}