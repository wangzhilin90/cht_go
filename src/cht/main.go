package main

import (
	_ "cht/initial"
	"cht/service/borrowerthriftservice"
	"cht/service/cashrecordthriftservice"
	"cht/service/collectionthriftservice"
	"cht/service/couponlistthriftservice"
	"cht/service/emailattestationthriftservice"
	"cht/service/gettendercouponthriftservice"
	"cht/service/gettenderredbagthriftservice"
	"cht/service/loguserloginservice"
	"cht/service/makeborrowservice"
	"cht/service/messagethriftservice"
	"cht/service/phoneattestationthriftservice"
	"cht/service/rechargerecordthriftservice"
	"cht/service/securedthriftservice"
	"cht/service/subledgerthriftservice"
	"cht/service/updatepasswdthriftservice"
	"cht/service/userloginservice"
)

func main() {
	ch := make(chan bool)

	go func() {
		/*开启加息券服务API*/
		couponlistthriftservice.StartCouponServer()
	}()

	go func() {
		/*开启用户登录服务*/
		userloginservice.StartUserLoginServer()
	}()

	go func() {
		/*开启登录日志服务*/
		loguserloginservice.StartLogUserLoginServer()
	}()

	go func() {
		/*开启忘记密码重置密码服务*/
		updatepasswdthriftservice.StartUpdatePasswdsServer()
	}()

	go func() {
		/*开启查询充值记录服务*/
		rechargerecordthriftservice.StartRechargeRecordServer()
	}()

	go func() {
		/*开启发标服务*/
		makeborrowservice.StartMakeBorrowServer()
	}()

	go func() {
		/*开启立即投资，获取红包金额服务*/
		gettenderredbagthriftservice.StartGetTenderRedBagServer()
	}()

	go func() {
		/*开启立即投资，获取投标加息值服务*/
		gettendercouponthriftservice.StartGetCouponServer()
	}()

	go func() {
		/*开启充值提现，获取提现记录服务*/
		cashrecordthriftservice.StartCashRecordServer()
	}()

	go func() {
		/*获取我的账户回款明细信息*/
		collectionthriftservice.StartGetCollectionListServer()
	}()

	go func() {
		/*开启短信服务，包括获取短信详情和短信记录数*/
		messagethriftservice.StartMessageServer()
	}()

	go func() {
		/*开启做标服务---分账人服务*/
		subledgerthriftservice.StartsubledgerServer()
	}()

	go func() {
		/*开启做标服务---担保人服务*/
		securedthriftservice.StartSecuredServer()
	}()

	go func() {
		/*开启做标服务---借款人服务*/
		borrowerthriftservice.StartCashRecordServer()
	}()

	go func() {
		/*开启手机密码重置服务*/
		phoneattestationthriftservice.StartPhoneAttestationServer()
	}()

	go func() {
		/*开启邮箱认证服务*/
		emailattestationthriftservice.StartEmailAttestationServer()
	}()
	<-ch
}
