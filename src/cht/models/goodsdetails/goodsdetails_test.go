package goodsdetails

import (
	_ "cht/initial"
	"testing"
)

func NewGoodsDetailsRequest(id int32) *GoodsDetailsRequest {
	return &GoodsDetailsRequest{
		ID: id,
	}
}

func TestGetGoodsDetails(t *testing.T) {
	gdr := NewGoodsDetailsRequest(12)
	res, err := GetGoodsDetails(gdr)
	if err != nil {
		t.Fatalf("TestGetGoodsDetails query failed:%v", err)
	}

	t.Logf("TestGetGoodsDetails return value:%v", res)
}
