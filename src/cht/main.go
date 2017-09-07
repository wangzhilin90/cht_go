package main

import (
	_ "cht/initial"
	"cht/service/cashrecordthriftservice"
	"cht/service/couponlistthriftservice"
	"cht/service/gettendercouponthriftservice"
	"cht/service/gettenderredbagthriftservice"
	"cht/service/loguserloginservice"
	"cht/service/makeborrowservice"
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

	go func() {
		makeborrowservice.StartMakeBorrowServer()
	}()

	go func() {
		gettenderredbagthriftservice.StartGetTenderRedBagServer()
	}()

	go func() {
		gettendercouponthriftservice.StartGetCouponServer()
	}()

	go func() {
		cashrecordthriftservice.StartCashRecordServer()
	}()

	<-ch
}
