package subledgerthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewSubledgerRequestStruct(arrUser string) *SubledgerRequestStruct {
	return &SubledgerRequestStruct{
		HsZhuanrangrenStr: arrUser,
	}
}

func TestGetSubledgerList(t *testing.T) {
	srs := NewSubledgerRequestStruct("1,2,3,4444")
	ss := &subledgerservice{}
	res, err := ss.GetSubledgerList(srs)
	if err != nil {
		t.Fatalf("TestGetSubledgerList failed %v", err)
	}
	t.Logf("TestGetSubledgerList res:%v", res)
}
