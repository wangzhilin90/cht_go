package updateuserloginlogdetails

import (
	_ "cht/initial"
	"testing"
)

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
