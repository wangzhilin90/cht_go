package kefulist

import (
	_ "cht/initial"
	"testing"
)

func NewKeFuListRequestStruct(roleId int32, status int32, customerType string) *KeFuListRequestStruct {
	return &KeFuListRequestStruct{
		RoleID:       roleId,
		Status:       status,
		CustomerType: customerType,
	}
}

func TestGetKeFuList(t *testing.T) {
	kfrs := NewKeFuListRequestStruct(2, 0, "2,3")
	res, err := GetKeFuList(kfrs)
	if err != nil {
		t.Fatalf("TestGetKeFuList failed %v", err)
	}
	t.Logf("TestGetKeFuList res:%v", res)
}
