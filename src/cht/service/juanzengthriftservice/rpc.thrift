namespace go juanzengthriftservice

struct RequestStruct {
	1: i32 user_id,
	2: string content,
	3: string chengHuiTongTraceLog
}

// 1 messlist留言 返回结构体
struct MesslistResultStruct {
    1: string username,
    2: string avatar,
    3: i32 addtime,
	4: string content,
    5: string reply,
	6: string up_content,
	7: string up_reply
}

// 2 fundlist返回结构体
struct FundlistResultStruct {
    1: i32 type,
	2: i32 addtime,
    3: string username,
	4: string money
}

// 3 numlist 返回结构体
struct NumlistResultStruct {
	1: string num,
	2: string money
}

struct JuanzengResponseStruct {
	1:list<MesslistResultStruct> messlist,
	2:list<FundlistResultStruct> fundlist,
	3:NumlistResultStruct numlist,
	4:string tzr                 
}

service JuanzengThriftService {
	JuanzengResponseStruct getInfo(1: RequestStruct requestObj)
	i32 addMess(1: RequestStruct requestObj)
}
