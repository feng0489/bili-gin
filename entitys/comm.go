package entitys


type Result struct {
	Total int `json:"total"  gorm:"column:total"`
	Allsum []uint8 `json:"allsum"  gorm:"column:allsum"`
}

