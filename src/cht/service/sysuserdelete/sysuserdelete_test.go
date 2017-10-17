package sysuserdelete

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserDeleteRequestStruct() *SysUserDeleteRequestStruct {
	return &SysUserDeleteRequestStruct{
		UserIDStr: "207,208",
	}
}

func TestDeleteSysUser(t *testing.T) {
	sudrs := NewSysUserDeleteRequestStruct()
	suds := sysuserdeleteservice{}
	res, _ := suds.DeleteSysUser(sudrs)
	if res.Status == DELETE_SYS_USER_FAILED {
		t.Fatalf("TestDeleteSysUser failed")
	}
}
