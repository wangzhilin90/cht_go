namespace php SysUser.SysUserAdd
namespace go  sysuseradd

//添加后台管理用户

struct SysUserAddRequestStruct {
    1:string account,
    2:string password,//加密后的密码
    3:string realname,
    4:string mobile,
    5:string qq,
    6:i32 status,
    7:i32 role_id,
    8:i32  ,//角色类型
    9:i32 create_time,//添加时间
    10:i32 lastlogintime,//最后登录时间
	11:i32 views,//客服分配的次数
    12:string lastloginip,//最后登录ip
    13:string chengHuiTongTraceLog
}

struct SysUserAddResponseStruct {
	1:i32 status,
	3:string msg,
}

service SysUserAddThriftService {
    SysUserAddResponseStruct addSysUser (1:SysUserAddRequestStruct requestObj)
}
//sql: insert into jl_sys_user (id,account,password,realname,mobile,qq,status,role_id,customer_type,create_time,lastlogintime,views,lastloginip) values(next VALUE FOR MYCATSEQ_SYS_USER,?.....)