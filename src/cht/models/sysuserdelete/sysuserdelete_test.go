package sysuserdelete

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserDeleteRequest() *SysUserDeleteRequest {
	return &SysUserDeleteRequest{
		UserIDStr: "210,211",
	}
}

func TestDeleteSysUser(t *testing.T) {
	sudr := NewSysUserDeleteRequest()
	b := DeleteSysUser(sudr)
	if b == false {
		t.Fatalf("TestDeleteSysUser failed")
	}
}
