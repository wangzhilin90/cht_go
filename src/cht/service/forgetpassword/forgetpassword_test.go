package forgetpassword

import (
	_ "cht/initial"
	"testing"
)

func NewForgetPasswordRequestStruct(id int32, password string) *ForgetPasswordRequestStruct {
	return &ForgetPasswordRequestStruct{
		ID:       id,
		Password: password,
	}
}

func TestForgetPassword(t *testing.T) {
	fprs := NewForgetPasswordRequestStruct(134234, "newmamaipi")
	fps := forgetpasswordservice{}
	res, _ := fps.ForgetPassword(fprs)
	if res.Status != UPDATE_FORGET_PASSWORD_SUCCESS {
		t.Fatalf("TestForgetPassword failed")
	}
}
