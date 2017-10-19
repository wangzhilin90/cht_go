package sysuseredit

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserEditRequestStruct(account, mobile string, user_id int32) *SysUserEditRequestStruct {
	return &SysUserEditRequestStruct{
		Account: account,
		Mobile:  mobile,
		UserID:  user_id,
	}
}

func TestEditSysUser(t *testing.T) {
	suer := NewSysUserEditRequestStruct("nini", "8012", 12)
	sues := sysusereditservice{}
	res, _ := sues.EditSysUser(suer)
	if res.Status == SYS_USER_EDIT_FAILED {
		t.Fatalf("TestEditSysUser failed")
	}
}
