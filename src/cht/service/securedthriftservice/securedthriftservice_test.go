package securedthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewSecuredRequestStruct() *SecuredRequestStruct {
	return &SecuredRequestStruct{}
}

func TestGetSecuredList(t *testing.T) {
	srs := NewSecuredRequestStruct()
	ss := &securedservice{}
	res, err := ss.GetSecuredList(srs)
	if err != nil {
		t.Fatalf("TestGetSecuredList failed %v", res)
	}
	t.Logf("TestGetSecuredList res:%v", res)
}
