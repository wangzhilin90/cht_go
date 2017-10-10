package roleadd

import (
	_ "cht/initial"
	"testing"
)

func NewRoleAddRequestStruct(rolename, remark string) *RoleAddRequestStruct {
	return &RoleAddRequestStruct{
		RoleName: rolename,
		Remark:   remark,
	}
}

func TestAddRole(t *testing.T) {
	rars := NewRoleAddRequestStruct("wzl", "niubi")
	b := AddRole(rars)
	if b == false {
		t.Fatalf("TestAddRole insert failed")
	}
}
