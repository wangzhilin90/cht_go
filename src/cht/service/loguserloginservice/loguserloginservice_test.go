package loguserloginservice

import (
	_ "cht/initial"
	"testing"
)

func TestUpdateLogUserlLogin(t *testing.T) {
	lulrs := &LogUserlLoginRequestStruct{
		UserID:               5003,
		LoginIP:              "192.168.8.35",
		LoginStyle:           0,
		ChengHuiTongTraceLog: "testlog",
	}
	luls := LogUserLoginService{}
	res, err := luls.UpdateLogUserlLogin(lulrs)
	if err != nil || res.UserID != 5003 {
		t.Fatalf("TestUpdateLogUserlLogin failed")
	}
}
