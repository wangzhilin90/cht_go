namespace go updateuserpasswword
namespace php User.UpdateUserPasswWord

struct UpdateUserPasswWordRequestStruct {
    1: i32 id
	2: string password
}

struct UpdateUserPasswWordResponseStruct {
     1:i32 status
	 2:string msg
}

service UpdateUserPasswWordThriftService {
    UpdateUserPasswWordResponseStruct updateUserPasswWord (1:UpdateUserPasswWordRequestStruct requestObj)
}