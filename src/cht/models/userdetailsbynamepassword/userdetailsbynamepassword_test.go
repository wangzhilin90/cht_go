package userdetailsbynamepassword

import (
	_ "cht/initial"
	"testing"
)

func NewUserDetailsByNamePasswordRequest(name, passwd string) *UserDetailsByNamePasswordRequest {
	return &UserDetailsByNamePasswordRequest{
		Name:     name,
		Password: passwd,
	}
}

func TestGetUseDetailsrByNamePassword(t *testing.T) {
	udbpr := NewUserDetailsByNamePasswordRequest("keke", "2b4e2374c6c2afa0b88bc67c43faa499")
	res, err := GetUseDetailsrByNamePassword(udbpr)
	if err != nil {
		t.Fatalf("TestGetUseDetailsrByNamePassword failed", err)
	}
	t.Logf("TestGetUseDetailsrByNamePassword res :%v", res)
}
