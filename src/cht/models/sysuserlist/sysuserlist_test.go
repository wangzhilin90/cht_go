package sysuserlist

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserListRequest() *SysUserListRequest {
	return &SysUserListRequest{}
}

func TestGetSysUserList(t *testing.T) {
	sulr := NewSysUserListRequest()
	res, err := GetSysUserList(sulr)
	if err != nil {
		t.Fatalf("TestGetSysUserList query failed :%v", err)
	}
	t.Logf("TestGetSysUserList return value:%v", res)
}
