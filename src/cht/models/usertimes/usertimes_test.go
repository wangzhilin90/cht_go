package usertimes

import (
	_ "cht/initial"
	"testing"
)

func NewUserTimesDetailsRequest(username string, isAdmin int32, type1 int32) *UserTimesDetailsRequest {
	return &UserTimesDetailsRequest{
		Username: username,
		Isadmin:  isAdmin,
		Type:     type1,
	}
}

func NewUserTimesUpdateRequest(username string, isAdmin int32, ip string, type1 int32) *UserTimesUpdateRequest {
	return &UserTimesUpdateRequest{
		Username: username,
		Isadmin:  isAdmin,
		IP:       ip,
		Type:     type1,
	}
}

func NewUserTimesInsertRequest(username string, isAdmin int32, ip string, type1 int32) *UserTimesInsertRequest {
	return &UserTimesInsertRequest{
		Username: username,
		Isadmin:  isAdmin,
		IP:       ip,
		Type:     type1,
	}
}

func NewUserTimesDeleteRequest(username string, type1 int32) *UserTimesDeleteRequest {
	return &UserTimesDeleteRequest{
		Username: username,
		Type:     type1,
	}
}

func TestGetUserTimesDetails(t *testing.T) {
	utdr := NewUserTimesDetailsRequest("wzl", 0, 1)
	res, err := GetUserTimesDetails(utdr)
	if err != nil {
		t.Fatalf("TestGetUserTimesDetails failed:%v", err)
	}
	t.Logf("TestGetUserTimesDetails return value:%v", res)
}

func TestUpdateUserTimes(t *testing.T) {
	utur := NewUserTimesUpdateRequest("wzl", 1, "192.168.8.208", 1)
	b := UpdateUserTimes(utur)
	if b == false {
		t.Fatalf("TestUpdateUserTimes failed")
	}
}

func TestInsertUserTimes(t *testing.T) {
	utir := NewUserTimesInsertRequest("wzl", 1, "192.168.8.209", 1)
	b := InsertUserTimes(utir)
	if b == false {
		t.Fatalf("TestInsertUserTimes failed")
	}
}

func TestDeleteUserTimes(t *testing.T) {
	utdr := NewUserTimesDeleteRequest("wzl", 1)
	b := DeleteUserTimes(utdr)
	if b == false {
		t.Fatalf("TestDeleteUserTimes failed")
	}
}
