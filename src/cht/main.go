package main

import (
	_ "cht/initial"
	"cht/service/couponlistthriftservice"
	"cht/service/loguserloginservice"
	"cht/service/rechargerecordthriftservice"
	"cht/service/updatepasswdthriftservice"
	"cht/service/userloginservice"
)

func main() {
	ch := make(chan bool)
	go func() {
		couponlistthriftservice.StartCouponServer()
	}()

	go func() {
		userloginservice.StartUserLoginServer()
	}()

	go func() {
		loguserloginservice.StartLogUserLoginServer()
	}()

	go func() {
		updatepasswdthriftservice.StartUpdatePasswdsServer()
	}()

	go func() {
		rechargerecordthriftservice.StartRechargeRecordServer()
	}()

	<-ch
}
