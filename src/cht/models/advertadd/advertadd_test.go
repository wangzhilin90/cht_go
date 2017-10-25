package advertadd

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertAddRequest(adType int32, title string) *AdvertAddRequest {
	return &AdvertAddRequest{
		Type:  adType,
		Img:   " ",
		Title: title,
	}
}

func TestAddAdvert(t *testing.T) {
	aar := NewAdvertAddRequest(20, "周年庆1")
	b := AddAdvert(aar)
	if b == false {
		t.Fatalf("TestAddAdvert failed")
	}
}
