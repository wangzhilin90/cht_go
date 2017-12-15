package usertimes

import (
	_ "cht/initial"
	"testing"
)

func NewUserTimesDetailsRequest(username string, isAdmin int32) *UserTimesDetailsRequest {
	return &UserTimesDetailsRequest{
		Username: username,
		Isadmin:  isAdmin,
	}
}

func NewUserTimesUpdateRequest(username string, isAdmin int32, times int32, ip string) *UserTimesUpdateRequest {
	return &UserTimesUpdateRequest{
		Username: username,
		Isadmin:  isAdmin,
		Times:    times,
		IP:       ip,
	}
}

func NewUserTimesInsertRequest(username string, isAdmin int32, times int32, ip string) *UserTimesInsertRequest {
	return &UserTimesInsertRequest{
		Username: username,
		Isadmin:  isAdmin,
		Times:    times,
		IP:       ip,
	}
}

func NewUserTimesDeleteRequest(username string) *UserTimesDeleteRequest {
	return &UserTimesDeleteRequest{
		Username: username,
	}
}

func TestGetUserTimesDetails(t *testing.T) {
	utdr := NewUserTimesDetailsRequest("wzl", 0)
	res, err := GetUserTimesDetails(utdr)
	if err != nil {
		t.Fatalf("TestGetUserTimesDetails failed:%v", err)
	}
	t.Logf("TestGetUserTimesDetails return value:%v", res)
}

func TestUpdateUserTimes(t *testing.T) {
	utur := NewUserTimesUpdateRequest("wzl", 1, 3, "192.168.8.208")
	b := UpdateUserTimes(utur)
	if b == false {
		t.Fatalf("TestUpdateUserTimes failed")
	}
}

func TestInsertUserTimes(t *testing.T) {
	utir := NewUserTimesInsertRequest("wzl", 1, 1, "192.168.8.209")
	b := InsertUserTimes(utir)
	if b == false {
		t.Fatalf("TestInsertUserTimes failed")
	}
}

func TestDeleteUserTimes(t *testing.T) {
	utdr := NewUserTimesDeleteRequest("wzl")
	b := DeleteUserTimes(utdr)
	if b == false {
		t.Fatalf("TestDeleteUserTimes failed")
	}
}
