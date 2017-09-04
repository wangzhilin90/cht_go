namespace go  updatepasswdthriftservice

struct UpdatePasswdRequestStruct {
    1: i32 id 
	2: string  password
}

struct UpdatePasswdResponseStruct {
     1:i32 status
	 2:string msg
}

service UpdatePasswdThriftService {
    UpdatePasswdResponseStruct updatePasswd(1:UpdatePasswdRequestStruct requestObj)
}


