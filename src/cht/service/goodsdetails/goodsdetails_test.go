package goodsdetails

import (
	_ "cht/initial"
	"testing"
)

func NewGoodsDetailsRequestStruct(id int32) *GoodsDetailsRequestStruct {
	return &GoodsDetailsRequestStruct{
		ID: id,
	}
}

func TestGetGoodsDetails(t *testing.T) {
	gdr := NewGoodsDetailsRequestStruct(12)
	gds := goodsdetailsservice{}
	res, _ := gds.GetGoodsDetails(gdr)
	if res.Status != QUERY_GOODS_DETAILS_SUCCESS {
		t.Fatalf("TestGetGoodsDetails query failed")
	}
	t.Logf("TestGetGoodsDetails response:%v", res)
}
