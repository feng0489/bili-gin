package entitys

//+-------------------------------+---------------+------+-----+---------+----------------+
//| id                            | int(11)       | NO   | PRI | NULL    | auto_increment |
//| deal_type                     | tinyint(1)    | NO   |     | 0       |                |
//| is_tonglian                   | tinyint(1)    | YES  |     | 0       |                |
//| real_time_status              | int(5)        | YES  |     | 0       |                |
//| real_time_repayplan           | tinyint(1)    | YES  |     | 0       |                |
//| real_time_creditor            | tinyint(1)    | YES  |     | 0       |                |
//| cg_sync                       | tinyint(1)    | YES  |     | 0       |                |
//| cg_subjectNo                  | varchar(32)   | YES  |     | NULL    |                |
//| cg_signing                    | tinyint(4)    | YES  |     | 0       |                |
//| can_repay_ben                 | tinyint(1)    | YES  |     | 0       |                |
//| name                          | text          | NO   |     | NULL    |                |
//| sub_name                      | varchar(255)  | NO   |     | NULL    |                |
//| voucher                       | varchar(250)  | YES  |     | NULL    |                |
//| mortgage_name                 | varchar(30)   | NO   |     | NULL    |                |
//| mortgage_idno                 | varchar(32)   | NO   |     | NULL    |                |
//| auth_name                     | varchar(30)   | NO   |     | NULL    |                |
//| cate_id                       | int(11)       | NO   | MUL | NULL    |                |
//| contract_id                   | int(5)        | YES  |     | NULL    |                |
//| agency_id                     | int(11)       | NO   |     | NULL    |                |
//| user_id                       | int(11)       | NO   | MUL | NULL    |                |
//| tr_id                         | int(11)       | NO   |     | 0       |                |
//| description                   | text          | NO   |     | NULL    |                |
//| eqb_zhaiyao_data              | varchar(255)  | YES  |     | 0       |                |
//| is_effect                     | tinyint(1)    | NO   |     | NULL    |                |
//| is_delete                     | tinyint(1)    | NO   |     | NULL    |                |
//| sort                          | int(11)       | NO   | MUL | NULL    |                |
//| type_id                       | int(11)       | NO   |     | NULL    |                |
//| icon_type                     | tinyint(1)    | NO   |     | NULL    |                |
//| icon                          | varchar(255)  | NO   |     | NULL    |                |
//| seo_title                     | text          | NO   |     | NULL    |                |
//| seo_keyword                   | text          | NO   |     | NULL    |                |
//| seo_description               | text          | NO   |     | NULL    |                |
//| is_hot                        | tinyint(1)    | NO   |     | NULL    |                |
//| is_new                        | tinyint(1)    | NO   |     | NULL    |                |
//| is_best                       | tinyint(1)    | NO   |     | NULL    |                |
//| borrow_amount                 | decimal(10,0) | NO   |     | NULL    |                |
//| min_loan_money                | decimal(10,0) | NO   |     | 50      |                |
//| max_loan_money                | decimal(10,0) | NO   |     | 0       |                |
//| repay_time                    | int(11)       | NO   |     | NULL    |                |
//| rate                          | decimal(10,2) | NO   |     | NULL    |                |
//| day                           | int(1)        | NO   |     | NULL    |                |
//| create_time                   | int(11)       | NO   | MUL | NULL    |                |
//| update_time                   | int(11)       | NO   | MUL | NULL    |                |
//| name_match                    | text          | NO   | MUL | NULL    |                |
//| name_match_row                | text          | NO   |     | NULL    |                |
//| deal_cate_match               | text          | NO   | MUL | NULL    |                |
//| deal_cate_match_row           | text          | NO   |     | NULL    |                |
//| tag_match                     | text          | NO   | MUL | NULL    |                |
//| tag_match_row                 | text          | NO   |     | NULL    |                |
//| type_match                    | text          | NO   |     | NULL    |                |
//| type_match_row                | text          | NO   |     | NULL    |                |
//| is_recommend                  | tinyint(1)    | NO   |     | NULL    |                |
//| buy_count                     | int(11)       | NO   |     | NULL    |                |
//| load_money                    | decimal(20,4) | NO   |     | NULL    |                |
//| repay_money                   | decimal(20,4) | NO   |     | NULL    |                |
//| start_time                    | int(11)       | NO   |     | NULL    |                |
//| success_time                  | int(11)       | NO   |     | NULL    |                |
//| repay_start_time              | int(11)       | NO   |     | NULL    |                |
//| last_repay_time               | int(11)       | NO   |     | NULL    |                |
//| next_repay_time               | int(11)       | NO   |     | NULL    |                |
//| bad_time                      | int(11)       | NO   |     | NULL    |                |
//| deal_status                   | tinyint(4)    | NO   | MUL | NULL    |                |
//| enddate                       | int(11)       | NO   |     | NULL    |                |
//| voffice                       | tinyint(1)    | NO   |     | NULL    |                |
//| vposition                     | tinyint(1)    | NO   |     | NULL    |                |
//| services_fee                  | varchar(20)   | NO   |     | NULL    |                |
//| publish_wait                  | tinyint(1)    | NO   |     | NULL    |                |
//| is_send_bad_msg               | tinyint(1)    | NO   |     | NULL    |                |
//| is_send_success_msg           | tinyint(1)    | NO   |     | NULL    |                |
//| bad_msg                       | text          | NO   |     | NULL    |                |
//| send_half_msg_time            | int(11)       | NO   |     | NULL    |                |
//| send_three_msg_time           | int(11)       | NO   |     | NULL    |                |
//| is_send_half_msg              | tinyint(1)    | NO   |     | NULL    |                |
//| is_has_loans                  | tinyint(11)   | NO   |     | NULL    |                |
//| loantype                      | tinyint(1)    | NO   |     | NULL    |                |
//| warrant                       | tinyint(1)    | NO   |     | NULL    |                |
//| titlecolor                    | varchar(20)   | NO   |     | NULL    |                |
//| is_send_contract              | tinyint(1)    | NO   |     | NULL    |                |
//| repay_time_type               | tinyint(1)    | NO   |     | 1       |                |
//| risk_rank                     | tinyint(1)    | NO   |     | 0       |                |
//| deal_sn                       | varchar(50)   | NO   |     | NULL    |                |
//| is_has_received               | tinyint(1)    | NO   |     | NULL    |                |
//| manage_fee                    | varchar(30)   | NO   |     | NULL    |                |
//| user_loan_manage_fee          | varchar(30)   | NO   |     | NULL    |                |
//| manage_impose_fee_day1        | varchar(30)   | NO   |     | NULL    |                |
//| manage_impose_fee_day2        | varchar(30)   | NO   |     | NULL    |                |
//| impose_fee_day1               | varchar(30)   | NO   |     | NULL    |                |
//| impose_fee_day2               | varchar(30)   | NO   |     | NULL    |                |
//| user_load_transfer_fee        | varchar(30)   | NO   |     | NULL    |                |
//| compensate_fee                | varchar(30)   | NO   |     | NULL    |                |
//| ips_do_transfer               | tinyint(1)    | NO   |     | 0       |                |
//| ips_over                      | tinyint(1)    | NO   |     | 0       |                |
//| delete_msg                    | text          | NO   |     | NULL    |                |
//| user_bid_rebate               | varchar(30)   | NO   |     | NULL    |                |
//| guarantees_amt                | decimal(20,2) | NO   |     | 0.00    |                |
//| real_freezen_amt              | decimal(20,2) | YES  |     | 0.00    |                |
//| un_real_freezen_amt           | decimal(20,2) | YES  |     | 0.00    |                |
//| guarantor_amt                 | decimal(20,2) | YES  |     | 0.00    |                |
//| guarantor_margin_amt          | decimal(20,2) | YES  |     | 0.00    |                |
//| guarantor_real_freezen_amt    | decimal(20,2) | YES  |     | 0.00    |                |
//| un_guarantor_real_freezen_amt | decimal(20,2) | YES  |     | 0.00    |                |
//| guarantor_pro_fit_amt         | decimal(20,2) | YES  |     | 0.00    |                |
//| guarantor_real_fit_amt        | decimal(20,2) | YES  |     | 0.00    |                |
//| mer_bill_no                   | varchar(30)   | YES  |     | NULL    |                |
//| ips_bill_no                   | varchar(30)   | YES  |     | NULL    |                |
//| ips_guarantor_bill_no         | varchar(30)   | YES  |     | NULL    |                |
//| mer_guarantor_bill_no         | varchar(30)   | YES  |     | NULL    |                |
//| isdf                          | tinyint(2)    | NO   |     | 0       |                |
//| extend_1                      | varchar(255)  | YES  |     | NULL    |                |
//| extend_2                      | varchar(255)  | YES  |     | NULL    |                |
//| extend_3                      | varchar(255)  | YES  |     | NULL    |                |
//| extend_4                      | varchar(255)  | YES  |     | NULL    |                |
//| extend_5                      | varchar(255)  | YES  |     | NULL    |                |
//| extend_6                      | varchar(255)  | YES  |     | NULL    |                |
//| extend_7                      | varchar(255)  | YES  |     | NULL    |                |
//| extend_8                      | varchar(255)  | YES  |     | NULL    |                |
//| extend_9                      | varchar(255)  | YES  |     | NULL    |                |
//| extend_10                     | varchar(255)  | YES  |     | NULL    |                |
//+-------------------------------+---------------+------+-----+------

type Deel struct {
	Id                        int64   `json:"id"  gorm:"column:id"`
	DealType                  int     `json:"deal_type"  gorm:"column:deal_type"`
	IsTonglian                int     `json:"is_tonglian"  gorm:"column:is_tonglian"`
	RealTimeStatus            int     `json:"real_time_status"  gorm:"column:real_time_status"`
	RealTimeRepayplan         int     `json:"real_time_repayplan"  gorm:"column:real_time_repayplan"`
	RealTimeCreditor          int     `json:"real_time_creditor"  gorm:"column:real_time_creditor"`
	CgSync                    int     `json:"cg_sync"  gorm:"column:cg_sync"`
	CgSubjectNo               string  `json:"cg_subjectNo"  gorm:"column:cg_subjectNo"`
	CgSigning                 int     `json:"cg_signing"  gorm:"column:cg_signing"`
	Name                      string  `json:"name"  gorm:"column:name"`
	SubName                   string  `json:"sub_name"  gorm:"column:sub_name"`
	Voucher                   string  `json:"voucher"  gorm:"column:voucher"`
	MortgageName              string  `json:"mortgage_name"  gorm:"column:mortgage_name"`
	MortgageIdno              string  `json:"mortgage_idno"  gorm:"column:mortgage_idno"`
	AuthName                  string  `json:"auth_name"  gorm:"column:auth_name"`
	CateId                    int64   `json:"cate_id"  gorm:"column:cate_id"`
	ContractId                int64   `json:"contract_id"  gorm:"column:contract_id"`
	AgencyId                  int64   `json:"agency_id"  gorm:"column:agency_id"`
	UserId                    int64   `json:"user_id"  gorm:"column:user_id"`
	TrId                      int64   `json:"tr_id"  gorm:"column:tr_id"`
	Description               string  `json:"description"  gorm:"column:description"`
	EqbZhaiyaoData            string  `json:"eqb_zhaiyao_data"  gorm:"column:eqb_zhaiyao_data"`
	IsEffect                  int     `json:"is_effect"  gorm:"column:is_effect"`
	IsDelete                  int     `json:"is_delete"  gorm:"column:is_delete"`
	Sort                      int     `json:"sort"  gorm:"column:sort"`
	TypeId                    int     `json:"type_id"  gorm:"column:type_id"`
	IconYype                  int     `json:"icon_type"  gorm:"column:icon_type"`
	Icon                      string  `json:"icon"  gorm:"column:icon"`
	SeoTitle                  string  `json:"seo_title"  gorm:"column:seo_title"`
	SeoKeyword                string  `json:"seo_keyword"  gorm:"column:seo_keyword"`
	SeoDescription            string  `json:"seo_description"  gorm:"column:seo_description"`
	IsHot                     int     `json:"is_hot"  gorm:"column:is_hot"`
	IsNew                     int     `json:"is_new"  gorm:"column:is_new"`
	IsBest                    int     `json:"is_best"  gorm:"column:is_best"`
	BorrowAmount              float32 `json:"borrow_amount"  gorm:"column:borrow_amount"`
	MinLoanMoney              float32 `json:"min_loan_money"  gorm:"column:min_loan_money"`
	MaxLoanMoney              float32 `json:"max_loan_money"  gorm:"column:max_loan_money"`
	RepayTime                 int64   `json:"repay_time"  gorm:"column:repay_time"`
	Rate                      float32 `json:"rate"  gorm:"column:rate"`
	Day                       int     `json:"day"  gorm:"column:day"`
	CreateTime                int64   `json:"create_time"  gorm:"column:create_time"`
	NameMatch                 string  `json:"name_match"  gorm:"column:name_match"`
	NameMatchRow              string  `json:"name_match_row"  gorm:"column:name_match_row"`
	DealCateMatch             string  `json:"deal_cate_match"  gorm:"column:deal_cate_match"`
	DealCateMatchRow          string  `json:"deal_cate_match_row"  gorm:"column:deal_cate_match_row"`
	TagMatch                  string  `json:"tag_match"  gorm:"column:tag_match"`
	TagMatchRow               string  `json:"tag_match_row"  gorm:"column:tag_match_row"`
	TypeMatch                 string  `json:"type_match"  gorm:"column:type_match"`
	TypeMatchRow              string  `json:"type_match_row"  gorm:"column:type_match_row"`
	IsRecommend               int     `json:"is_recommend"  gorm:"column:is_recommend"`
	BuyCount                  int64   `json:"buy_count"  gorm:"column:buy_count"`
	LoadMoney                 float64 `json:"load_money"  gorm:"column:load_money"`
	RepayMoney                float64 `json:"repay_money"  gorm:"column:repay_money"`
	StartTime                 int64   `json:"start_time"  gorm:"column:start_time"`
	SuccessTime               int64   `json:"success_time"  gorm:"column:success_time"`
	RepayStartTime            int64   `json:"repay_start_time"  gorm:"column:repay_start_time"`
	LastRepayTime             int64   `json:"last_repay_time"  gorm:"column:last_repay_time"`
	NextRepayTime             int64   `json:"next_repay_time"  gorm:"column:next_repay_time"`
	BadTime                   int64   `json:"bad_time"  gorm:"column:bad_time"`
	DealStatus                int     `json:"deal_status"  gorm:"column:deal_status"`
	Enddate                   int64   `json:"enddate"  gorm:"column:enddate"`
	Voffice                   int     `json:"voffice"  gorm:"column:voffice"`
	Vposition                 int     `json:"vposition"  gorm:"column:vposition"`
	ServicesFee               string  `json:"services_fee"  gorm:"column:services_fee"`
	PublishWait               int     `json:"publish_wait"  gorm:"column:publish_wait"`
	IsSendBadMsg              int     `json:"is_send_bad_msg"  gorm:"column:is_send_bad_msgsss"`
	IsSendSuccessMsg          int     `json:"is_send_success_msg"  gorm:"column:is_send_success_msg"`
	BadMsg                    string  `json:"bad_msg"  gorm:"column:bad_msg"`
	SendHalfMsgTime           int64   `json:"send_half_msg_time"  gorm:"column:send_half_msg_time"`
	SendThreeMsgTime          int64   `json:"send_three_msg_time"  gorm:"column:send_three_msg_time"`
	IsSendHalfMsg             int     `json:"is_send_half_msg"  gorm:"column:is_send_half_msg"`
	IsHasLoans                int     `json:"is_has_loans"  gorm:"column:is_has_loans"`
	Loantype                  int     `json:"loantype"  gorm:"column:loantype"`
	Warrant                   int     `json:"warrant"  gorm:"column:warrant"`
	Titlecolor                string  `json:"titlecolor"  gorm:"column:titlecolor"`
	IsSendContract            int     `json:"is_send_contract"  gorm:"column:is_send_contract"`
	RepayTimeType             int     `json:"repay_time_type"  gorm:"column:repay_time_type"`
	RiskRank                  int     `json:"risk_rank"  gorm:"column:risk_rank"`
	DealSn                    string  `json:"deal_sn"  gorm:"column:deal_sn"`
	IsHasReceived             int     `json:"is_has_received"  gorm:"column:is_has_received"`
	ManageFee                 string  `json:"manage_fee"  gorm:"column:manage_fee"`
	UserLoanManageFee         string  `json:"user_loan_manage_fee"  gorm:"column:user_loan_manage_fee"`
	ManageImposeFeeDay1       string  `json:"manage_impose_fee_day1"  gorm:"column:manage_impose_fee_day1"`
	ManageImposeFeeDay2       string  `json:"manage_impose_fee_day2"  gorm:"column:manage_impose_fee_day2"`
	ImposeFeeDay1             string  `json:"impose_fee_day1"  gorm:"column:impose_fee_day1"`
	ImposeFeeDay2             string  `json:"impose_fee_day2"  gorm:"column:impose_fee_day2"`
	UserLoadTransferFee       string  `json:"user_load_transfer_fee"  gorm:"column:user_load_transfer_fee"`
	CompensateFee             string  `json:"compensate_fee"  gorm:"column:compensate_fee"`
	IpsDoTransfer             int     `json:"ips_do_transfer"  gorm:"column:ips_do_transfer"`
	IpsOver                   int     `json:"ips_over"  gorm:"column:ips_over"`
	DeleteMsg                 string  `json:"delete_msg"  gorm:"column:delete_msg"`
	UserBidRebate             string  `json:"user_bid_rebate"  gorm:"column:user_bid_rebate"`
	GuaranteesAmt             float64 `json:"guarantees_amt"  gorm:"column:guarantees_amt"`
	RealFreezenAmt            float64 `json:"real_freezen_amt"  gorm:"column:real_freezen_amt"`
	UnRealFreezenAmt          float64 `json:"un_real_freezen_amt"  gorm:"column:un_real_freezen_amt"`
	GuarantorAmt              float64 `json:"guarantor_amt"  gorm:"column:guarantor_amt"`
	GuarantorMarginAmt        float64 `json:"guarantor_margin_amt"  gorm:"column:guarantor_margin_amt"`
	GuarantorRealFreezenAmt   float64 `json:"guarantor_real_freezen_amt"  gorm:"column:guarantor_real_freezen_amt"`
	UnGuarantorRealFreezenAmt float64 `json:"un_guarantor_real_freezen_amt"  gorm:"column:un_guarantor_real_freezen_amt"`
	GuarantorProFitAmt        float64 `json:"guarantor_pro_fit_amt"  gorm:"column:guarantor_pro_fit_amt"`
	GuarantorRealFitAmt       float64 `json:"guarantor_real_fit_amt"  gorm:"column:guarantor_real_fit_amt"`
	MerBillNo                 string  `json:"mer_bill_no"  gorm:"column:mer_bill_no"`
	IpsBillNo                 string  `json:"ips_bill_no"  gorm:"column:ips_bill_no"`
	IpsGuarantorBillNo        string  `json:"ips_guarantor_bill_no"  gorm:"column:ips_guarantor_bill_no"`
	MerGuarantorBillNo        string  `json:"mer_guarantor_bill_no"  gorm:"column:mer_guarantor_bill_no"`
	Isdf                      int     `json:"isdf"  gorm:"column:isdf"`
	Extend1                   string  `json:"extend_1"  gorm:"column:extend_1"`
	Extend2                   string  `json:"extend_2"  gorm:"column:extend_2"`
	Extend3                   string  `json:"extend_3"  gorm:"column:extend_3"`
	Extend4                   string  `json:"extend_4"  gorm:"column:extend_4"`
	Extend5                   string  `json:"extend_5"  gorm:"column:extend_5"`
	Extend6                   string  `json:"extend_6"  gorm:"column:extend_6"`
	Extend7                   string  `json:"extend_7"  gorm:"column:extend_7"`
	Extend8                   string  `json:"extend_8"  gorm:"column:extend_8"`
	Extend9                   string  `json:"extend_9"  gorm:"column:extend_9"`
	Extend10                  string  `json:"extend_10"  gorm:"column:extend_10"`
}

func (t Deel) TableName() string {
	return "nhyd_deal"
}