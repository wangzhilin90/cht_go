namespace php User.UserAttestionCardInfoSave
namespace go  userattestioncardinfosave

struct UserAttestionCardInfoSaveRequestStruct {
    1:i32 user_id,
    2:i32 card_type,
    3:string card_id,
    4:string chengHuiTongTraceLog
}

struct UserAttestionCardInfoSaveResponseStruct {
    1:i32 status,
    2:string msg,
}

service UserAttestionCardInfoSaveThriftService {
    UserAttestionCardInfoSaveResponseStruct saveUserAttestionCardInfo (1:UserAttestionCardInfoSaveRequestStruct requestObj)
}

//sql:UPDATE jl_user_attestation SET .... WHERE user_id = $user_id
