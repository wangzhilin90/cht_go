namespace php Log.VipCustomerLogList
namespace go  vipcustomerloglist

//专属客服日志记录

struct VipCustomerLogListRequestStruct {
    1:i32 start_time,
    2:i32 end_time,
    3:string keywords,
    4:i32 type,//查询类型,默认为1(会员名称)
    5:i32 limitOffset,
    6:i32 limitNum,
    7:string chengHuiTongTraceLog
}

struct VipCustomerDetailsStruct {
    1:i32 id,
	2:i32 user_id,
	3:string username,
	4:string email,
	5:string realname,
	6:string phone,
	7:i32 scene_passtime,//通过时间
	8:i32 vip_status, //VIP状态
	9:i32 vip_passtime,//VIP通过时间
	10:i32 vip_verifytime,//VIP有效期时间
	11:i32 old_customer,//旧的客服ID
	12:i32 new_customer,//新的客服ID
	13:i32 updatetime,//VIP有效期时间
	14:string remark //记录操作用户
}

struct VipCustomerLogListResponseStruct {
	1:i32 status,
    2:list<VipCustomerDetailsStruct> VipCustomerLogList,
    3:i32 total_num,
	4:string msg,
}


service VipCustomerLogListThriftService {
    VipCustomerLogListResponseStruct getVipCustomerLogList (1:VipCustomerLogListRequestStruct requestObj)
}

//SQL:
//    SELECT COUNT(1) FROM  jl_change_vip WHERE $where ;
//    SELECT * FROM jl_change_vip WHERE $where ORDER BY id desc LIMIT $limitOffset , $limitNum;

//主要的业务逻辑

/** 获取所有的vip 变更记录* */
//   public function getAllchangevip () {
//       $where = array();
//       if (isset($_GET['start_time']) && $_GET['start_time']) {
//           $start_time = strtotime($_GET['start_time']);
//           $where[] = "vip_passtime>='$start_time'";
//        }
//        if (isset($_GET['end_time']) && $_GET['end_time']) {
//            $end_time = strtotime($_GET['end_time']);
//            if ($start_time <= $end_time) {
//                $end_time = strtotime("next day", $end_time);
//                $where[] = "vip_passtime<='$end_time'";
//            }
//        }
//        if (isset($_GET['keywords']) && $_GET['keywords']) {
//            if ($_GET['type'] == 1) {
//                $where[] = "username like '%" . $_GET['keywords'] . "%'";
//            }
//            if ($_GET['type'] == 2) {
//                $where[] = "user_id =" . intval($_GET['keywords']);
//            }
//            if ($_GET['type'] == 3) {
//                $where[] = "realname like '%" . $_GET['keywords'] . "%'";
//            }
//        }
//        $count = D("Changevip")->getCount($where);
//        $page = isset($_GET['page']) && $_GET['page'] ? intval($_GET['page']) : 1;
//        unset($_GET['page'], $_GET['m'], $_GET['c'], $_GET['a']);
//        $page = pagelist($count, $page, $_GET, 20);
//        $list = D("Changevip")->getlist('', $where, '', $page['limit']);
//        $kefu_old =  D("admin_user")->getUserlist(array('role_id' => 2, 'status' => 0), '', '', 'id');
//        $kefu= array();
//        if(!empty($kefu_old)){
//            foreach($kefu_old as $k=>$v){
//                $kefu[$v['id']] = $v;
//            }
//        }
//        include template('vip-allchange');
//    }