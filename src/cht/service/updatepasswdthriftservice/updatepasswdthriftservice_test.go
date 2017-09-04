package updatepasswdthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewUpdatePasswdRequestStruct(id int32, passwd string) *UpdatePasswdRequestStruct {
	return &UpdatePasswdRequestStruct{
		ID:       id,
		Password: passwd,
	}
}

func TestUpdatePasswd(t *testing.T) {
	uprs := NewUpdatePasswdRequestStruct(134234, "mimamamaipi")
	ups := updatepasswdservice{}
	res, err := ups.UpdatePasswd(uprs)
	if err != nil {
		t.Fatal("TestUpdatePasswd failed")
	}
	t.Log("TestUpdatePasswd res:", res)
}
