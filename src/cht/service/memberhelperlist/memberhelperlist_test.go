package memberhelperlist

import (
	_ "cht/initial"
	"testing"
)

func NewMemberHelperListRequestStruct(utype int32, keywords string) *MemberHelperListRequestStruct {
	return &MemberHelperListRequestStruct{
		Type:     utype,
		Keywords: keywords,
	}
}

func TestGetMemberHelperList(t *testing.T) {
	mhlrs := NewMemberHelperListRequestStruct(2, "ha")
	mhls := memberhelperlistservice{}
	res, _ := mhls.GetMemberHelperList(mhlrs)
	if res.Status != QUERY_MEMBER_HELP_LIST_SUCCESS {
		t.Fatalf("TestGetMemberHelperList query failed")
	}
	t.Logf("TestGetMemberHelperList return value:%v", res)
}
