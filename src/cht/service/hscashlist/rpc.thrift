//资金管理---徽商提现记录
namespace php Hsbank.HsCashList
namespace go  hscashlist

struct HsCashListRequestStruct{
	1:i32 start_time,
	2:i32 end_time,
	3:i32 timetype,
	4:i32 utype,//借款人,默认0：不限
	5:i32 type,
	6:string keywords,
	7:i32 pay_way,//提现途径，默认-1不限
	8:i32 status,//提现状态，默认-1不限
	9:i32 is_export,//是否导出，默认不导出 0
	10:i32 limitOffset,
	11:i32 limitNum,
	12:string chengHuiTongTraceLog
}

struct HsCashListResultStruct{
	1:i32 id,
	2:i32 user_id,
	3:string order_sn, // '提现订单号',
	4:string money, // '提现金额',
	5:string credited, // 实际到账金额
	6:string fee, // 提现手续费
	7:string use_return_money, //使用回款金额
	8:i32 use_free_num, //是否使用免费提现次数
	9:i32 addtime, //'添加时间',
	10:i32 status, // '状态：0审核中，1到账成功，2审核失败',
	11:i32 pay_way, // 提现途径：对应字典51类值；（原值：0为PC，1为APP，2微信）
	12:i32 deal_time, // 到账成功或审核失败的处理时间
	13:string fail_result, // 提现失败错误码
	14:string username,
	15:string realname,
	16:i32 regtime
}

struct HsCashListResponseStruct{
	1:i32 status,
	2:string msg,
	3:i32 total_num,
	4:list<HsCashListResultStruct> HsCashList
}

service HsCashListThriftService {
    HsCashListResponseStruct getHsCashList (1:HsCashListRequestStruct requestObj)
}


//   $where = array();
//        if (isset($_GET['start_time']) && $_GET['start_time']) {
//            $start_time = strtotime($_GET['start_time']);
//            if (empty($_GET['timetype'])) {
//                $where[] = "HC.addtime>='$start_time'";
//            } else {
//                $where[] = "HC.deal_time>='$start_time'";
//            }
//        }
//        if (isset($_GET['end_time']) && $_GET['end_time']) {
//            $end_time = strtotime($_GET['end_time']);
//            if ($start_time <= $end_time) {
//                $end_time = strtotime("next day", $end_time);
//                if (empty($_GET['timetype'])) {
//                    $where[] = "HC.addtime<'$end_time'";
//                } else {
//                    $where[] = "HC.deal_time<'$end_time'";
//                }
//            }
//        }
//        if(isset($_GET['utype']) && $_GET['utype']){
//            $is_borrower = intval($_GET['utype']);
//            if($is_borrower == 1){//保胜借款人
//                $where[] = 'U.is_borrower>0';
//            }
//            if($is_borrower == 640){//普通客户
//                 $where['U.is_borrower'] = 640;
//            }
//            if($is_borrower == 641){//深圳保胜
//                $where['U.is_borrower'] = 641;
//            }
//            if($is_borrower == 642){//贵州保胜
//                $where['U.is_borrower'] = 642;
//            }
//            if($is_borrower == 643){//广州保胜
//                $where['U.is_borrower'] = 643;
//            }
//        }
//        if (isset($_GET['type']) && isset($_GET['keywords']) && !empty($_GET['keywords'])) {
//            if ($_GET['type'] == 1) {
//                $where[] = "U.username='" . $_GET['keywords'] . "'";
//            } elseif ($_GET['type'] == 2) {
//                $where[] = "U.realname='" . $_GET['keywords'] . "'";
//            } elseif ($_GET['type'] == 3) {
//                $where[] = "HC.order_sn='" . $_GET['keywords'] . "'";
//            }
//        }
//        if (isset($_GET['pay_way']) && $_GET['pay_way'] != -1) {
//            $where[] = "HC.pay_way='" . intval($_GET['pay_way']) . "'";
//        }
//        if (isset($_GET['status']) && $_GET['status'] != -1) {
//            $where[] = "HC.status='" . intval($_GET['status']) . "'";
//        }

//导出：SELECT HC.*,U.username,U.realname,U.addtime AS regtime FROM jl_hs_cash HC LEFT JOIN jl_user U ON HC.user_id=U.id where $where ORDER BY HC.id DESC

//统计总数:SELECT COUNT(1) FROM jl_hs_cash HC LEFT JOIN jl_user U ON HC.user_id=U.id
//SELECT HC.*,U.username,U.realname,U.addtime AS regtime FROM jl_hs_cash HC LEFT JOIN jl_user U ON HC.user_id=U.id where $where ORDER BY HC.id DESC LIMIT 0,20




