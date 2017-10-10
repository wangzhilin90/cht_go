namespace php SysUser.RoleRightSet
namespace go rolerightset

// +++ 角色权限编辑

struct RoleRightSetRequestStruct {
    1:i32 role_id ,//角色id
    2:string power_config,//角色权限，角色id之间通过逗号(,)串联成的字符串
    5:string chengHuiTongTraceLog
}

struct RoleRightSetResponseStruct {
    1:i32 status,
    2:string msg
}

service RoleRightSetThriftService {
    RoleRightSetResponseStruct setRoleRight (1:RoleRightSetRequestStruct requestObj) //sql = "UPDATE jl_sys_role SET power_config = $power_config WHERE id = $role_id";
}