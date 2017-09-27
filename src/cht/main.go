package main

import (
	_ "cht/initial"
	"cht/service/advertmanagethriftservice"
	"cht/service/borrowuserdetails"
	"cht/service/cashrecordthriftservice"
	"cht/service/emailattestationthriftservice"
	"cht/service/juanzengthriftservice"
	"cht/service/makeborrowservice"
	"cht/service/messagethriftservice"
	"cht/service/phoneattestationthriftservice"
	"cht/service/securedlist"
	"cht/service/subledgerlist"
	"cht/service/sysconfigthriftservice"
	"cht/service/updateuserloginlogdetails"
	"cht/service/updateuserpasswword"
	"cht/service/usercollectionlist"
	"cht/service/usercouponlist"
	"cht/service/userloginservice"
	"cht/service/userrechargerecordlist"
	"cht/service/usertendercoupondetails"
	"cht/service/usertenderredbagdestails"
)

func main() {
	ch := make(chan bool)

	go func() {
		/*开启加息券服务API*/
		usercouponlist.StartCouponServer()
	}()

	go func() {
		/*开启用户登录服务*/
		userloginservice.StartUserLoginServer()
	}()

	go func() {
		/*用户登录记录日志服务*/
		updateuserloginlogdetails.StartLogUserLoginServer()
	}()

	go func() {
		/*修改用户密码服务*/
		updateuserpasswword.StartUpdatePasswdsServer()
	}()

	go func() {
		/*开启查询充值记录服务*/
		userrechargerecordlist.StartRechargeRecordServer()
	}()

	go func() {
		/*开启发标服务*/
		makeborrowservice.StartMakeBorrowServer()
	}()

	go func() {
		/*获取用户投资红包服务*/
		usertenderredbagdestails.StartGetTenderRedBagServer()
	}()

	go func() {
		/*获取用户投资加息值服务*/
		usertendercoupondetails.StartGetCouponServer()
	}()

	go func() {
		/*开启充值提现，获取提现记录服务*/
		cashrecordthriftservice.StartCashRecordServer()
	}()

	go func() {
		/*用户回款列表服务*/
		usercollectionlist.StartGetCollectionListServer()
	}()

	go func() {
		/*开启短信服务，包括获取短信详情和短信记录数*/
		messagethriftservice.StartMessageServer()
	}()

	go func() {
		/*做标服务---分账人列表服务*/
		subledgerlist.StartsubledgerServer()
	}()

	go func() {
		/*做标服务---担保人列表服务*/
		securedlist.StartSecuredServer()
	}()

	go func() {
		/*做标服务---借款人详情*/
		borrowuserdetails.StartBorrowerServer()
	}()

	go func() {
		/*开启手机密码重置服务*/
		phoneattestationthriftservice.StartPhoneAttestationServer()
	}()

	go func() {
		/*开启邮箱认证服务*/
		emailattestationthriftservice.StartEmailAttestationServer()
	}()

	go func() {
		/*开启查询系统配置服务*/
		sysconfigthriftservice.StartSysConfigServer()
	}()

	go func() {
		/*开启广告管理服务*/
		advertmanagethriftservice.StartAdvertManageServer()
	}()

	go func() {
		/*开启捐赠服务*/
		juanzengthriftservice.StartJuanzengServer()
	}()

	<-ch
}
