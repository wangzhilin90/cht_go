namespace php SysUser.RoleAdd
namespace go roleadd

// +++ 添加角色
struct RoleAddRequestStruct {
    1:string role_name,//角色名称
    2:string remark,//角色说明，非必填
    3:string chengHuiTongTraceLog
}

struct RoleAddResponseStruct {
    1:i32 status,
    2:string msg
}

service RoleAddThriftService {
    RoleAddResponseStruct addRole (1:RoleAddRequestStruct requestObj)    //sql = "INSERT INTO jl_sys_role (role_name,remark) VALUES ('','')"
}