package juanzengthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewRequestStruct(user_id int32, content string) *RequestStruct {
	return &RequestStruct{
		UserID:  user_id,
		Content: content,
	}
}

func TestGetInfo(t *testing.T) {
	rs := NewRequestStruct(214, "selecttestcontent'")
	js := juanzengservice{}
	res, err := js.GetInfo(rs)
	if err != nil {
		t.Fatalf("TestGetInfo return failed %v", err)
	}
	t.Logf("TestGetInfo res %v", res)
}

func TestAddMess(t *testing.T) {
	rs := NewRequestStruct(224, "testcontent")
	js := juanzengservice{}
	res, err := js.AddMess(rs)
	if err != nil {
		t.Fatalf("TestAddMess return failed %v", err)
	}
	t.Logf("TestAddMess res %v", res)
}
