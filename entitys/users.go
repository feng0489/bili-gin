package entitys

type Users struct {
	Id int64    `json:"id"  gorm:"column:id"`
	Username string `json:"user_name"  gorm:"column:username"`
	Nickname string `json:"nick_name"  gorm:"column:nickname"`
	Password string `json:"pass"  gorm:"column:password"`
	Phone string `json:"phone"  gorm:"column:phone"`
	HeadUrl string `json:"head_url"  gorm:"column:head_url"`
	Bcoin int64 `json:"coin"  gorm:"column:b_coin"`   //b币
	UserTab int64 `json:"tabs"  gorm:"column:user_tab"`  //作品数量
	Channel int64 `json:"channel"  gorm:"column:channel"`//频道数量
	StoreUp int64 `json:"store"  gorm:"column:store_up"`//收藏数量
	Focus int64 `json:"focus"  gorm:"column:focus"`//关注数量
	Follower int64 `json:"follower"  gorm:"column:follower"`//粉丝数量
	CreateTime int64 `json:"create_time"  gorm:"column:create_time"`//注册时间
	LastTime int64 `json:"last_time"  gorm:"column:last_time"`//最后登录时间
	LastIp int `json:"last_ip"  gorm:"column:last_ip"`//最后登录时间
}


func (t Users) TableName() string {
	return "bili_users"
}


type UserReg struct {
	NickName string `form:"nick_name" json:"nick_name" xml:"nick_name"  binding:"required"`
	Phone string `form:"phone" json:"phone" xml:"phone"  binding:"required"`
	Password string `form:"pass" json:"pass" xml:"pass"  binding:"required"`
}