package advertlist

import (
	_ "cht/initial"
	"testing"
)

func NewAdvertListRequest() *AdvertListRequest {
	return &AdvertListRequest{}
}

func TestGetAdvertListTatalNum(t *testing.T) {
	alr := NewAdvertListRequest()
	num, err := GetAdvertListTatalNum(alr)
	if err != nil {
		t.Fatalf("TestGetAdvertListTatalNum failed:%v", err)
	}
	t.Logf("TestGetAdvertListTatalNum return num:%v", num)
}

func TestGetAdvertList(t *testing.T) {
	alr := NewAdvertListRequest()
	res, err := GetAdvertList(alr)
	if err != nil {
		t.Fatalf("TestGetAdvertList failed:%v", err)
	}
	t.Logf("TestGetAdvertList res:%v", res)
}
