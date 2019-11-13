package entitys

//
//| id                 | int(11)       | NO   | PRI | NULL    | auto_increment |
//| deal_id            | int(11)       | NO   |     | NULL    |                |
//| user_id            | int(11)       | NO   | MUL | NULL    |                |
//| repay_money        | decimal(20,4) | NO   |     | NULL    |                |
//| manage_money       | decimal(20,4) | NO   |     | NULL    |                |
//| impose_money       | decimal(20,4) | NO   |     | NULL    |                |
//| repay_time         | int(11)       | NO   |     | NULL    |                |
//| true_repay_time    | int(11)       | NO   |     | NULL    |                |
//| status             | tinyint(1)    | NO   |     | NULL    |                |
//| l_key              | int(11)       | NO   | MUL | 0       |                |
//| has_repay          | tinyint(1)    | NO   |     | 0       |                |
//| manage_fee_status  | tinyint(1)    | NO   |     | NULL    |                |
//| mange_impose_money | decimal(20,2) | NO   |     | NULL    |                |

type DealRepay struct {
	Id               int64   `json:"id"  gorm:"column:id"`
	DealId           int64   `json:"deal_id"  gorm:"column:deal_id"`
	UserIsd          int64   `json:"user_id"  gorm:"column:user_id"`
	RepayMoney       float64 `json:"repay_money"  gorm:"column:repay_money"`
	ManageMoney      float64 `json:"manage_money"  gorm:"column:manage_money"`
	ImposeMoney      float64 `json:"impose_money"  gorm:"column:impose_money"`
	RepayTime        int64   `json:"repay_time"  gorm:"column:repay_time"`
	TrueRepayTime    int64   `json:"true_repay_time"  gorm:"column:true_repay_time"`
	Status           int     `json:"status"  gorm:"column:status"`
	LKey             int64   `json:"l_key"  gorm:"column:l_key"`
	HasRepay         int     `json:"has_repay"  gorm:"column:has_repay"`
	ManageFeeStatus  int     `json:"manage_fee_status"  gorm:"column:manage_fee_status"`
	MangeImposeMoney float64 `json:"mange_impose_money"  gorm:"column:mange_impose_money"`
}

func (t DealRepay) TableName() string {
	return "nhyd_deal_repay"
}