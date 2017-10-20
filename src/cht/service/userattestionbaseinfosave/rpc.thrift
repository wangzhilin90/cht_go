namespace php User.UserAttestionBaseInfoSave
namespace go userattestionbaseinfosave

struct UserAttestionBaseInfoSaveRequestStruct {
    1:i32 user_id,
    2:string video_pic,
    3:i32 real_status,
    4:i32 email_status,
    5:i32 phone_status,
    6:i32 video_status,
    7:i32 scene_status,
    8:i32 real_passtime,
    9:string chengHuiTongTraceLog
}

struct UserAttestionBaseInfoSaveResponseStruct {
    1:i32 status,
    2:string msg,
}

service UserAttestionBaseInfoSaveThriftService {
    UserAttestionBaseInfoSaveResponseStruct saveUserAttestionBaseInfo (1:UserAttestionBaseInfoSaveRequestStruct requestObj)
}

//sql:UPDATE jl_user_attestation SET .... WHERE user_id = $user_id
