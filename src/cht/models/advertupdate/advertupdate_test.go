package advertupdate

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertUpdateRequest(id int32, title string) *AdvertUpdateRequest {
	return &AdvertUpdateRequest{
		Title: title,
		ID:    id,
	}
}

func TestUpdateAdvert(t *testing.T) {
	aur := NewAdvertUpdateRequest(124, "周年庆jjja")
	b := UpdateAdvert(aur)
	if b == false {
		t.Fatalf("TestUpdateAdvert failed")
	}
}
