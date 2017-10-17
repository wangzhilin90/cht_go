namespace php Common.SysUserDelete
namespace go  sysuserdelete

//删除后台管理用户
struct SysUserDeleteRequestStruct {
    1:string user_id_str,
    2:string  chengHuiTongTraceLog
}

struct SysUserDeleteResponseStruct {
	1:i32 status,
	2:string msg
}

service SysUserDeleteThriftService {
    SysUserDeleteResponseStruct deleteSysUser (1:SysUserDeleteRequestStruct requestObj)
}

//sql :注释 where条件为空时禁止删除（若做啦删除操作相当清空数据表）DELETE FROM  jl_sys_user WHERE id  IN ($user_id_str);


