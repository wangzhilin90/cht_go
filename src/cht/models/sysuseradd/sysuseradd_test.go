package sysuseradd

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserAddRequest() *SysUserAddRequest {
	return &SysUserAddRequest{
		Account:     "15",
		RoleID:      3442,
		Password:    "mima",
		Realname:    " ",
		Mobile:      " ",
		Qq:          " ",
		Lastloginip: " ",
	}
}

func TestAddSysUser(t *testing.T) {
	suars := NewSysUserAddRequest()
	b := AddSysUser(suars)
	if b == false {
		t.Fatalf("TestAddSysUser failed")
	}
}
