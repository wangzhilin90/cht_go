package userdetailsbynamepassword

import (
	_ "cht/initial"
	"testing"
)

func NewUserDetailsByNamePasswordRequestStruct(name, passwd string) *UserDetailsByNamePasswordRequestStruct {
	return &UserDetailsByNamePasswordRequestStruct{
		Name:     name,
		Password: passwd,
	}
}

func TestGetUseDetailsrByNamePassword(t *testing.T) {
	udbpr := NewUserDetailsByNamePasswordRequestStruct("keke", "2b4e2374c6c2afa0b88bc67c43faa499")
	udbs := userdetailsbynamepasswordservice{}
	res, err := udbs.GetUseDetailsrByNamePassword(udbpr)
	if err != nil {
		t.Fatalf("TestGetUseDetailsrByNamePassword failed", err)
	}
	t.Logf("TestGetUseDetailsrByNamePassword res :%v", res)
}
