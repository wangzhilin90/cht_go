package makeborrowservice

import (
	_ "cht/initial"
	"testing"
)

func NewMakeBorrowRequestStruct(userID int32, borrowtype int32, borrowUse int32) *MakeBorrowRequestStruct {
	return &MakeBorrowRequestStruct{
		BorrowType:   borrowtype,
		UserID:       userID,
		BorrowUse:    borrowUse,
		Content:      "后台测试-信3",
		Title:        "biaoti",
		Account:      "20000",
		VerifyRemark: "审核意见",
		MostAccount:  "100",
		Secured:      "23k4",
	}
}

func TestMakeBorrow(t *testing.T) {
	brs := NewMakeBorrowRequestStruct(29, 3, 0)
	bs := borrowservice{}
	res, _ := bs.makeBorrow(brs)
	if res.Status != ISSURE_SUCCESS {
		t.Fatal("TestMakeBorrowNotDepositAccount failed:%v", res)
	}
	t.Logf("TestMakeBorrowNotDepositAccount return value:%v", res)
}
