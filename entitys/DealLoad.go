package entitys


//| id                 | int(11)       | NO   | PRI | NULL    | auto_increment |
//| deal_id            | int(11)       | NO   | MUL | NULL    |                |
//| user_id            | int(11)       | NO   | MUL | NULL    |                |
//| user_name          | varchar(50)   | NO   |     | NULL    |                |
//| money              | decimal(20,4) | NO   |     | NULL    |                |
//| create_time        | int(11)       | NO   |     | NULL    |                |
//| is_repay           | tinyint(1)    | NO   |     | NULL    |                |
//| is_rebate          | tinyint(1)    | NO   |     | NULL    |                |
//| is_auto            | tinyint(1)    | NO   |     | NULL    |                |
//| cg_subjectAuthCode | varchar(32)   | YES  |     | NULL    |                |
//| pP2PBillNo         | varchar(30)   | YES  |     | NULL    |                |
//| pContractNo        | varchar(30)   | YES  |     | NULL    |                |
//| pMerBillNo         | varchar(30)   | YES  |     | NULL    |                |
//| is_has_loans       | tinyint(1)    | NO   |     | 0       |                |
//| msg                | varchar(100)  | YES  |     | NULL    |                |


type DealLoad struct {
	Id      int64   `json:"id"  gorm:"column:id"`
	DealId int64       `json:"deal_id"  gorm:"column:deal_id"`
	UserId  int64   `json:"user_id"  gorm:"column:user_id"`
	UserName string  `json:"user_name"  gorm:"column:user_name"`
	Money   float64  `json:"money"  gorm:"column:money"`
	CreateTime int64  `json:"create_time"  gorm:"column:create_time"`
	IsRepay  int   `json:"create_time"  gorm:"column:create_time"`
	IsRebate  int     `json:"is_rebate"  gorm:"column:is_rebate"`
	IsAuto  int    `json:"is_auto"  gorm:"column:is_auto"`
	CgSubjectAuthCode string `json:"cg_subjectAuthCode"  gorm:"column:cg_subjectAuthCode"`
	PP2PBillNo  string    `json:"pP2PBillNo"  gorm:"column:pP2PBillNo"`
	PContractNo string   `json:"pContractNo"  gorm:"column:pContractNo"`
	PMerBillNo  string   `json:"pMerBillNo"  gorm:"column:pMerBillNo"`
	IsHasLoans int      `json:"is_has_loans"  gorm:"column:is_has_loans"`
	Msg    string    `json:"is_has_loans"  gorm:"column:is_has_loans"`
}

func (t DealLoad) TableName() string {
	return "nhyd_deal_load"
}