namespace php SysUser.RoleDetails
namespace go roledetails

// +++ 角色详情
struct RoleDetailsRequestStruct {
    1:i32 role_id ,//角色id
    2:string chengHuiTongTraceLog
}

struct RoleDetailsStruct {
    1:i32 id,
    2:string role_name,
    3:string remark,
    4:string power_config,
    5:i32 create_time,
}

struct RoleDetailsResponseStruct {
    1:i32 status,
    2:RoleDetailsStruct RoleDetails
    3:string msg
}

service RoleDetailsThriftService {
    RoleDetailsResponseStruct getRoleDetails (1:RoleDetailsRequestStruct requestObj) //sql = "SELECT * FROM jl_sys_role WHERE id = $role_id LIMIT 1";
}