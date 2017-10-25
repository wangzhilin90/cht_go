package advertdel

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertDelRequestStruct(id int32) *AdvertDelRequestStruct {
	return &AdvertDelRequestStruct{
		ID: id,
	}
}

func TestDelAdvert(t *testing.T) {
	adr := NewAdvertDelRequestStruct(123)
	ads := advertdelservice{}
	res, _ := ads.DelAdvert(adr)
	if res.Status != DELETE_ADVERT_SUCCESS {
		t.Fatalf("TestDelAdvert failed")
	}
}
