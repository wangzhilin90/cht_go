package sysuseradd

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserAddRequestStruct() *SysUserAddRequestStruct {
	return &SysUserAddRequestStruct{
		Account:     "17",
		RoleID:      3442,
		Password:    "mima",
		Realname:    " ",
		Mobile:      " ",
		Qq:          " ",
		Lastloginip: " ",
	}
}

func TestAddSysUser(t *testing.T) {
	suars := NewSysUserAddRequestStruct()
	suas := sysuseraddservice{}
	res, _ := suas.AddSysUser(suars)
	if res.Status == INSERT_SYS_USER_FAILED {
		t.Fatalf("TestAddSysUser failed")
	}
}
