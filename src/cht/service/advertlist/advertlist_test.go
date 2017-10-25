package advertlist

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertListRequestStruct() *AdvertListRequestStruct {
	return &AdvertListRequestStruct{}
}

func TestGetAdvertList(t *testing.T) {
	alrs := NewAdvertListRequestStruct()
	als := advertlistservice{}
	res, _ := als.GetAdvertList(alrs)
	if res.Status != QUERY_ADVERT_LIST_SUCCESS {
		t.Fatalf("TestGetAdvertList failed")
	}
	t.Logf("TestGetAdvertList response:%v", res)
}
