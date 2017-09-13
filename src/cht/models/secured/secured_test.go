package secured

import (
	_ "cht/initial"
	"testing"
)

func NewSecuredRequestStruct() *SecuredRequestStruct {
	return &SecuredRequestStruct{}
}

func TestGetSecuredList(t *testing.T) {
	srs := NewSecuredRequestStruct()
	res, err := GetSecuredList(srs)
	if err != nil {
		t.Fatalf("TestGetSecuredList failed %v", err)
	}
	t.Logf("TestGetSecuredList res %v", res)
}
