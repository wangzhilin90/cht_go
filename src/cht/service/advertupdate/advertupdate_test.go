package advertupdate

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertUpdateRequestStruct(id int32, title string) *AdvertUpdateRequestStruct {
	return &AdvertUpdateRequestStruct{
		Title: title,
		ID:    id,
	}
}

func TestUpdateAdvert(t *testing.T) {
	aur := NewAdvertUpdateRequestStruct(124, "周年庆fasdf")
	aus := advertupdateservice{}
	res, _ := aus.UpdateAdvert(aur)
	if res.Status != UPDATE_ADVERT_SUCCESS {
		t.Fatalf("TestUpdateAdvert failed")
	}
}
