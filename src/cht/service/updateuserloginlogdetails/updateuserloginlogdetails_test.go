package updateuserloginlogdetails

import (
	_ "cht/initial"
	"testing"
)

func NewUserLoginLogDetailsRequestStruct(user_id int32) *UserLoginLogDetailsRequestStruct {
	return &UserLoginLogDetailsRequestStruct{
		UserID: user_id,
	}
}

func TestUpdateUserLoginLogDetails(t *testing.T) {
	lulrs := &UpdateUserLoginLogDetailsRequestStruct{
		UserID:               31,
		LoginIP:              "192.168.8.36",
		LoginStyle:           0,
		ChengHuiTongTraceLog: "testlog",
	}
	luls := LogUserLoginService{}
	res, err := luls.UpdateUserLoginLogDetails(lulrs)
	if err != nil {
		t.Fatalf("TestupdateUserLoginLogDetails failed")
	}
	t.Log("TestupdateUserLoginLogDetails res:", res)
}

func TestGetUserLoginLogDetails(t *testing.T) {
	ulldr := NewUserLoginLogDetailsRequestStruct(44428)
	luls := LogUserLoginService{}
	res, _ := luls.GetUserLoginLogDetails(ulldr)
	if res.Status != QUERY_USER_LOGIN_LOG_SUCCESS {
		t.Fatalf("TestGetUserLoginLogDetails failed")
	}
	t.Logf("TestGetUserLoginLogDetails return value:%v", res)
}
