//客户管理---会员紧急联系人
namespace php User.MemberHelperList
namespace go memberhelperlist

struct MemberHelperListRequestStruct{
	1:i32 type,
	2:string keywords,
	3:i32 limitOffset,
	4:i32 limitNum,
	5:string chengHuiTongTraceLog
}

struct MemberHelperDetailsStruct {
	1:string linkman,
	2:string linkrelation, //联系人关系(对应字黄表)
	3:string linkphone,
	4:i32 updatetime,
	5:i32 id,
	6:string username,
	7:string realname,
	8:string phone,
	9:i32 customer //客服ID
}

struct MemberHelperListResponseStruct{
	1:i32 status,
	2:list<MemberHelperDetailsStruct> linkManList
	3:i32 total_num,
	4:string msg,
}

service MemberHelperListThriftService {
    MemberHelperListResponseStruct getMemberHelperList (1:MemberHelperListRequestStruct requestObj)
}

//        $where = array();
//        if (isset($_GET['type']) && !empty($_GET['keywords'])) {
//            if ($_GET['type'] == 1) {
//                $where[] = " U.username like '%" . $_GET['keywords'] . "%'";
//            } elseif ($_GET['type'] == 2) {
//                $where[] = " U.realname like '%" . $_GET['keywords'] . "%'";
//            } elseif ($_GET['type'] == 3) {
//                $where[] = " U.phone='" . $_GET['keywords'] . "'";
//            } else {
//                $where[] = "id='" . $_GET['keywords'] . "'";
//            }
//        }
//        $where[] = " UF.linkman <>'' ";

//总数：SELECT COUNT(1) FROM jl_user_field UF LEFT JOIN jl_user U ON UF.user_id=U.id WHERE UF.linkman <>''
//列表：SELECT UF.linkman,UF.linkrelation,UF.linkphone,UF.updatetime,U.id,U.username,U.realname,U.phone,U.customer FROM jl_user_field UF LEFT JOIN jl_user U ON UF.user_id=U.id WHERE UF.linkman <>'' ORDER BY id DESC LIMIT 0,20








