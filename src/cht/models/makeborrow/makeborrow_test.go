package makeborrow

import (
	_ "cht/initial"
	"testing"
)

func NewMakeBorrowRequest(userID int32, borrowtype int32, borrowUse int32) *MakeBorrowRequest {
	return &MakeBorrowRequest{
		BorrowType: borrowtype,
		UserID:     userID,
		BorrowUse:  borrowUse,
		// Title:      "biaoti",
		Title:   " ",
		Content: " ",
		Litpic:  " ",
		// TimeLimit:     1,
		Account:       "1000000.00",
		AccountTender: "0.00",
		Apr:           "0.0000",
		AprAdd:        "0.0000",
		MortgageFile:  " ",
		VerifyRemark:  " ",
		Pwd:           " ",
		LowestAccount: "50.00",
		MostAccount:   "0.00",
		ValidTime:     1,
		Bonus:         "0.00",
		OpenAccount:   1,
		OpenBorrow:    1,
		OpenTender:    1,
		OpenCredit:    1,
		OpenZiliao:    1,
		Addip:         " ",
		Secured:       "241234",
		Zhuanrangren:  " ",
		SignDate:      " ",
		FeeRate:       "0.00",
		BorrowName:    " ",
	}
}

func TestInsertBorrowTbl(t *testing.T) {
	mbr := NewMakeBorrowRequest(30, 5, 0)
	err := InsertBorrowTbl(mbr)
	if err != nil {
		t.Fatalf("TestInsertBorrowTbl failed :%v", err)
	}
}

func TestTInsertBorrowTbl(t *testing.T) {
	err := TInsertBorrowTbl()
	if err != nil {
		t.Fatalf("TInsertBorrowTbl failed :%v", err)
	}
}
