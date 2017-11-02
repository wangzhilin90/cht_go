package accessconfig

import (
	_ "cht/initial"
	"testing"
)

func NewAccessConfigRequest(source string) *AccessConfigRequest {
	return &AccessConfigRequest{
		Source: source,
	}
}

func TestGetAccessConfig(t *testing.T) {
	acr := NewAccessConfigRequest("ac34752a516542b2694ad24b3c3f70b7")
	res, err := GetAccessConfig(acr)
	if err != nil {
		t.Fatalf("TestGetAccessConfig failed:%v", err)
	}
	t.Logf("TestGetAccessConfig return value:%v", res)
}
