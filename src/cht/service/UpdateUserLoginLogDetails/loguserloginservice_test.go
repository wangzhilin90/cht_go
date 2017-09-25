package UpdateUserLoginLogDetails

import (
	_ "cht/initial"
	"testing"
)

func TestUpdateLogUserlLogin(t *testing.T) {
	lulrs := &LogUserlLoginRequestStruct{
		UserID:               28,
		LoginIP:              "192.168.8.36",
		LoginStyle:           0,
		ChengHuiTongTraceLog: "testlog",
	}
	luls := LogUserLoginService{}
	res, err := luls.UpdateLogUserlLogin(lulrs)
	if err != nil {
		t.Fatalf("TestUpdateLogUserlLogin failed")
	}
	t.Log("TestUpdateLogUserlLogin res:", res)
}
