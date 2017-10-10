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
	rdrs := NewRoleDeleteRequestStruct("124,125")
	b := DeleteRole(rdrs)
	if b == false {
		t.Fatalf("TestDeleteRole delete failed")
	}
}
