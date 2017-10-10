package roledetails

import (
	_ "cht/initial"
	"testing"
)

func NewRoleDetailsRequestStruct(roleid int32) *RoleDetailsRequestStruct {
	return &RoleDetailsRequestStruct{
		RoleID: roleid,
	}
}

func TestGetRoleDetails(t *testing.T) {
	rdrs := NewRoleDetailsRequestStruct(128)
	rds := roledetailsservice{}
	res, _ := rds.GetRoleDetails(rdrs)
	if res.Status != 1000 {
		t.Errorf("TestGetRoleDetails query failed")
	}
	t.Logf("TestGetRoleDetails return value %v", res)
}
