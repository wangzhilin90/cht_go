//客户管理---专属客服---列表
namespace php User.CustomerList
namespace go  customerlist

struct CustomerListRequestStruct{
	1:i32 customer,
	2:i32 start_time,
	3:i32 end_time,
	4:i32 islock,
	5:string username,
	6:i32 is_export,//是否导出，默认不导出 0
	7:i32 limitOffset,
	8:i32 limitNum,
	9:string chengHuiTongTraceLog
}

struct CustomerListResultStruct{
	1:i32 id,
	2:string username,
	3:string password,
	4:string paypassword,
	5:i32 point,
	6:string email,
	7:string avatar,
	8:i32 sex,
	9:string realname,
	10:string phone,
	11:string tel,
	12:i32 birthday,
	13:i32 nation,
	14:i32 province,
	15:i32 city,
	16:i32 area,
	17:string address,
	18:i32 customer,
	19:i32 logintime,
	20:string loginip,
	21:i32 addtime,
	22:string addip,
	23:i32 islock,
	24:i32 isvest,
	25:i32 os_type,
	26:string device_token,
	27:string weinxin_id,
	28:i32 bind_time,
	29:string invitation_code,
	30:string source,
	31:string hsid,
	32:i32 g_status,
	33:string g_password,
	34:string auto_protocol_code,
	35:i32 is_borrower,
	36:i32 is_worker, //保胜业务员类型，对应字典64类值：（原值：1是深圳，2是广州，3是贵州
	37:i32 hswaitactivate
}

struct CustomerListResponseStruct{
	1:i32 status,
	2:string msg,
	3:i32 total_num,
	4:list<CustomerListResultStruct> CustomerList
}

service CustomerListThriftService {
	CustomerListResponseStruct getCustomerList (1:CustomerListRequestStruct requestObj)
}

//		  $where = array();
//        $where[] = " customer=" . intval($this->admin['id']);
//        if (isset($_GET['start_time']) && $_GET['start_time']) {
//            $start_time = strtotime($_GET['start_time']);
//            $where[] = " `addtime` > '$start_time'";
//        }
//        if (isset($_GET['end_time']) && $_GET['end_time']) {
//            $end_time = strtotime($_GET['end_time']);
//            if ($start_time <= $end_time) {
//                $end_time = strtotime("next day", $end_time);
//                $where[] = " `addtime` < '$end_time'";
//            }
//        }
//        if (isset($_GET['islock']) && $_GET['islock'] != -1) {
//            $where[] = "islock=" . intval($_GET['islock']);
//        }
//        if (isset($_GET['username']) && $_GET['username']) {
//            $where[] = " username like '%" . $_GET['username'] . "%'";
//        }


//总数：SELECT COUNT(1) FROM jl_user WHERE customer=1
//列表：SELECT * FROM jl_user U WHERE customer=1 ORDER BY id DESC LIMIT 0,20

//导出 不加分页




