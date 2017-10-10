package roledelete

import (
	_ "cht/initial"
	"testing"
)

func NewRoleDeleteRequestStruct(roleid string) *RoleDeleteRequestStruct {
	return &RoleDeleteRequestStruct{
		RoleIDStr: roleid,
	}
}

func TestDeleteRole(t *testing.T) {
	rdrs := NewRoleDeleteRequestStruct("126,127")
	rds := roledeleteservice{}
	res, _ := rds.DeleteRole(rdrs)
	if res.Status != 1000 {
		t.Fatalf("TestDeleteRole delete failed")
	}
}
