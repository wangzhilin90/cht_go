package subledger

import (
	_ "cht/initial"
	"testing"
)

func NewSubledgerRequest(arrUser string) *SubledgerRequest {
	return &SubledgerRequest{
		HsZhuanrangrenStr: arrUser,
	}
}

func TestGetSubledgerList(t *testing.T) {
	sr := NewSubledgerRequest("143242314")
	res, err := GetSubledgerList(sr)
	if err != nil {
		t.Fatalf("TestGetSubledgerList failed %v", err)
	}
	t.Logf("TestGetSubledgerList res %v", res)
}
