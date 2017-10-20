namespace php User.UserAttestionList
namespace go userattestionlist

struct UserAttestionListRequestStruct {
    1:string username,
    2:string realname,
    3:i32 real_status,
    4:i32 email_status,
    5:i32 phone_status,
    6:i32 video_status,
    7:i32 scene_status,
    8:i32 limitOffset,
    9:i32 limitNum,
    10:string chengHuiTongTraceLog
}

struct UserAttestationDetailsStruct {
	1: i32 user_id,// 用户ID,
	2: i32 card_type,// 证件类型（对应字典表ID）
	3: string hs_card_type,//徽商证件类型
	4: string card_id,//证件号码
	5: string card_pic1,//证件照片1
	6: string card_pic2,//证件照片2
	7: string video_pic,//视频认识图片
	8: i32 real_status,//实名认证状态：对应字典39类值；（原值：0：未申请认证，1：申请中，2：已经认证,3未通过认证）
	9: i32 real_passtime,//通过时间
	10: i32 email_status,//邮箱认证状态（0：未申请认证，1：申请中，2：已经认证,3未通过认证）
	11: i32 email_passtime,//通过时间
	12: i32 phone_status,//手机认证状态（0：未申请认证，1：申请中，2：已经认证,3未通过认证）
	13: i32 phone_passtime,//通过时间
	14: i32 video_status,//视频认证状态（0：未申请认证，1：申请中，2：已经认证,3未通过认证）
	15: i32 video_passtime,//通过时间,
	16: i32 scene_status,//现场认证状态（0：未申请认证，1：申请中，2：已经认证,3未通过认证）
	17: i32 scene_passtime,//通过时间
	18: i32 vip_status,//VIP状态（0：未申请认证，1：申请中，2：已经认证,3未通过认证）
	19: i32 vip_passtime,//VIP通过时间
	20: i32 vip_verifytime//VIP有效期时间
}

struct UserAttestionListResponseStruct {
    1:i32 status,
    2:string msg,
    3:i32 total,
    4:list<UserAttestationDetailsStruct> UserAttestionList
}

service UserAttestionListThriftService {
    UserAttestionListResponseStruct userAttestionList(1:UserAttestionListRequestStruct requestObj)
}
//查询总数total的sql = "select count(1) from jl_user_attestation UA LEFT JOIN  jl_user U ON UA.user_id=U.id LEFT JOIN jl_glossary G on UA.card_type=G.id where 拼接的where条件"

//查询列表数据的sql = "select UA.*,U.username,U.realname,U.phone,U.email,G.name from jl_user_attestation UA LEFT JOIN  jl_user U ON UA.user_id=U.id LEFT JOIN jl_glossary G on UA.card_type=G.id where 拼接的where条件 order by U.id desc limit $limitOffset,$limitNum"

//        $where = array();
//        if (isset($_GET['username']) && $_GET['username']) {
//            $where[] = "U.username like '%" . $_GET['username'] . "%'";
//        }
//        if (isset($_GET['realname']) && $_GET['realname']) {
//            $where[] = "U.realname like '%" . $_GET['realname'] . "%'";
//        }
//        if (isset($_GET['real_status']) && $_GET['real_status'] && $_GET['real_status'] != -1) {
//            $where[] = "UA.real_status='" . $_GET['real_status'] . "'";
//        }
//        if (isset($_GET['email_status']) && $_GET['email_status'] && $_GET['email_status'] != -1) {
//            $where[] = "UA.email_status='" . $_GET['email_status'] . "'";
//        }
//        if (isset($_GET['phone_status']) && $_GET['phone_status'] && $_GET['phone_status'] != -1) {
//            $where[] = "UA.phone_status='" . $_GET['phone_status'] . "'";
//        }
//        if (isset($_GET['video_status']) && $_GET['video_status'] && $_GET['video_status'] != -1) {
//            $where[] = "UA.video_status='" . $_GET['video_status'] . "'";
//        }
//        if (isset($_GET['scene_status']) && $_GET['scene_status'] && $_GET['scene_status'] != -1) {
//            $where[] = " UA.scene_status='" . $_GET['scene_status'] . "'";
//        }