package vipmemberranklist

import (
	_ "cht/initial"
	"testing"
)

func NewVipMemberRankListRequest(uType int32, keywords string, limitnum, limitoffset int32) *VipMemberRankListRequest {
	return &VipMemberRankListRequest{
		Type:        uType,
		Keywords:    keywords,
		LimitNum:    limitnum,
		LimitOffset: limitoffset,
	}
}

func TestGetVipMemberRankListTotalNum(t *testing.T) {
	vmrlr := NewVipMemberRankListRequest(1, "", 20, 0)
	num, err := GetVipMemberRankListTotalNum(vmrlr)
	if err != nil {
		t.Fatalf("TestGetVipMemberRankListTotalNum failed:%v", err)
	}
	t.Logf("TestGetVipMemberRankListTotalNum return num:%v", num)
}

func TestGetVipMemberRankList(t *testing.T) {
	vmrlr := NewVipMemberRankListRequest(1, "", 20, 0)
	num, err := GetVipMemberRankList(vmrlr)
	if err != nil {
		t.Fatalf("TestGetVipMemberRankList failed:%v", err)
	}
	t.Logf("TestGetVipMemberRankList return num:%v", num)
}
