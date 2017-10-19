package sysuseredit

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserEditRequest(account, mobile string, user_id int32) *SysUserEditRequest {
	return &SysUserEditRequest{
		Account: account,
		Mobile:  mobile,
		UserID:  user_id,
	}
}

func TestEditSysUser(t *testing.T) {
	suer := NewSysUserEditRequest("nini", "8012", 12)
	b := EditSysUser(suer)
	if b == false {
		t.Fatalf("TestEditSysUser failed")
	}
}
