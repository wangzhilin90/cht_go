package memberhelperlist

import (
	_ "cht/initial"
	"testing"
)

func NewMemberHelperListRequest(utype int32, keywords string) *MemberHelperListRequest {
	return &MemberHelperListRequest{
		Type:     utype,
		Keywords: keywords,
	}
}

func TestGetMemberHelperListTotalNum(t *testing.T) {
	mhlr := NewMemberHelperListRequest(2, "ha")
	num, err := GetMemberHelperListTotalNum(mhlr)
	if err != nil {
		t.Fatalf("TestGetMemberHelperListTotalNum query failed:%v", err)
	}
	t.Logf("TestGetMemberHelperListTotalNum return num:%v", num)
}

func TestGetMemberHelperList(t *testing.T) {
	mhlr := NewMemberHelperListRequest(2, "ha")
	res, err := GetMemberHelperList(mhlr)
	if err != nil {
		t.Fatalf("TestGetMemberHelperList query failed:%v", err)
	}
	t.Logf("TestGetMemberHelperList return value:%v", res)
}
