package controllers

import (
	"bili-gin/entitys"
	"bili-gin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const Shxydm  = "91450100092739725T"

func AllProjeck(c *gin.Context){
	db :=util.GetInstance().MyDB()
	var deel entitys.Deel
	var deelsForTime []entitys.Deel
	var allDeels []entitys.Deel

	var data [][]string

	db.Table(deel.TableName()).Select("repay_start_time ").Where("deal_status>=0 and publish_wait=0 and is_effect=1 and is_delete=0").Group("FROM_UNIXTIME(repay_start_time,'%Y-%m-%d')").Find(&deelsForTime)//按天分区
	if len(deelsForTime) >0 {

		for _,times :=range deelsForTime{
			timeStart,err :=util.DateStart(times.RepayStartTime)
			if err !=nil{
				fmt.Println("get time start err",err.Error())
			}
			timeEnd,errs := util.DateEnd(times.RepayStartTime)
			if errs != nil {
				fmt.Println("get time end err",errs.Error())
			}

			db.Table(deel.TableName()).Select("id,repay_start_time,last_repay_time,borrow_amount,rate,services_fee,manage_fee,guarantor_real_freezen_amt,loantype,type_id,buy_count,repay_time,user_id").Where("deal_status>=0 and publish_wait=0 and is_effect=1 and is_delete=0 and repay_start_time>=? and repay_start_time<=?",timeStart,timeEnd).Find(&allDeels)//获取每天的订单数

			if len(allDeels)>0{


				//处理订单
				for _,deel := range allDeels{

					 var info []string
					/*
					 $data = [];
                        $data['5001'] = $this->shxydm . $deal['id'];//项目唯一编号
                        $data['5002'] = to_date($deal['repay_start_time'], 'Ymd');//借款项目成立的日期，格式为 YYYYMMDD
                        $data['5003'] = $this->shxydm;//社会信用代码
                        $data['5004'] = $deal['id'];//项目编号
                        $data['5005'] = $this->number_float(4, $deal['borrow_amount']);//借款金额
                        $data['5006'] = to_date($deal['repay_start_time'], 'Ymd');//借款起息日
                        $last_repay = M('DealRepay')->field('repay_time')->where(array('deal_id' => $deal['id']))->order('id desc')->find();
                        $data['5007'] = to_date($last_repay['repay_time'], 'Ymd');//借款到期日期
                        $data['5008'] = $this->number_float(4, $deal['rate'] * 0.01);// 出借利率
                        $data['5009'] = $this->number_float(4, $deal['borrow_amount'] * $deal['manage_fee'] * 0.01);// 借款项目平台服务费 保留小数点后 4 位，不够 4 位的补 0
                        $data['5010'] = $this->number_float(4, $deal['guarantor_real_freezen_amt']);// 第三方代偿保障费用 保留小数点后 4 位，不够 4 位的补 0
                        $data['5011'] = info[$deal['loantype']];//还款方式
                        $data['5012'] = info[$deal['type_id']];//借款用途
                        $data['5013'] = $deal['buy_count'];//出借人个数
                        $data['5014'] = $deal['repay_time'];//约定还款总期数
                        $all_data[] = $data;
					*/
					var repay entitys.DealRepay
					db.Table(repay.TableName()).Select("repay_time").Where("deal_id=?",deel.Id).Order("id asc").Last(&repay)
					info = append(info, fmt.Sprintf("%s%d",Shxydm,deel.Id),util.StempToTime(deel.RepayStartTime),Shxydm,strconv.FormatInt(deel.Id, 10),fmt.Sprintf("%.4f",deel.BorrowAmount),util.StempToTime(deel.RepayStartTime),util.StempToTime(repay.RepayTime),fmt.Sprintf("%.4f",deel.Rate*0.01),fmt.Sprintf("%.4f",deel.Rate*0.01*deel.BorrowAmount),fmt.Sprintf("%.4f",deel.GuarantorRealFreezenAmt),loanType(deel.Loantype),typeId(deel.TypeId),strconv.FormatInt(deel.BuyCount, 10),strconv.FormatInt(deel.RepayTime, 10))

					data = append(data, info)
				}

				er :=util.WriteFile("data/24EXPORTBUSINESSZHAIQ.csv",data)
				if er !=nil{
					fmt.Println("写入文件异常：",er.Error())
				}


			}

		}
	}

	c.JSON(200, gin.H{
		"code":200,
		"msg": "ok",
		"data":data,

	})
}

func AllProjeckBytine(c *gin.Context)  {

	db :=util.GetInstance().MyDB()
	var deel entitys.Deel
	var deels []entitys.Deel
	var times = make(chan int64)
	db.Table(deel.TableName()).Select("repay_start_time ").Where("deal_status>=0 and publish_wait=0 and is_effect=1 and is_delete=0").Group("FROM_UNIXTIME(repay_start_time,'%Y-%m-%d')").Find(&deels)

	go deelTimes(times)

	for _,v:=range deels{
		times <-v.RepayStartTime
	}

	close(times)

	//deelData :=make([]map[string]interface{},len(deels))
	c.JSON(200, gin.H{
		"code":200,
		"msg": "ok",
		"data":deels,

	})
}

func deelTimes(time  chan int64) (times *[]map[string]string){
	for v:=range time{
		fmt.Println("\r\n时间:",v)
	}
	return times
}
//借款用途
func typeId(typeid int) string {
	if typeid>10{
		return "01"
	}
	info :=make(map[int]string,10)
	info[1] = "02"//'短期周转';
	info[2] = "04"//'购房借款';
	info[3] = "01"//'装修借款';
	info[4] = "01"//'个人消费';
	info[5] = "01"//'婚礼筹备';
	info[6] = "03"//'教育培训';
	info[7] = "01"//'汽车消费';
	info[8] = "05"//'投资创业';
	info[9] = "03"//'医疗支出';
	info[10] = "01"//'其他借款';
  	return info[typeid]
}

//还款方式
func loanType(loantype int) string {
	if loantype>5{
		return "04"
	}

	info :=make(map[int]string,10)
	info[0] = "01"//等额本息
	info[1] = "04"//先息后本
	info[2] = "05"//一次性还本付息
	info[3] = "03"//等本等息
	info[4] = "02"//等额本金
	info[5] = "06"//随时还款
	return info[loantype]
}