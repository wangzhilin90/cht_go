package loguserlogin

import (
	_ "cht/initial"
	"testing"
)

func NewLogUserlLoginRequest(userID int32, loginIP string, loginstyle int32, log string) *LogUserlLoginRequest {
	lulr := &LogUserlLoginRequest{
		UserID:               userID,
		LoginIP:              loginIP,
		LoginStyle:           loginstyle,
		ChengHuiTongTraceLog: log,
	}
	return lulr
}

func NewUserLoginLogDetailsRequest(user_id int32) *UserLoginLogDetailsRequest {
	return &UserLoginLogDetailsRequest{
		UserID: user_id,
	}
}

func TestGetBorrowInfo(t *testing.T) {
	lulr := NewLogUserlLoginRequest(28, "192.168.8.35", 0, "testlog")
	res, err := GetBorrowInfo(lulr)
	if err != nil {
		t.Fatalf("GetBorrowInfo failed", err)
	}
	t.Logf("GetBorrowInfo res", res)
}

func TestUpdateLogUserlLogin(t *testing.T) {
	lulr := NewLogUserlLoginRequest(28, "192.168.8.35", 0, "testlog")
	b, _ := UpdateLogUserlLogin(lulr)
	if b == false {
		t.Fatalf("TestUpdateLogUserlLogin failed")
	}
}

func TestGetUserLoginLogDetails(t *testing.T) {
	ulldr := NewUserLoginLogDetailsRequest(44428)
	res, err := GetUserLoginLogDetails(ulldr)
	if err != nil {
		t.Fatalf("TestGetUserLoginLogDetails failed:%v", err)
	}
	t.Logf("TestGetUserLoginLogDetails return value:%v", res)
}
