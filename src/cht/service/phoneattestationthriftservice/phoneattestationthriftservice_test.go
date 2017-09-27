package phoneattestationthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewCheckPhoneUseRequestStruct(phone string) *CheckPhoneUseRequestStruct {
	return &CheckPhoneUseRequestStruct{
		Phone: phone,
	}
}

func NewGetUserIdByhsidRequestStruct(hsid string) *GetUserIdByhsidRequestStruct {
	return &GetUserIdByhsidRequestStruct{
		Hsid: hsid,
	}
}

func NewUpdatePhoneRequestStruct(user_id int32, phone string) *UpdatePhoneRequestStruct {
	return &UpdatePhoneRequestStruct{
		Phone:  phone,
		UserID: user_id,
	}
}

func TestCheckPhoneByPhone(t *testing.T) {
	cprs := NewCheckPhoneUseRequestStruct("13537573273")
	pts := &phoneattestationservice{}
	str, _ := pts.CheckPhoneByPhone(cprs)
	if str == "1000" {
		t.Fatalf("TestCheckPhoneUse failed")
	}
}

func TestGetUserIdByhsid(t *testing.T) {
	uprs := NewGetUserIdByhsidRequestStruct("9930040050160000010")
	pts := &phoneattestationservice{}
	user_id, _ := pts.GetUserIdByhsid(uprs)
	if user_id != 313576 {
		t.Fatalf("TestGetUserIdByhsid failed")
	}
}

func TestUpdatePhoneByTransaction(t *testing.T) {
	ubrs := NewUpdatePhoneRequestStruct(204742, "1371405792")
	pts := &phoneattestationservice{}
	str, _ := pts.UpdatePhoneByTransaction(ubrs)
	if str == "1001" {
		t.Fatalf("TestUpdatePhone failed")
	}
}
