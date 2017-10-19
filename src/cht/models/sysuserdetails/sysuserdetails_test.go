package sysuserdetails

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserDetailsRequest(user_id int32) *SysUserDetailsRequest {
	return &SysUserDetailsRequest{
		UserID: user_id,
	}
}

func TestGetSysUserDetails(t *testing.T) {
	sudr := NewSysUserDetailsRequest(12)
	res, err := GetSysUserDetails(sudr)
	if err != nil {
		t.Fatalf("TestGetSysUserDetails query failed :%v", err)
	}
	t.Logf("TestGetSysUserDetails return value:%v", res)
}
