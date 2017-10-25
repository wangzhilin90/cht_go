package advertdetails

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertDetailsRequest(id int32) *AdvertDetailsRequest {
	return &AdvertDetailsRequest{
		ID: id,
	}
}

func TestGetAdvertDetails(t *testing.T) {
	adr := NewAdvertDetailsRequest(18)
	res, err := GetAdvertDetails(adr)
	if err != nil {
		t.Fatalf("TestGetAdvertDetails failed:%v", err)
	}
	t.Logf("TestGetAdvertDetails return value:%v", res)
}
