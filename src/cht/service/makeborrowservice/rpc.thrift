namespace php Borrow.MakeBorrow
namespace go  makeborrowservice

struct MakeBorrowRequestStruct {
   1: i32 id
	2: i32 borrow_type = 0
	3: i32 user_id = 0
    4: string title = " "
	5: string content  = " "
	6: string litpic  = " "
	7: i32    borrow_use = 0
	8: i32 	is_datetype = 0
	9: i32	time_limit = 1
	10:i32	style = 0
	11:string account = "0.00"
	12:string account_tender = "0.00"
	13:string apr = "0.0000"
	14:string apr_add = "0.0000"
	15:string mortgage_file = " "
	16:i32	is_dxb = 0
	17:string pwd = " "
	18:string lowest_account = "50.00"
	19:string most_account = "0.00"
	20:i32    valid_time = 1
	21:i32    award = 0
	22:string bonus = "0.00"
	23:i32	  is_false = 0
	24:i32	  open_account = 1
	25:i32	open_borrow = 1
	26:i32 		open_tender = 1
	27:i32  open_credit = 1
	28:i32 	open_ziliao = 1
	29:i32  material = 0
	30:i32	addtime = 0
	31:string addip = " "
	32:i32   status = 0
	33:i32	ruten_allnumber = 0
	34:i32   ruten_number = 0
	35:i32   verify_user = 0
	36:i32   verify_time = 0
	37:string verify_remark= " "
	38:i32   review_user = 0
	39:i32   review_time_local = 0
	40:i32   review_time = 0
	41:string secured = " "
	42:string zhuanrangren = " "
	43:i32    huodong = 0
	44:string sign_date = " "
	45:i32    subledger = 0
	46:i32    repay_sign = 0
	47:i32    auto_tender_lock = 0
	48:i32    is_auto = 0
	49:i32    is_check = 0
	50:i32    review_lock = 0
	51:string fee_rate = "0.0000"
	52:string borrow_name = " "
}

struct MakeBorrowResponseStruct {
     1:i32 status
	 2:string msg
}

service MakeBorrowThriftService {
    MakeBorrowResponseStruct makeBorrow(1:MakeBorrowRequestStruct requestObj)
}
