package roleedit

import (
	_ "cht/initial"
	"testing"
)

func NewRoleEditRequestStruct(roledId int32, roleName, remark string) *RoleEditRequestStruct {
	return &RoleEditRequestStruct{
		RoleID:   roledId,
		RoleName: roleName,
		Remark:   remark,
	}
}

func TestEditRole(t *testing.T) {
	rers := NewRoleEditRequestStruct(128, "huaa", "lihai")
	rs := roleeditservice{}
	res, _ := rs.EditRole(rers)
	if res.Status != 1000 {
		t.Fatalf("TestEditRole edit failed")
	}
}
