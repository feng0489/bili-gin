package entitys

//| id                    | int(11)       | NO   | PRI | NULL     | auto_increment |
//| type_id               | tinyint(4)    | NO   |     | 0        |                |
//| is_company            | tinyint(1)    | YES  |     | 0        |                |
//| company_name          | varchar(255)  | YES  |     | NULL     |                |
//| company_code          | varchar(50)   | YES  |     | NULL     |                |
//| user_name             | varchar(255)  | NO   | UNI | NULL     |                |
//| user_pwd              | varchar(255)  | NO   |     | NULL     |                |
//| create_time           | int(11)       | NO   |     | NULL     |                |
//| update_time           | int(11)       | NO   |     | NULL     |                |
//| real_time_status      | tinyint(1)    | YES  |     | 0        |                |
//| cg_custNo             | varchar(50)   | YES  |     | NULL     |                |
//| cg_acctNo             | varchar(50)   | YES  |     | NULL     |                |
//| cg_data               | text          | YES  |     | NULL     |                |
//| cg_amtType            | varchar(10)   | YES  |     | NULL     |                |
//| login_ip              | varchar(255)  | NO   |     | NULL     |                |
//| group_id              | int(11)       | NO   |     | NULL     |                |
//| is_effect             | tinyint(1)    | NO   |     | NULL     |                |
//| is_delete             | tinyint(1)    | NO   |     | NULL     |                |
//| email                 | varchar(255)  | NO   |     | NULL     |                |
//| idno                  | varchar(20)   | NO   |     | NULL     |                |
//| idcardpassed          | tinyint(1)    | NO   |     | 0        |                |
//| idcardpassed_time     | int(11)       | NO   |     | NULL     |                |
//| idcard_v_return       | text          | NO   |     | NULL     |                |
//| real_name             | varchar(50)   | NO   |     | NULL     |                |
//| mobile                | varchar(255)  | NO   |     | NULL     |                |
//| mobilepassed          | tinyint(1)    | NO   |     | NULL     |                |
//| eqb_accountId         | varchar(50)   | YES  |     | NULL     |                |
//| eqb_sealData          | text          | YES  |     | NULL     |                |
//| score                 | int(11)       | NO   |     | NULL     |                |
//| money                 | decimal(20,4) | NO   |     | NULL     |                |
//| quota                 | decimal(20,0) | NO   |     | 0        |                |
//| lock_money            | decimal(10,0) | NO   |     | NULL     |                |
//| hx_money              | decimal(20,2) | NO   |     | 0.00     |                |
//| hx_lock_money         | decimal(20,2) | NO   |     | 0.00     |                |
//| verify                | varchar(255)  | NO   |     | NULL     |                |
//| code                  | varchar(255)  | NO   |     | NULL     |                |
//| pid                   | int(11)       | NO   |     | NULL     |                |
//| login_time            | int(11)       | NO   |     | NULL     |                |
//| referral_count        | int(11)       | NO   |     | NULL     |                |
//| password_verify       | varchar(255)  | NO   |     | NULL     |                |
//| integrate_id          | int(11)       | NO   |     | NULL     |                |
//| sina_id               | int(11)       | NO   |     | NULL     |                |
//| renren_id             | int(11)       | NO   |     | NULL     |                |
//| kaixin_id             | int(11)       | NO   |     | NULL     |                |
//| sohu_id               | int(11)       | NO   |     | NULL     |                |
//| bind_verify           | varchar(255)  | NO   |     | NULL     |                |
//| verify_create_time    | int(11)       | NO   |     | NULL     |                |
//| tencent_id            | varchar(255)  | NO   |     | NULL     |                |
//| referer               | varchar(255)  | NO   |     | NULL     |                |
//| login_pay_time        | int(11)       | NO   |     | NULL     |                |
//| focus_count           | int(11)       | NO   |     | NULL     |                |
//| focused_count         | int(11)       | NO   |     | NULL     |                |
//| n_province_id         | int(11)       | NO   |     | NULL     |                |
//| n_city_id             | int(11)       | NO   |     | NULL     |                |
//| province_id           | int(11)       | NO   |     | NULL     |                |
//| city_id               | int(11)       | NO   |     | NULL     |                |
//| sex                   | tinyint(1)    | NO   |     | -1       |                |
//| step                  | tinyint(1)    | NO   |     | NULL     |                |
//| byear                 | int(4)        | NO   |     | NULL     |                |
//| bmonth                | int(2)        | NO   |     | NULL     |                |
//| bday                  | int(2)        | NO   |     | NULL     |                |
//| graduation            | varchar(15)   | NO   |     | NULL     |                |
//| graduatedyear         | int(5)        | NO   |     | NULL     |                |
//| university            | varchar(100)  | NO   |     | NULL     |                |
//| edu_validcode         | varchar(20)   | NO   |     | NULL     |                |
//| has_send_video        | tinyint(1)    | NO   |     | NULL     |                |
//| marriage              | varchar(15)   | NO   |     | NULL     |                |
//| haschild              | tinyint(1)    | NO   |     | NULL     |                |
//| hashouse              | tinyint(1)    | NO   |     | NULL     |                |
//| houseloan             | tinyint(1)    | NO   |     | NULL     |                |
//| hascar                | tinyint(1)    | NO   |     | NULL     |                |
//| carloan               | tinyint(4)    | NO   |     | NULL     |                |
//| car_brand             | varchar(50)   | NO   |     | NULL     |                |
//| car_year              | int(4)        | NO   |     | NULL     |                |
//| car_number            | varchar(50)   | NO   |     | NULL     |                |
//| address               | varchar(150)  | NO   |     | NULL     |                |
//| work_address          | varchar(100)  | YES  |     | NULL     |                |
//| phone                 | varchar(50)   | NO   |     | NULL     |                |
//| postcode              | varchar(20)   | NO   |     | NULL     |                |
//| locate_time           | int(11)       | NO   |     | 0        |                |
//| xpoint                | float(10,6)   | NO   |     | 0.000000 |                |
//| ypoint                | float(10,6)   | NO   |     | 0.000000 |                |
//| topic_count           | int(11)       | NO   |     | NULL     |                |
//| fav_count             | int(11)       | NO   |     | NULL     |                |
//| faved_count           | int(11)       | NO   |     | NULL     |                |
//| insite_count          | int(11)       | NO   |     | NULL     |                |
//| outsite_count         | int(11)       | NO   |     | NULL     |                |
//| level_id              | int(11)       | NO   |     | NULL     |                |
//| point                 | int(11)       | NO   |     | NULL     |                |
//| sina_app_key          | varchar(255)  | NO   |     | NULL     |                |
//| sina_app_secret       | varchar(255)  | NO   |     | NULL     |                |
//| is_syn_sina           | tinyint(1)    | NO   |     | NULL     |                |
//| tencent_app_key       | varchar(255)  | NO   |     | NULL     |                |
//| tencent_app_secret    | varchar(255)  | NO   |     | NULL     |                |
//| is_syn_tencent        | tinyint(1)    | NO   |     | NULL     |                |
//| t_access_token        | varchar(250)  | NO   |     | NULL     |                |
//| t_openkey             | varchar(250)  | NO   |     | NULL     |                |
//| t_openid              | varchar(250)  | NO   |     | NULL     |                |
//| sina_token            | varchar(250)  | NO   |     | NULL     |                |
//| is_borrow_out         | tinyint(1)    | NO   |     | NULL     |                |
//| is_borrow_int         | tinyint(1)    | NO   |     | NULL     |                |
//| creditpassed          | tinyint(1)    | NO   |     | NULL     |                |
//| creditpassed_time     | int(11)       | NO   |     | NULL     |                |
//| workpassed            | tinyint(1)    | NO   |     | NULL     |                |
//| workpassed_time       | int(11)       | NO   |     | NULL     |                |
//| incomepassed          | tinyint(1)    | NO   |     | NULL     |                |
//| incomepassed_time     | int(11)       | NO   |     | NULL     |                |
//| housepassed           | tinyint(1)    | NO   |     | NULL     |                |
//| housepassed_time      | int(11)       | NO   |     | NULL     |                |
//| carpassed             | tinyint(1)    | NO   |     | NULL     |                |
//| carpassed_time        | int(11)       | NO   |     | NULL     |                |
//| marrypassed           | tinyint(1)    | NO   |     | NULL     |                |
//| marrypassed_time      | int(11)       | NO   |     | NULL     |                |
//| edupassed             | tinyint(1)    | NO   |     | NULL     |                |
//| edupassed_time        | int(11)       | NO   |     | NULL     |                |
//| skillpassed           | tinyint(1)    | NO   |     | NULL     |                |
//| skillpassed_time      | int(11)       | NO   |     | NULL     |                |
//| videopassed           | tinyint(1)    | NO   |     | NULL     |                |
//| videopassed_time      | int(11)       | NO   |     | NULL     |                |
//| mobiletruepassed      | tinyint(1)    | NO   |     | NULL     |                |
//| mobiletruepassed_time | int(11)       | NO   |     | NULL     |                |
//| residencepassed       | tinyint(1)    | NO   |     | NULL     |                |
//| residencepassed_time  | int(11)       | NO   |     | NULL     |                |
//| alipay_id             | varchar(255)  | NO   |     | NULL     |                |
//| qq_id                 | varchar(255)  | NO   |     | NULL     |                |
//| info_down             | varchar(255)  | NO   |     | NULL     |                |
//| sealpassed            | tinyint(1)    | NO   |     | NULL     |                |
//| paypassword           | varchar(50)   | NO   |     | NULL     |                |
//| apns_code             | varchar(255)  | YES  |     | NULL     |                |
//| emailpassed           | tinyint(1)    | NO   |     | NULL     |                |
//| tmp_email             | varchar(255)  | NO   |     | NULL     |                |
//| view_info             | text          | NO   |     | NULL     |                |
//| ips_acct_no           | varchar(30)   | YES  |     | NULL     |                |
//| signature             | text          | YES  |     | NULL     |                |
type User struct {
	Id                   int64   `json:"id"  gorm:"column:id"`
	TypeId               int     `json:"type_id"  gorm:"column:type_id"`
	IsCompany            int     `json:"is_company"  gorm:"column:is_company"`
	CompanyName          string  `json:"company_name"  gorm:"column:company_name"`
	CompanyCode          string  `json:"company_code"  gorm:"column:company_code"`
	UserName             string  `json:"user_name"  gorm:"column:user_name"`
	UserPwd              string  `json:"user_pwd"  gorm:"column:user_pwd"`
	CreateTime           int64   `json:"create_time"  gorm:"column:create_time"`
	UpdateTime           int64   `json:"update_time"  gorm:"column:update_time"`
	RealTimeStatus       int     `json:"real_time_status"  gorm:"column:real_time_status"`
	CgCustNo             string  `json:"cg_custNo"  gorm:"column:cg_custNo"`
	CgAcctNo             string  `json:"cg_acctNo"  gorm:"column:cg_acctNo"`
	CgData               string  `json:"cg_data"  gorm:"column:cg_data"`
	CgAmtType            string  `json:"cg_amtType"  gorm:"column:cg_amtType"`
	LoginIp              string  `json:"login_ip"  gorm:"column:login_ip"`
	GroupId              int64   `json:"group_id"  gorm:"column:group_id"`
	IsEffect             int     `json:"is_effect"  gorm:"column:is_effect"`
	IsDelete             int     `json:"is_delete"  gorm:"column:is_delete"`
	Email                string  `json:"email"  gorm:"column:email"`
	Idno                 string  `json:"idno"  gorm:"column:idno"`
	Idcardpassed         int     `json:"idcardpassed"  gorm:"column:idcardpassed"`
	IdcardpassedTime     int     `json:"idcardpassed_time"  gorm:"column:idcardpassed_time"`
	IdcardVReturn        string  `json:"idcard_v_return"  gorm:"column:idcard_v_return"`
	RealName             string  `json:"real_name"  gorm:"column:real_name"`
	Mobile               string  `json:"mobile"  gorm:"column:mobile"`
	Mobilepassed         int     `json:"mobilepassed"  gorm:"column:mobilepassed"`
	EqbAccountId         string  `json:"eqb_accountId"  gorm:"column:eqb_accountId"`
	EqbSealData          string  `json:"eqb_sealData"  gorm:"column:eqb_sealData"`
	Score                int     `json:"score"  gorm:"column:score"`
	Money                float64 `json:"money"  gorm:"column:money"`
	Quota                float64 `json:"quota"  gorm:"column:quota"`
	LockMoney            float64 `json:"lock_money"  gorm:"column:lock_money"`
	HxMoney              float64 `json:"hx_money"  gorm:"column:hx_money"`
	HxLockMoney          float64 `json:"hx_lock_money"  gorm:"column:hx_lock_money"`
	Verify               string  `json:"verify"  gorm:"column:verify"`
	Code                 string  `json:"code"  gorm:"column:code"`
	Pid                  int64   `json:"pid"  gorm:"column:pid"`
	LoginTime            int64   `json:"login_time"  gorm:"column:login_time"`
	ReferralCount        int64   `json:"referral_count"  gorm:"column:referral_count"`
	PasswordVerify       string  `json:"password_verify"  gorm:"column:password_verify"`
	IntegrateId          int64   `json:"integrate_id"  gorm:"column:integrate_id"`
	SinaId               int64   `json:"sina_id"  gorm:"column:sina_id"`
	RenrenId             int64   `json:"renren_id"  gorm:"column:renren_id"`
	KaixinId             int64   `json:"kaixin_id"  gorm:"column:kaixin_id"`
	SohuId               int64   `json:"sohu_id"  gorm:"column:sohu_id"`
	BindVerify           string  `json:"bind_verify"  gorm:"column:bind_verify"`
	VerifyCreateTime     int64   `json:"verify_create_time"  gorm:"column:verify_create_time"`
	TencentId            string   `json:"tencent_id"  gorm:"column:tencent_id"`
	Referer              string  `json:"referer"  gorm:"column:referer"`
	LoginPayTime         int64   `json:"login_pay_time"  gorm:"column:login_pay_time"`
	FocusCount           int     `json:"focus_count"  gorm:"column:focus_count"`
	FocusedCount         int     `json:"focused_count"  gorm:"column:focused_count"`
	NProvinceId          int64   `json:"n_province_id"  gorm:"column:n_province_id"`
	NCityId              int64   `json:"n_city_id"  gorm:"column:n_city_id"`
	ProvinceId           int64   `json:"province_id"  gorm:"column:province_id"`
	CityId               int64   `json:"city_id"  gorm:"column:city_id"`
	Sex                  int     `json:"sex"  gorm:"column:sex"`
	Step                 int     `json:"step"  gorm:"column:step"`
	Byear                int     `json:"byear"  gorm:"column:byear"`
	Bmonth               int     `json:"bmonth"  gorm:"column:bmonth"`
	Bday                 int     `json:"bday"  gorm:"column:bday"`
	Graduation           string  `json:"graduation"  gorm:"column:graduation"`
	Graduatedyear        int     `json:"graduatedyear"  gorm:"column:graduatedyear"`
	University           string  `json:"university"  gorm:"column:university"`
	EeduValidcode        string  `json:"edu_validcode"  gorm:"column:edu_validcode"`
	HasSendVideo         int     `json:"has_send_video"  gorm:"column:has_send_video"`
	Marriage             string  `json:"marriage"  gorm:"column:marriage"`
	Haschild             int     `json:"haschild"  gorm:"column:haschild"`
	Hashouse             int     `json:"hashouse"  gorm:"column:hashouse"`
	Houseloan            int     `json:"houseloan"  gorm:"column:houseloan"`
	Hascar               int     `json:"hascar"  gorm:"column:hascar"`
	Carloan              int     `json:"hascar"  gorm:"column:hascar"`
	CarBrand             string  `json:"car_brand"  gorm:"column:car_brand"`
	CarYear              int     `json:"car_year"  gorm:"column:car_year"`
	CarNumber            string  `json:"car_number"  gorm:"column:car_number"`
	Aaddress             string  `json:"address"  gorm:"column:address"`
	WorkAddress          string  `json:"work_address"  gorm:"column:work_address"`
	Phone                string  `json:"phone"  gorm:"column:phone"`
	Postcode             string  `json:"postcode"  gorm:"column:postcode"`
	LocateTime           int64   `json:"locate_time"  gorm:"column:locate_time"`
	Xpoint               float32 `json:"xpoint"  gorm:"column:xpoint"`
	Ypoint               float32 `json:"ypoint"  gorm:"column:ypoint"`
	TopicCount           int     `json:"topic_count"  gorm:"column:topic_count"`
	FavCount             int     `json:"fav_count"  gorm:"column:fav_count"`
	FavedCount           int     `json:"faved_count"  gorm:"column:faved_count"`
	InsiteCount          int     `json:"insite_count"  gorm:"column:insite_count"`
	OoutsiteCount        int     `json:"outsite_count"  gorm:"column:outsite_count"`
	LevelId              int     `json:"level_id"  gorm:"column:level_id"`
	Point                int     `json:"point"  gorm:"column:point"`
	SinaAppKey           string  `json:"sina_app_key"  gorm:"column:sina_app_key"`
	SinaAppSecret        string  `json:"sina_app_secret"  gorm:"column:sina_app_secret"`
	IsSynSina            int     `json:"is_syn_sina"  gorm:"column:is_syn_sina"`
	TencentAppKey        string  `json:"tencent_app_key"  gorm:"column:tencent_app_key"`
	TencentAppSecret     string  `json:"tencent_app_secret"  gorm:"column:tencent_app_secret"`
	IsSynTencent         int     `json:"is_syn_tencent"  gorm:"column:is_syn_tencent"`
	TAccessToken         string  `json:"t_access_token"  gorm:"column:t_access_token"`
	TOpenkey             string  `json:"t_openkey"  gorm:"column:t_openkey"`
	TOpenid              string  `json:"t_openid"  gorm:"column:t_openid"`
	SinaToken            string  `json:"sina_token"  gorm:"column:sina_token"`
	IsBorrowOut          int     `json:"is_borrow_out"  gorm:"column:is_borrow_out"`
	IsBorrowInt          int     `json:"is_borrow_int"  gorm:"column:is_borrow_int"`
	Creditpassed         int     `json:"creditpassed"  gorm:"column:creditpassed"`
	CreditpassedTime     int64   `json:"creditpassed_time"  gorm:"column:creditpassed_time"`
	Workpassed           int     `json:"workpassed"  gorm:"column:workpassed"`
	WorkpassedTime       int64   `json:"workpassed_time"  gorm:"column:workpassed_time"`
	Incomepassed         int     `json:"incomepassed"  gorm:"column:incomepassed"`
	IncomepassedTime     int64   `json:"incomepassed_time"  gorm:"column:incomepassed_time"`
	Housepassed          int     `json:"housepassed"  gorm:"column:housepassed"`
	HousepassedTime      int     `json:"housepassed_time"  gorm:"column:housepassed_time"`
	Carpassed            int     `json:"carpassed"  gorm:"column:carpassed"`
	CarpassedTime        int64   `json:"carpassed_time"  gorm:"column:carpassed_time"`
	Marrypassed          int     `json:"marrypassed"  gorm:"column:marrypassed"`
	MarrypassedTime      int64   `json:"marrypassed_time"  gorm:"column:marrypassed_time"`
	Edupassed            int     `json:"edupassed"  gorm:"column:edupassed"`
	EdupassedTime        int64   `json:"edupassed_time"  gorm:"column:edupassed_time"`
	Skillpassed          int     `json:"skillpassed"  gorm:"column:skillpassed"`
	SkillpassedTime      int64   `json:"skillpassed_time"  gorm:"column:skillpassed_time"`
	Videopassed          int     `json:"videopassed"  gorm:"column:videopassed"`
	VideopassedTime      int64   `json:"videopassed_time"  gorm:"column:videopassed_time"`
	Mobiletruepassed     int     `json:"mobiletruepassed"  gorm:"column:mobiletruepassed"`
	MobiletruepassedTime int64   `json:"mobiletruepassed"  gorm:"column:mobiletruepassed"`
	Residencepassed      int     `json:"residencepassed"  gorm:"column:residencepassed"`
	ResidencepassedTime  int64   `json:"residencepassed"  gorm:"column:residencepassed"`
	AlipayId             string  `json:"alipay_id"  gorm:"column:alipay_id"`
	QqId                 string  `json:"qq_id"  gorm:"column:qq_id"`
	InfoDown             string  `json:"info_down"  gorm:"column:info_down"`
	Sealpassed           int     `json:"sealpassed"  gorm:"column:sealpassed"`
	Paypassword          string  `json:"paypassword"  gorm:"column:paypassword"`
	ApnsCode             string  `json:"apns_code"  gorm:"column:apns_code"`
	Emailpassed          int     `json:"emailpassed"  gorm:"column:emailpassed"`
	TmpEmail             string  `json:"tmp_email"  gorm:"column:tmp_email"`
	ViewInfo             string  `json:"view_info"  gorm:"column:view_info"`
	IpsAcctNo            string  `json:"ips_acct_no"  gorm:"column:ips_acct_no"`
	Signature            string  `json:"signature"  gorm:"column:signature"`
}


func (t User) TableName() string {
	return "nhyd_user"
}