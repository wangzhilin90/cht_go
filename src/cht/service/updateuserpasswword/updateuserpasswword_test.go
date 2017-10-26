package updateuserpasswword

import (
	_ "cht/initial"
	"testing"
)

func NewUpdatePasswdRequestStruct(id int32, oldpw, newpw string) *UpdateUserPasswWordRequestStruct {
	return &UpdateUserPasswWordRequestStruct{
		ID:           id,
		OldPassword:  oldpw,
		NewPassword_: newpw,
	}
}

func TestUpdatePasswd(t *testing.T) {
	uprs := NewUpdatePasswdRequestStruct(134234, "mimamamaipi", "newmimamamaipi")
	ups := updatepasswdservice{}
	res, err := ups.UpdateUserPasswWord(uprs)
	if err != nil {
		t.Fatal("TestUpdatePasswd failed")
	}
	t.Log("TestUpdatePasswd res:", res)
}
