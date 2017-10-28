//客户管理---客服值班---删除值班
namespace php User.KefuDutyDelete
namespace go  kefudutydelete

struct KefuDutyDeleteRequestStruct {
	1:string idstr,//需要删除的id字符串 ,例如： 2或者 5,7,9,11
	2:string chengHuiTongTraceLog
}

struct KefuDutyDeleteResponseStruct {
	1:i32 status,
	2:string msg
}

service KefuDutyDeleteThriftService {
    KefuDutyDeleteResponseStruct deleteKefuDuty (1:KefuDutyDeleteRequestStruct requestObj)
}


//DELETE FROM `jl_customer_plan` WHERE id in(203)

//删除计划
//public function delplan(){
//    $id = intval($_GET['id']);
//    if(!$id){
//        alert("参数错误");
//    }
//    $result = D('customerplan')->delPlans(array('id'=>$id));
//    if(!$result){
//         alert("删除失败");
//    }
//    header('location:' . url('', '', 'dutyList'));
//    exit;
//}