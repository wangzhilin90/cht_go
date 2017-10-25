package advertdetails

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertDetailsRequestStruct(id int32) *AdvertDetailsRequestStruct {
	return &AdvertDetailsRequestStruct{
		ID: id,
	}
}

func TestGetAdvertDetails(t *testing.T) {
	adr := NewAdvertDetailsRequestStruct(18)
	ads := advertdetailsservice{}
	res, _ := ads.GetAdvertDetails(adr)
	if res.Status != GET_ADVERT_DETAILS_SUCCESS {
		t.Fatalf("TestGetAdvertDetails failed")
	}
	t.Logf("TestGetAdvertDetails return value:%v", res)
}
