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
	rars := NewRoleAddRequestStruct("fas", "niubi")
	ras := roleaddservice{}
	res, _ := ras.AddRole(rars)
	if res.Status != 1000 {
		t.Fatalf("TestAddRole insert failed")
	}
}
