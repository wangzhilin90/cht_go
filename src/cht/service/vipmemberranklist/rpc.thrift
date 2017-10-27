//客户管理---VIP会员等级
namespace php User.VipMemberRankList
namespace go  vipmemberranklist

struct VipMemberRankListRequestStruct {
	1:i32 type,
	2:string keywords,
	3:i32 limitOffset,
	4:i32 limitNum,
	5:string chengHuiTongTraceLog
}

struct VipMemberRankDetailsStruct {
	1:i32 user_id,
	2:string money,
	3:string money_freeze,
	4:string money_usable,
	5:string money_wait,
	6:i32 cash_time,
	7:string return_money,
	8:i32 vip_level,
	9:string vip_wait_money,
	10:string hsreturn_money,
	11:string hsmoney_freeze,
	12:string hsmoney_wait,
	13:string username,
	14:string realname,
	15:i32 addtime
}

struct VipMemberRankListReponseStruct {
	1:i32 status,
	2:list<VipMemberRankDetailsStruct> VipMemberRankList
	3:i32 total_num,
	4:string msg,
}

service VipMemberRankListThriftService {
    VipMemberRankListReponseStruct getVipMemberRankList (1:VipMemberRankListRequestStruct requestObj)
}


//		  $where = array();
//        if (isset($_GET['type']) && !empty($_GET['keywords'])) {
//            if ($_GET['type'] == 1) {
//                $where[] = " U.username='" . $_GET['keywords'] . "'";
//            } else {
//                $where[] = " U.realname='" . $_GET['keywords'] . "'";
//            }
//        }

//总数：SELECT COUNT(1) FROM jl_user U LEFT JOIN jl_account A ON A.user_id=U.id
//列表：SELECT A.*,U.username,U.realname,U.addtime FROM jl_user U LEFT JOIN jl_account A ON A.user_id=U.id ORDER BY A.vip_level DESC,A.vip_wait_money DESC LIMIT 0,20
