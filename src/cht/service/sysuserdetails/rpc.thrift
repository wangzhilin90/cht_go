namespace php SysUser.SysUserDetails
namespace go sysuserdetails

//后台管理员详情

struct SysUserDetailsRequestStruct {
    1:i32 user_id,
    2:string chengHuiTongTraceLog
}

struct SysUserDetailsStruct {
  1:i32 id,//后台系统用户表
  2:i32 role_id,//角色表中id
  3:string account,//登录账号
  4:string realname,//真实姓名
  5:string password,//登录密码
  6:string mobile,//手机号
  7:string qq,//qq
  8:string lastloginip,//最后登录IP
  9:i32 lastlogintime,//最后登录时间
  10:i32 create_time,//创建时间
  11:i32 status,//是否启用 默认 0 启用 1 停用
  12:i32 views,//客服分配的次数
  13:i32 customer_type //客服类型：0无，1咨新客服，2指导客服，3专属客服
}

struct SysUserDetailsResponseStruct {
	1:i32 status,
    2:SysUserDetailsStruct SysUserDetails,
	3:string msg,
}

service SysUserDetailsThriftService {
    SysUserDetailsResponseStruct getSysUserDetails (1:SysUserDetailsRequestStruct requestObj)
}