package makeborrowservice

import (
	_ "cht/initial"
	"testing"
)

func NewMakeBorrowRequestStruct(userID int32, borrowtype int32, borrowUse int32) *MakeBorrowRequestStruct {
	return &MakeBorrowRequestStruct{
		BorrowType: borrowtype,
		UserID:     userID,
		BorrowUse:  borrowUse,
		Title:      "biaoti",
		// Title:         ",",
		// Content:       ",",
		// Litpic:        ",",
		// TimeLimit:     1,
		Account: "1000000.00",
		// AccountTender: "0.00",
		// Apr:           "0.0000",
		// AprAdd:        "0.0000",
		// MortgageFile:  ",",
		// VerifyRemark:  ",",
		// Pwd:           ",",
		// LowestAccount: "50.00",
		// MostAccount:   "0.00",
		// ValidTime:     1,
		// Bonus:         "0.00",
		// OpenAccount:   1,
		// OpenBorrow:    1,
		// OpenTender:    1,
		// OpenCredit:    1,
		// OpenZiliao:    1,
		// Addip:         ",",
		// Secured:       "241234",
		// Zhuanrangren:  ",",
		// SignDate:      ",",
		FeeRate: "20.00",
		// BorrowName:    ",",
	}
}

func TestMakeBorrowNotDepositAccount(t *testing.T) {
	/*测试存管账户不存在*/
	brs := NewMakeBorrowRequestStruct(29, 5, 0)
	bs := borrowservice{}
	res, _ := bs.makeBorrow(brs)
	if res.Status != ISSURE_SUCCESS {
		t.Fatal("TestMakeBorrowNotDepositAccount failed")
	}
}

// func TestMakeBorrowAddCredit(t *testing.T) {
// 	/*测试不是加信贷*/
// 	brs := NewMakeBorrowRequestStruct(5004, 5, 0, "biaoti")
// 	bs := borrowservice{}
// 	res, _ := bs.makeBorrow(brs)
// 	if res.Status != NOT_DEPOSIT_ACCOUNT {
// 		t.Fatal("TestmakeBorrow failed")
// 	}
// }

// func TestMakeBorrowExceedReditLimit(t *testing.T) {
// 	/*测试不是加信贷*/
// 	brs := NewMakeBorrowRequestStruct(5003, 4, 0, "biaoti")
// 	bs := borrowservice{}
// 	res, _ := bs.makeBorrow(brs)
// 	if res.Status != NOT_DEPOSIT_ACCOUNT {
// 		t.Fatal("TestmakeBorrow failed")
// 	}
// }
