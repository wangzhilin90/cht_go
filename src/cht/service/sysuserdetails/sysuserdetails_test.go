package sysuserdetails

import (
	_ "cht/initial"
	"testing"
)

func NewSysUserDetailsRequestStruct(user_id int32) *SysUserDetailsRequestStruct {
	return &SysUserDetailsRequestStruct{
		UserID: user_id,
	}
}

func TestGetSysUserDetails(t *testing.T) {
	sudr := NewSysUserDetailsRequestStruct(12)
	suds := sysuserdetailsservice{}
	res, _ := suds.GetSysUserDetails(sudr)
	if res.Status == QUERY_SYS_USER_DETAILS_FAILED {
		t.Fatalf("TestGetSysUserDetails query failed")
	}
	t.Logf("TestGetSysUserDetails return value:%v", res)
}
