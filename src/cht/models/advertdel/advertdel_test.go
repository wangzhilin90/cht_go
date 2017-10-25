package advertdel

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertDelRequest(id int32) *AdvertDelRequest {
	return &AdvertDelRequest{
		ID: id,
	}
}

func TestGetAdvertFid(t *testing.T) {
	adr := NewAdvertDelRequest(1)
	res, err := GetAdvertFid(adr)
	if err != nil {
		t.Fatalf("TestGetAdvertFid failed")
	}
	t.Logf("TestGetAdvertFid return value:%v", res)
}

func TestDelAdvert(t *testing.T) {
	adr := NewAdvertDelRequest(12)
	b := DelAdvert(adr)
	if b == false {
		t.Fatalf("TestDelAdvert failed")
	}
}
