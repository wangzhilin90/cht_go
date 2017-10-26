//客户管理---专属客服---会员启用、禁用
namespace php User.CustomerUpdate
namespace go  customerupdate

struct CustomerUpdateRequestStruct {
	1:string id,
	2:i32 islock,
	3:string chengHuiTongTraceLog
}

struct CustomerUpdateResponseStruct {
	1:i32 status,
	2:string msg
}

service CustomerUpdateThriftService {
	CustomerUpdateResponseStruct updateCustomer (1:CustomerUpdateRequestStruct requestObj)
}


//	$role_list = D('User')->userDel('id in(' . $id . ')', $islock);
//	public function userDel($where = '', $islock = 1) {
//        if (empty($where))
//            return 0;
//        return $this->db->update('#@_user', array('islock' => intval($islock)), $where);
//    }