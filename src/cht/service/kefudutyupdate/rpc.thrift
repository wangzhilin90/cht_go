//客户管理---客服值班---修改值班
namespace php User.KefuDutyUpdate
namespace go  kefudutyupdate

struct KefuDutyUpdateRequestStruct {
	1:i32 id,
	2:string customer,
	3:i32 duty_time,
	4:string holiday_user,
	5:i32 is_rest,
	6:i32 starttime, //对应数据库“start”
	7:i32 endtime,   //对应数据库“end”
	8:string chengHuiTongTraceLog
}

struct KefuDutyUpdateResponseStruct {
	1:i32 status,
	2:string msg
}

service KefuDutyUpdateThriftService {
    KefuDutyUpdateResponseStruct updateKefuDuty (1:KefuDutyUpdateRequestStruct requestObj)
}

/**
 * 修改值班计划
 */
//public function editPlan(){
//    $id = isset($_GET['id']) && $_GET['id'] ? intval($_GET['id']) : intval($_POST['id']);
//    if (empty($id)) {
//        alert("参数错误");
//    }
//    if(isset($_POST['savetype'])){
//        $data = array();
//        $data['duty_time'] = intval(strtotime($_POST['dutytime']));
//        $data['is_rest'] = isset($_POST['isrest']) && $_POST['isrest']==1 ? 1:0;
//        $data['customer']= !$_POST['kefu'] ? '': implode(',', $_POST['kefu']);
//        $data['holiday_user'] = !$_POST['restkfval'] ? '': implode(',', $_POST['restkfval']);
//        $data['start'] = intval(strtotime(($_POST['dutytime'].' '.$_POST['start'])));
//        $data['end'] = intval(strtotime(($_POST['dutytime'].' '.$_POST['end'])));
//        if( !$data['duty_time'] || !$data['customer']){
//            alert( '指班人和值班时间不可以为空！');
//        }
//        $result = D('customerplan')->addPlan($data,array('id'=>$id));
//        if($result){
//            alert("设置成功", HTTP_REFERER, 1, 11, 'dutyList');
//        }else{
//            alert( '修改失败！');
//        }
//    }
//    $kefu = D("admin_user")->getUserlist(array('role_id' => 2, 'status' => 0), '', '', 'id');
//    $list = D('customerplan')->getPlan('*',array('id'=>$id),'',1);
//    $list['restkf']= explode(',', $list['holiday_user']);
//    $list['work']= explode(',', $list['customer']);
//    include template('customer-duty-edit');
//}