package rolerightset

import (
	_ "cht/initial"
	"testing"
)

func NewRoleRightSetRequestStruct(roleId int32, powerConfig string) *RoleRightSetRequestStruct {
	return &RoleRightSetRequestStruct{
		RoleID:      roleId,
		PowerConfig: powerConfig,
	}
}

func TestSetRoleRight(t *testing.T) {
	rrsrs := NewRoleRightSetRequestStruct(126, "12,24")
	b := SetRoleRight(rrsrs)
	if b == false {
		t.Fatalf("TestSetRoleRight update failed")
	}
}
