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
	rdrs := NewRoleDetailsRequestStruct(126)
	res, err := GetRoleDetails(rdrs)
	if err != nil {
		t.Errorf("TestGetRoleDetails query failed %v", err)
	}
	t.Logf("TestGetRoleDetails return value %v", res)
}
