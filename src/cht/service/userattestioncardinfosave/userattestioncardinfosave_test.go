package userattestioncardinfosave

import (
	_ "cht/initial"
	"testing"
)

func NewUserAttestionCardInfoSaveRequestStruct(user_id, card_type int32, card_id string) *UserAttestionCardInfoSaveRequestStruct {
	return &UserAttestionCardInfoSaveRequestStruct{
		UserID:   user_id,
		CardID:   card_id,
		CardType: card_type,
	}
}

func TestSaveUserAttestionCardInfo(t *testing.T) {
	uacisr := NewUserAttestionCardInfoSaveRequestStruct(2, 1, "362204198512184332")
	uis := userattestioncardinfosaveservice{}
	res, _ := uis.SaveUserAttestionCardInfo(uacisr)
	if res.Status == SAVE_USER_ATTESTION_CARDINFO_FAILED {
		t.Fatalf("TestSaveUserAttestionCardInfo failed")
	}
}
