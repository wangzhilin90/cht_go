package usertimes

import (
	_ "cht/initial"
	"testing"
)

func NewUserTimesDetailsRequestStruct(username string, isAdmin int32) *UserTimesDetailsRequestStruct {
	return &UserTimesDetailsRequestStruct{
		Username: username,
		Isadmin:  isAdmin,
	}
}

func NewUserTimesUpdateRequestStruct(username string, isAdmin int32, times int32, ip string) *UserTimesUpdateRequestStruct {
	return &UserTimesUpdateRequestStruct{
		Username: username,
		Isadmin:  isAdmin,
		Times:    times,
		IP:       ip,
	}
}

func NewUserTimesInsertRequestStruct(username string, isAdmin int32, times int32, ip string) *UserTimesInsertRequestStruct {
	return &UserTimesInsertRequestStruct{
		Username: username,
		Isadmin:  isAdmin,
		Times:    times,
		IP:       ip,
	}
}

func TestGetUserTimesDetails(t *testing.T) {
	utdr := NewUserTimesDetailsRequestStruct("wzl", 0)
	uts := usertimesservice{}
	res, _ := uts.GetUserTimesDetails(utdr)
	if res.Status != QUERY_USER_TIMES_SUCCESS {
		t.Fatalf("TestGetUserTimesDetails failed")
	}
	t.Logf("TestGetUserTimesDetails return value:%v", res)
}

func TestUpdateUserTimes(t *testing.T) {
	utur := NewUserTimesUpdateRequestStruct("wzl", 1, 3, "192.168.8.208")
	uts := usertimesservice{}
	res, _ := uts.UpdateUserTimes(utur)
	if res.Status != UPDATE_USER_TIMES_SUCCESS {
		t.Fatalf("TestUpdateUserTimes failed")
	}
}

func TestInsertUserTimes(t *testing.T) {
	utir := NewUserTimesInsertRequestStruct("wzl1", 1, 1, "192.168.8.209")
	uts := usertimesservice{}
	res, _ := uts.InsertUserTimes(utir)
	if res.Status != INSERT_USER_TIMES_SUCCESS {
		t.Fatalf("TestInsertUserTimes failed")
	}
}
