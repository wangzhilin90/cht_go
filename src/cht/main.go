package main

import (
	_ "cht/initial"
	"cht/service/advertmanagethriftservice"
	"cht/service/articledetails"
	"cht/service/borrowuserdetails"
	"cht/service/dutydetails"
	"cht/service/emailattestationthriftservice"
	"cht/service/helplist"
	"cht/service/juanzengthriftservice"
	"cht/service/kefulist"
	"cht/service/makeborrowservice"
	"cht/service/messagethriftservice"
	"cht/service/operationaldata"
	"cht/service/phoneattestationthriftservice"
	"cht/service/roleadd"
	"cht/service/roledelete"
	"cht/service/roledetails"
	"cht/service/roleedit"
	"cht/service/rolerightset"
	"cht/service/securedlist"
	"cht/service/subledgerlist"
	"cht/service/sysconfigthriftservice"
	"cht/service/talking"
	"cht/service/updateuserloginlogdetails"
	"cht/service/updateuserpasswword"
	"cht/service/usercashrecordList"
	"cht/service/usercollectionlist"
	"cht/service/usercouponlist"
	"cht/service/userdetailsbynamepassword"
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
		usercashrecordList.StartCashRecordServer()
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

	go func() {
		/*开启查询后台用户详情服务*/
		userdetailsbynamepassword.StartUseDetailsrServer()
	}()

	go func() {
		/*开启查询客服列表服务*/
		kefulist.StartKeFuListsServer()
	}()

	go func() {
		/*开启后台---值班人详情服务*/
		dutydetails.StartDutyDetailServer()
	}()

	go func() {
		/*开启后台---用户角色添加服务*/
		roleadd.StartRoleAddServer()
	}()

	go func() {
		/*开启后台---用户角色删除服务*/
		roledelete.StartRoleDeleteServer()
	}()

	go func() {
		/*开启后台---用户角色详情获取服务*/
		roledetails.StartRoleDetailsServer()
	}()

	go func() {
		/*开启后台---用户角色编辑服务*/
		roleedit.StartRoleEditServer()
	}()

	go func() {
		/*开启后台---用户角色权限修改服务*/
		rolerightset.StartRoleEditServer()
	}()

	go func() {
		/*开启后台---小城交流日服务*/
		talking.StartTalkingServer()
	}()

	go func() {
		/*开启后台---运营数据查询服务*/
		operationaldata.StartOperationalDataServer()
	}()

	go func() {
		/*开启后台---获取帮助中心文章列表服务*/
		helplist.StartHelpListServer()
	}()

	go func() {
		/*开启后台---文章详情服务*/
		articledetails.StartArticleDetailsServer()
	}()

	<-ch
}
