package accessconfig

import (
	_ "cht/initial"
	"testing"
)

func NewAccessConfigRequestStruct(source string) *AccessConfigRequestStruct {
	return &AccessConfigRequestStruct{
		Source: source,
	}
}

func TestGetAccessConfig(t *testing.T) {
	acrs := NewAccessConfigRequestStruct("ac34752a516542b2694ad24b3c3f70b7")
	acs := accessconfigservice{}
	res, _ := acs.GetAccessConfig(acrs)
	if res.Status != QUERY_ACCESS_CONFIG_SUCCESS {
		t.Fatalf("TestGetAccessConfig failed")
	}
	t.Logf("TestGetAccessConfig response:%v", res)
}
