namespace php SysUser.SysUserEdit
namespace go sysuseredit

//编辑后台管理用户
struct SysUserEditRequestStruct {
    1:string account,
    2:string password,//加密后的密码,如果 password为空 标识为未修改密码不用更新该字段,否则需要修改密码
    3:string realname,
    4:string mobile,
    5:string qq,
    6:i32 status,
    7:i32 role_id,
    8:i32 customer_type,//角色类型
    9:string create_time,//更新修改时间
    10:i32 user_id,//后台用户id
    11:string chengHuiTongTraceLog
}

struct SysUserEditResponseStruct {
	1:i32 status,
	3:string msg,
}

service SysUserEditThriftService {
    SysUserEditResponseStruct editSysUser(1:SysUserEditRequestStruct requestObj)
}