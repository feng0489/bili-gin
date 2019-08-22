package entitys


type CarouselNav struct {
	Id int64 `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Tab string `json:"tab" gorm:"column:tab"`
}

type Carousel struct {
	Id int64 `json:"id" gorm:"column:id"`
	Title string `json:"title" gorm:"column:title"`
	Photo string `json:"photo" gorm:"column:photo"`
	JumpUrl string `json:"jump_url" gorm:"column:jump_url"`
	Tab string `json:"tab" gorm:"column:tab"`
	TabName string `json:"tab_name" gorm:"column:tab_name"`
	Sort int64 `json:"sort" gorm:"column:sort"`
	Opne int `json:"opne" gorm:"column:opne"`
}