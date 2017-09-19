package emailattestation

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
	ceurs := NewCheckEmailUseRequestStruct(28663, "1234234@qq.com")
	res := CheckEmailUse(ceurs)
	if res != 1 {
		t.Fatalf("TestCheckEmailUse failed")
	}
}

func TestUserAttestationSave(t *testing.T) {
	uass := NewUserAttestationSaveStruct(2866343, 2)
	res := UserAttestationSave(uass)
	if res == 0 {
		t.Fatalf("TestUserAttestationSave failed")
	}
}

func TestUserEmailSave(t *testing.T) {
	uesrs := NewUserEmailSaveRequestStruct(28663, "1234234@qq.com")
	res := UserEmailSave(uesrs)
	if res != 1 {
		t.Fatalf("TestUserEmailSave failed")
	}
}

func TestInsertEmailLog(t *testing.T) {
	sers := NewSendEmailRequestStruct(2866345, "744596025@qq.com", "sb", `亲爱的会员:cgg,您好!<br/><br/>\r\n
		感谢您注册诚汇通—安全可靠的民间借贷公司-致力于打造国内口碑很好的P2P网络借贷平台.<br>\r\n
		 您在诚汇通可以享受到以下服务：<br/> 1)您可以通过诚汇通发布借款请求以解决资金紧缺需要.<br/>\r\n   
		  2)可以将自己的闲散资金出借给有需要的人获得资金回报".<br/><br/>\r\n   
		 <br/><a href=\http://cht.com/content/regedit?verifykey=CFCD208495D565EFUAJUVlwIGFUMBVABBlMNBQIaUghfVg1SCFMCWlF2VwFQHgcJCQ
		 并从中获得乐趣！<br/>\r\n   `,
		"182.34.34.44")
	// sers := NewSendEmailRequestStruct(2866345, "744596025@qq.com", "sb", `
	// 	诚汇通首页：<br/><br/>\r\n`,
	// 	"182.34.34.44")
	res, err := InsertEmailLog(sers)
	if err != nil {
		t.Fatalf("TestInsertEmailLog failed")
	}
	t.Logf("TestInsertEmailLog res:%v", res)
}

func TestUpdateEmailLog(t *testing.T) {
	lastID := int32(2487902)
	err := UpdateEmailLog(lastID)
	if err != nil {
		t.Fatalf("TestUpdateEmailLog failed")
	}
}

func TestSendSmtpMail(t *testing.T) {
	sers := NewSendEmailRequestStruct(28663, "744596025@qq.com", "感谢注册诚汇通，请完成您的邮箱验证!", `<h3>
        亲爱的会员,您好!
        </h3>`, "182.34.34.44")
	err := SendSmtpMail(sers)
	if err != nil {
		t.Fatalf("TestSendSmtpMail failed %v", err)
	}
}

func TestSendEmail(t *testing.T) {
	sers := NewSendEmailRequestStruct(28663, "744596025@qq.com", "感谢注册诚汇通，请完成您的邮箱验证!", `<h3>
        亲爱的会员,您好!
        </h3>`, "182.34.34.44")
	res := SendEmail(sers)
	if res != 1 {
		t.Fatalf("TestSendEmail failed")
	}
}
