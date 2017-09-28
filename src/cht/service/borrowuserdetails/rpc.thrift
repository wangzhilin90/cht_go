namespace php User.BorrowUserDetails
namespace go borrowuserdetails
//发标服务---借款人查询服务

struct BorrowUserDetailsRequestStruct {
	1:string name,
	2:string chengHuiTongTraceLog
}

struct materialInfoStruct {
	1:i32 id,
	2:string name
}

struct BorrowUserDetailsStruct {
	1:i32 id,
	2:string realname,
	3:i32 is_borrower,
	4:string card_id,
	5:string credit,
	6:string guarantor,
	7:list < materialInfoStruct > materialList
}

struct BorrowUserDetailsResponseStruct {
	1:i32 status, //1000:"查询借款人信息成功" 1001:"无此借款人!"
	2:string msg,
	3:BorrowUserDetailsStruct BorrowUserDetails
}

service BorrowUserDetailsThriftService {
    BorrowUserDetailsResponseStruct getBorrowUserDetails (1: BorrowUserDetailsRequestStruct requestObj)
}