namespace php SysUser.RoleDelete
namespace go roledelete

// +++ 角色删除
struct RoleDeleteRequestStruct {
    1:string role_id_str ,//角色id
    2:string chengHuiTongTraceLog
}

struct RoleDeleteResponseStruct {
    1:i32 status,
    2:string msg
}

service RoleDeleteThriftService {
    RoleDeleteResponseStruct deleteRole (1:RoleDeleteRequestStruct requestObj) //sql = "DELETE FROM jl_sys_role WHERE id !=1 AND id IN (' . $role_id_str . ')'"; //接收的id 为 逗号隔开的字符串, id= 1 标识为管理员
}