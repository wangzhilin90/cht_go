namespace php SysUser.KeFuList
namespace go kefulist

struct KeFuListRequestStruct {
    1:i32 role_id,//角色id
    2:i32 status,//启用状态 默认 0 ，启用 1 停用',
    3:string customer_type,//客服类型：对应字典49类值；（原值：0无，1咨新客服，2指导客服，3专属客服）多个值用,号隔开，例如： 2,3
    4:string chengHuiTongTraceLog
}

struct KeFuDetailsStruct {
    1:i32 id //smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '后台系统用户表',
    2:i32 role_id //smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '角色表中id',
    3:string account //varchar(30) NOT NULL COMMENT '登录账号',
    4:string realname //varchar(20) NOT NULL DEFAULT '' COMMENT '真实姓名',
    5:string password //char(32) NOT NULL COMMENT '登录密码',
    6:string mobile //varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
    7:string qq //varchar(11) NOT NULL DEFAULT '' COMMENT 'qq',
    8:string lastloginip //char(15) NOT NULL COMMENT '最后登录IP',
    9:i32 lastlogintime //int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后登录时间',
    10:i32 create_time //int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    11:i32 status //tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否启用 默认 0 启用 1 停用',
    12:i32 views //int(10) unsigned NOT NULL DEFAULT '0' COMMENT '客服分配的次数',
    13:i32 customer_type //smallint(5) unsigned NOT NULL DEFAULT '668' COMMENT '客服类型：对应字典49类值；（原值：0无，1咨新客服，2指导客服，3专属客服）',
}

struct KeFuListResponseStruct {
    1:i32 status,
    2:list<KeFuDetailsStruct> KeFuList,
    3:string msg
}
service KeFuListThriftService {
    //sql = "select * from jl_sys_user where status = 0 and role_id = 2 and customer_type IN (2,3)"
    KeFuListResponseStruct getKeFuList (1:KeFuListRequestStruct requestObj) 
}