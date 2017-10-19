package sysuserlist

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserListRequestStruct() *SysUserListRequestStruct {
	return &SysUserListRequestStruct{}
}

func TestGetSysUserList(t *testing.T) {
	sulrs := NewSysUserListRequestStruct()
	suls := sysuserlistservice{}
	res, _ := suls.GetSysUserList(sulrs)
	if res.Status == QUERY_SYS_USER_LIST_FAILED {
		t.Fatalf("TestGetSysUserList query failed")
	}
	t.Logf("TestGetSysUserList return value:%v", res)
}
