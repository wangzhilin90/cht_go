package juanzeng

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

func TestGetMesslistResult(t *testing.T) {
	rs := NewRequestStruct(214, "testcontent")
	res, err := GetMesslistResult(rs)
	if err != nil {
		t.Fatalf("TestGetMesslistResult failed %v", err)
	}
	t.Logf("TestGetMesslistResult res :%v", res)
}

func TestGetFundlistResult(t *testing.T) {
	rs := NewRequestStruct(214, "testcontent")
	res, err := GetFundlistResult(rs)
	if err != nil {
		t.Fatalf("TestGetFundlistResult failed %v", err)
	}
	t.Logf("TestGetFundlistResult res :%v", res)
}

func TestGetNumlistResult(t *testing.T) {
	rs := NewRequestStruct(214, "testcontent")
	res, err := GetNumlistResult(rs)
	if err != nil {
		t.Fatalf("TestGetNumlistResult failed %v", err)
	}
	t.Logf("TestGetNumlistResult res :%v", res)
}

func TestGetTotalJuanNum(t *testing.T) {
	rs := NewRequestStruct(214, "testcontent")
	res, err := GetTotalJuanNum(rs)
	if err != nil {
		t.Fatalf("TestGetTotalJuanNum failed %v", err)
	}
	t.Logf("TestGetTotalJuanNum res :%v", res)
}

func TestAddMess(t *testing.T) {
	rs := NewRequestStruct(214, "testcontent")
	res, err := AddMess(rs)
	if err != nil {
		t.Fatalf("TestAddMess failed %v", err)
	}
	t.Logf("TestAddMess res :%v", res)
}
