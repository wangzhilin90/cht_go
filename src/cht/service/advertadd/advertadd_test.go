package advertadd

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertAddRequestStruct(adType int32, title string) *AdvertAddRequestStruct {
	return &AdvertAddRequestStruct{
		Type:  adType,
		Img:   " ",
		Title: title,
	}
}

func TestAddAdvert(t *testing.T) {
	aar := NewAdvertAddRequestStruct(1, "biaoti")
	aas := advertaddservice{}
	res, _ := aas.AddAdvert(aar)
	if res.Status != ADD_ADVERT_SUCCESS {
		t.Fatalf("TestAddAdvert failed")
	}
}
