package goodsdel

import (
	_ "cht/initial"
	"testing"
)

func NewGoodsDeLRequest(id int32) *GoodsDeLRequest {
	return &GoodsDeLRequest{
		ID: id,
	}
}

func TestDelGoods(t *testing.T) {
	gdr := NewGoodsDeLRequest(118)
	b := DelGoods(gdr)
	if b == false {
		t.Fatalf("TestDelGoods failed")
	}
}
