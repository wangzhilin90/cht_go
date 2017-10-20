package userattestionbaseinfosave

import (
	_ "cht/initial"
	"testing"
)

func NewUserAttestionBaseInfoSaveRequest(user_id int32, vidio_pid string, sceneStatus int32) *UserAttestionBaseInfoSaveRequest {
	return &UserAttestionBaseInfoSaveRequest{
		UserID:      user_id,
		VideoPic:    vidio_pid,
		SceneStatus: sceneStatus,
	}
}

func TestSaveUserAttestionBaseInfo(t *testing.T) {
	uabisr := NewUserAttestionBaseInfoSaveRequest(1, "vidio_paht", 1)
	b := SaveUserAttestionBaseInfo(uabisr)
	if b == false {
		t.Fatalf("TestSaveUserAttestionBaseInfo failed")
	}
}
