package usertimes

import (
	_ "cht/initial"
	"testing"
)

func NewUserTimesDetailsRequestStruct(username string, isAdmin int32, type1 int32) *UserTimesDetailsRequestStruct {
	return &UserTimesDetailsRequestStruct{
		Username: username,
		Isadmin:  isAdmin,
		Type:     type1,
	}
}

func NewUserTimesUpdateRequestStruct(username string, isAdmin int32, ip string, type1 int32) *UserTimesUpdateRequestStruct {
	return &UserTimesUpdateRequestStruct{
		Username: username,
		Isadmin:  isAdmin,
		IP:       ip,
		Type:     type1,
	}
}

func NewUserTimesInsertRequestStruct(username string, isAdmin int32, ip string, type1 int32) *UserTimesInsertRequestStruct {
	return &UserTimesInsertRequestStruct{
		Username: username,
		Isadmin:  isAdmin,
		IP:       ip,
		Type:     type1,
	}
}

func NewUserTimesDeleteRequestStruct(username string, type1 int32) *UserTimesDeleteRequestStruct {
	return &UserTimesDeleteRequestStruct{
		Username: username,
		Type:     type1,
	}
}

func TestGetUserTimesDetails(t *testing.T) {
	utdr := NewUserTimesDetailsRequestStruct("wzl", 0, 1)
	uts := usertimesservice{}
	res, _ := uts.GetUserTimesDetails(utdr)
	if res.Status != QUERY_USER_TIMES_SUCCESS {
		t.Fatalf("TestGetUserTimesDetails failed")
	}
	t.Logf("TestGetUserTimesDetails return value:%v", res)
}

func TestUpdateUserTimes(t *testing.T) {
	utur := NewUserTimesUpdateRequestStruct("wzl", 1, "192.168.8.208", 1)
	uts := usertimesservice{}
	res, _ := uts.UpdateUserTimes(utur)
	if res.Status != UPDATE_USER_TIMES_SUCCESS {
		t.Fatalf("TestUpdateUserTimes failed")
	}
}

func TestInsertUserTimes(t *testing.T) {
	utir := NewUserTimesInsertRequestStruct("wzl", 1, "192.168.8.209", 1)
	uts := usertimesservice{}
	res, _ := uts.InsertUserTimes(utir)
	if res.Status != INSERT_USER_TIMES_SUCCESS {
		t.Fatalf("TestInsertUserTimes failed")
	}
}

func TestDeleteUserTimes(t *testing.T) {
	utdrs := NewUserTimesDeleteRequestStruct("wzl", 1)
	uts := usertimesservice{}
	res, _ := uts.DeleteUserTimes(utdrs)
	if res.Status != DELETE_USER_TIMES_SUCCESS {
		t.Fatalf("TestDeleteUserTimes failed")
	}
}
