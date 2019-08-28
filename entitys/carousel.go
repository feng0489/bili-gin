package entitys


type CarouselNav struct {
	Id int64 `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Tab string `json:"tab" gorm:"column:tab"`
}

func (t CarouselNav) TableName() string {
	return "bili_carousel_nav"
}

type Carousel struct {
	Id int64 `json:"id" gorm:"column:id"`
	Title string `json:"title" gorm:"column:title"`
	Photo string `json:"photo" gorm:"column:photo"`
	JumpUrl string `json:"jump_url" gorm:"column:jump_url"`
	Tab string `json:"tab" gorm:"column:tab"`
	TabName string `json:"tab_name" gorm:"column:tab_name"`
	Sort int64 `json:"sort" gorm:"column:sort"`
	Open int `json:"open" gorm:"column:open"`
}

func (t Carousel) TableName() string {
	return "bili_carousel"
}