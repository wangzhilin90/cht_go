package sysconfig

import (
	_ "cht/initial"
	"testing"
)

func NewSysConfigRequestStruct() *SysConfigRequestStruct {
	return &SysConfigRequestStruct{}
}

func TestGetSysConfig(t *testing.T) {
	scrs := NewSysConfigRequestStruct()
	res, err := GetSysConfig(scrs)
	if err != nil {
		t.Fatalf("TestGetSysConfig failed")
	}
	t.Logf("TestGetSysConfig res %v", res)
}
