package goodslist

import (
	_ "cht/initial"
	"testing"
)

func NewGoodsListRequest(name string, category, limitnum, limitOffset int32) *GoodsListRequest {
	return &GoodsListRequest{
		Name:        name,
		Category:    category,
		LimitNum:    limitnum,
		LimitOffset: limitOffset,
	}
}

func TestGetGoodsListTotalNum(t *testing.T) {
	glr := NewGoodsListRequest("", 1, 10, 2)
	num, err := GetGoodsListTotalNum(glr)
	if err != nil {
		t.Fatalf("TestGetGoodsListTotalNum query failed:%v", err)
	}
	t.Logf("TestGetGoodsListTotalNum return num:%v", num)
}

func TestGetGoodsList(t *testing.T) {
	glr := NewGoodsListRequest("", 1, 10, 2)
	res, err := GetGoodsList(glr)
	if err != nil {
		t.Fatalf("TestGetGoodsList failed:%v", err)
	}
	t.Logf("TestGetGoodsList return value:%v", res)
}
