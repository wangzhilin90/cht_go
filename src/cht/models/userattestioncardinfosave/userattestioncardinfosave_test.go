package userattestionbaseinfosave

import (
	_ "cht/initial"
	"testing"
)

func NewUserAttestionCardInfoSaveRequest(user_id, card_type int32, card_id string) *UserAttestionCardInfoSaveRequest {
	return &UserAttestionCardInfoSaveRequest{
		UserID:   user_id,
		CardID:   card_id,
		CardType: card_type,
	}
}

func TestSaveUserAttestionCardInfo(t *testing.T) {
	uacisr := NewUserAttestionCardInfoSaveRequest(2, 0, "362204198512184331")
	b := SaveUserAttestionCardInfo(uacisr)
	if b == false {
		t.Fatalf("TestSaveUserAttestionCardInfo failed")
	}
}
