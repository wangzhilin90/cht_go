//徽商日志明细
namespace php Hsbank.HsLogList
namespace go  hsloglist

struct HsLogListRequestStruct {
	1:i32 start_time,
	2:i32 end_time,
	3:i32 type,//操作类型,默认-1：不限
	4:i32 type2,
	5:string kws,
	6:i32 utype,//角色,默认0：不限
	7:i32 is_export,//是否导出，默认不导出 0
	8:i32 limitOffset,
	9:i32 limitNum,
	10:i32 borrow_id,//标的ID
	11:string chengHuiTongTraceLog
}

struct HsLogDetailsStruct {
	1:i32 id,
	2:i32 user_id,
	3:string orderno, // '订单号',
	4:i32 type, // '类型',
	5:string money, // '操作金额',
	6:string freeze_money,// '冻结金额',
	7:string wait_money, //'待收金额',
	8:i32 addtime, //'操作时间',
	9:i32 toid, // '对应业务id',
	10:string remark, // '备注',
	11:string username,
	12:string realname,
	13:i32 regtime
}


struct HsLogListReponseStruct {
	1:i32 status,
	2:string msg,
	3:i32 total_num,
	4:list<HsLogDetailsStruct> HslogList
}


service HsLogListThriftService {
    HsLogListReponseStruct getHslogList(1:HsLogListRequestStruct requestObj)
}

//  $where = array();
//        if (isset($_GET['start_time']) && $_GET['start_time']) {
//            if (empty($_GET['timetype'])) {
//                $where[] = "HL.addtime>='$start_time'";
//            } 
//        }
//        if (isset($_GET['end_time']) && $_GET['end_time']) {
//            $end_time = strtotime($_GET['end_time']);
//            if ($start_time <= $end_time) {
//                $end_time = strtotime("next day", $end_time);
//                if (empty($_GET['timetype'])) {
//                    $where[] = "HL.addtime<'$end_time'";
//                } 
//            }
//        }
//        if (isset($_GET['type']) && $_GET['type'] != -1) {
//            $where[] = "HL.type=" . intval($_GET['type']);
//        }
//        if (isset($_GET['type2']) && isset($_GET['kws']) && !empty($_GET['kws'])) {
//            if ($_GET['type2'] == 1) {
//                $where[] = "U.username = '" . $_GET['kws'] . "'";
//            } elseif ($_GET['type2'] == 2) {
//                $where[] = "U.realname = '" . $_GET['kws'] . "'";
//            } elseif ($_GET['type2'] == 3) {
//                $where[] = "HL.user_id='" . $_GET['kws'] . "'";
//            } elseif ($_GET['type2'] == 4) {
//                $where[] = "HL.orderno='" . $_GET['kws'] . "'";
//            }
//        }
//        if(isset($_GET['utype']) && $_GET['utype']){
//            if(intval($_GET['utype']) == 1){//保胜借款人
//                $where[] = 'U.is_borrower>0';
//            }
//            if(intval($_GET['utype']) == 2){//普通客户
//                $where['U.is_borrower'] = 0;
//            }
//        }

//if (strexists($where, '.')) {
//            $table = '#@_hs_log HL LEFT JOIN #@_user U ON HL.user_id=U.id';
//        } else {
//            $table = '#@_hs_log';
//        }
//统计总数:SELECT COUNT(1) FROM jl_hs_log
//统计总数:SELECT COUNT(1) FROM #@_hs_log HL LEFT JOIN #@_user U ON HL.user_id=U.id WHERE U.is_borrower>0

//导出：
//SELECT HL.*,U.username,U.realname,U.addtime AS regtime FROM #@_hs_log HL LEFT JOIN #@_user U ON HL.user_id=U.id  where $where ORDER BY HL.id DESC
查询:
//SELECT HL.*,U.username,U.realname,U.addtime AS regtime FROM #@_hs_log HL LEFT JOIN #@_user U ON HL.user_id=U.id where $where ORDER BY HL.id DESC LIMIT 0,20 

