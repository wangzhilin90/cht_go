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
	rrsrs := NewRoleRightSetRequestStruct(128, "12,24")
	rrs := rolerightsetservice{}
	res, _ := rrs.SetRoleRight(rrsrs)
	if res.Status != 1000 {
		t.Fatalf("TestSetRoleRight update failed")
	}
}
