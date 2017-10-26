package forgetpassword

import (
	_ "cht/initial"
	"testing"
)

func NewForgetPasswordRequest(id int32, passwd string) *ForgetPasswordRequest {
	return &ForgetPasswordRequest{
		ID:       id,
		Password: passwd,
	}
}

func TestForgetPassword(t *testing.T) {
	fpr := NewForgetPasswordRequest(134234, "mamaipi")
	b := ForgetPassword(fpr)
	if b == false {
		t.Fatalf("TestForgetPassword failed")
	}
}
