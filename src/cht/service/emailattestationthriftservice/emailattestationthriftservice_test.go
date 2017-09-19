package emailattestationthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewCheckEmailUseRequestStruct(user_id int32, email string) *CheckEmailUseRequestStruct {
	return &CheckEmailUseRequestStruct{
		Email:  email,
		UserID: user_id,
	}
}

func NewUserAttestationSaveStruct(user_id int32, emailStatus int32) *UserAttestationSaveStruct {
	return &UserAttestationSaveStruct{
		UserID:      user_id,
		EmailStatus: emailStatus,
	}
}

func NewUserEmailSaveRequestStruct(user_id int32, email string) *UserEmailSaveRequestStruct {
	return &UserEmailSaveRequestStruct{
		UserID: user_id,
		Email:  email,
	}
}

func NewSendEmailRequestStruct(user_id int32, send_to string, subject string, content string, ip string) *SendEmailRequestStruct {
	return &SendEmailRequestStruct{
		UserID:  user_id,
		SendTo:  send_to,
		Subject: subject,
		Content: content,
		IP:      ip,
	}
}

func TestCheckEmailUse(t *testing.T) {
	ceurs := NewCheckEmailUseRequestStruct(286623, "1234234@qq.com")
	eas := emailattestationservice{}
	res, _ := eas.CheckEmailUse(ceurs)
	if res != 1 {
		t.Fatalf("TestCheckEmailUse failed")
	}
	t.Logf("TestCheckEmailUse res:%v", res)
}

func TestUserAttestationSave(t *testing.T) {
	uass := NewUserAttestationSaveStruct(2866344, 2)
	eas := emailattestationservice{}
	res, _ := eas.UserAttestationSave(uass)
	if res == 0 {
		t.Fatalf("TestUserAttestationSave failed")
	}
	t.Logf("TestUserAttestationSave res:%v", res)
}

func TestUserEmailSave(t *testing.T) {
	uesrs := NewUserEmailSaveRequestStruct(28663, "12342345@qq.com")
	eas := emailattestationservice{}
	res, _ := eas.UserEmailSave(uesrs)
	if res != 1 {
		t.Fatalf("TestUserEmailSave failed")
	}
	t.Logf("TestUserEmailSave res:%v", res)
}

func TestSendEmail(t *testing.T) {
	sers := NewSendEmailRequestStruct(28663, "744596025@qq.com", "感谢注册诚汇通，请完成您的邮箱验证!", `<h3>
        亲爱的会员,您好!
        </h3>`, "182.34.34.45")
	eas := emailattestationservice{}
	res, _ := eas.SendEmail(sers)
	if res != 1 {
		t.Fatalf("TestSendEmail failed")
	}
	t.Logf("TestSendEmail res:%v", res)
}
