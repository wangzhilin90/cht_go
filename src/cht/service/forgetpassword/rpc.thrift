//忘记密码重置密码服务
namespace php User.ForgetPassword
namespace go  forgetpassword

struct ForgetPasswordRequestStruct {
    1: i32 id
	2: string password
}

struct ForgetPasswordResponseStruct {
     1:i32 status
	 2:string msg
}

service ForgetPasswordThriftService {
    ForgetPasswordResponseStruct forgetPassword (1:ForgetPasswordRequestStruct requestObj)
}