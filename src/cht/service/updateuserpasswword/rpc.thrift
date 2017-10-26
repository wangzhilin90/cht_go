//忘记密码重置密码服务
namespace go updateuserpasswword
namespace php User.UpdateUserPasswWord

struct UpdateUserPasswWordRequestStruct {
    1: i32 id,
	2: string newPassword,
	3: string oldPassword
}

struct UpdateUserPasswWordResponseStruct {
     1:i32 status,
	 2:string msg
}

service UpdateUserPasswWordThriftService {
    UpdateUserPasswWordResponseStruct updateUserPasswWord (1:UpdateUserPasswWordRequestStruct requestObj)
}

//1000:   "更新密码成功",
// 1001:   "获取数据库登录密码失败",
// 1002: "旧密码输入不正确",
// 1003:    "更新密码失败",

//SQL update jl_user set password=? where id=?


//!D('User')->getUsercount(array('id' => $this->user['id'], 'password' => D('User')->pwdhash($data['oldpass']))) && $this-jump_alert('旧密码输入不正确！');
//$result_user = D('User')->userSave(array('password' => $data['password']), array('id' => $this->user['id']));

