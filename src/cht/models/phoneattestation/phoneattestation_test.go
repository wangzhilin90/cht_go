package phoneattestation

import (
	_ "cht/initial"
	"testing"
)

func NewCheckPhoneUseRequest(phone string) *CheckPhoneUseRequest {
	return &CheckPhoneUseRequest{
		Phone: phone,
	}
}

func NewGetUserIdByhsidRequest(hsid string) *GetUserIdByhsidRequest {
	return &GetUserIdByhsidRequest{
		Hsid: hsid,
	}
}

func NewUpdatePhoneRequest(user_id int32, phone string) *UpdatePhoneRequest {
	return &UpdatePhoneRequest{
		Phone:  phone,
		UserID: user_id,
	}
}

func TestCheckPhoneUse(t *testing.T) {
	cpr := NewCheckPhoneUseRequest("13537573273")
	b := CheckPhoneUse(cpr)
	if b == false {
		t.Fatalf("TestCheckPhoneUse failed")
	}
}

func TestGetUserIdByhsid(t *testing.T) {
	gihr := NewGetUserIdByhsidRequest("9930040050160000010")
	userId, err := GetUserIdByhsid(gihr)
	if err != nil && userId != 313576 {
		t.Fatalf("TestGetUserIdByhsid failed %v", err)
	}
}

func TestUpdatePhone(t *testing.T) {
	upr := NewUpdatePhoneRequest(204742, "1371405795")
	b := UpdatePhone(upr)
	if b != "1000" {
		t.Fatalf("TestUpdatePhone failed")
	}
}
