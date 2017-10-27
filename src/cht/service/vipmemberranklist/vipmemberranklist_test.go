package vipmemberranklist

import (
	_ "cht/initial"
	"testing"
)

func NewVipMemberRankListRequestStruct(uType int32, keywords string, limitnum, limitoffset int32) *VipMemberRankListRequestStruct {
	return &VipMemberRankListRequestStruct{
		Type:        uType,
		Keywords:    keywords,
		LimitNum:    limitnum,
		LimitOffset: limitoffset,
	}
}

func TestGetVipMemberRankList(t *testing.T) {
	vmrlr := NewVipMemberRankListRequestStruct(1, "", 20, 0)
	vmrls := vipmemberranklistservice{}
	res, _ := vmrls.GetVipMemberRankList(vmrlr)
	if res.Status != QUERY_VIPMEMBERRANKLIST_SUCCESS {
		t.Fatalf("TestGetVipMemberRankList failed")
	}
	t.Logf("TestGetVipMemberRankList response:%v", res)
}
