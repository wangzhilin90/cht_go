namespace php SysUser.RoleEdit
namespace go roleedit

// +++ 角色编辑

struct RoleEditRequestStruct {
    1:i32 role_id ,//角色id
    2:string role_name,//角色名称
    3:string remark,//角色说明
    4:string chengHuiTongTraceLog
}

struct RoleEditResponseStruct {
    1:i32 status,
    2:string msg
}

service RoleEditThriftService {
    RoleEditResponseStruct editRole (1:RoleEditRequestStruct requestObj) //sql = "UPDATE jl_sys_role SET role_name = $role_name,remark = $remark .... WHERE id = $role_id";
}