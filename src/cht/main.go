package main

import (
	_ "cht/initial"
	"cht/service/couponlistthriftservice"
	"cht/service/loguserloginservice"
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
	<-ch
}
