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
	rers := NewRoleEditRequestStruct(126, "hua", "lihai")
	b := EditRole(rers)
	if b == false {
		t.Fatalf("TestEditRole edit failed")
	}
}
