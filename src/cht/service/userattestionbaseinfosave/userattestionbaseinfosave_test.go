package userattestionbaseinfosave

import (
	_ "cht/initial"
	"testing"
)

func NewUserAttestionBaseInfoSaveRequestStruct(user_id int32, vidio_pid string, sceneStatus int32) *UserAttestionBaseInfoSaveRequestStruct {
	return &UserAttestionBaseInfoSaveRequestStruct{
		UserID:      user_id,
		VideoPic:    vidio_pid,
		SceneStatus: sceneStatus,
	}
}

func TestSaveUserAttestionBaseInfo(t *testing.T) {
	uabisr := NewUserAttestionBaseInfoSaveRequestStruct(1, "vidio_paht", 1)
	uars := userattestionbaseinfosaveservice{}
	res, _ := uars.SaveUserAttestionBaseInfo(uabisr)
	if res.Status == SAVE_USER_ATTESTION_BASE_INFO_FAILED {
		t.Fatalf("TestSaveUserAttestionBaseInfo failed")
	}
}
