package makeborrow

import (
	_ "cht/initial"
	"testing"
)

func NewBorrowStruct(userID int32) *Borrow {
	id, _ := GetLatestBorrowID()

	return &Borrow{
		ID:           id,
		UserID:       userID,
		Title:        "biaoti1",
		Content:      "hello",
		VerifyRemark: "审核",
		MostAccount:  "2000000",
		OpenAccount:  23,
		Subledger:    32,
	}
}

func TestGetLatestBorrowID(t *testing.T) {
	res, err := GetLatestBorrowID()
	if err != nil {
		t.Fatalf("TestGetLatestID failed:%v", err)
	}
	t.Logf("TestGetLatestID return value:%v", res)
}

func TestInsertBorrowTbl(t *testing.T) {
	bs := NewBorrowStruct(30)
	num, err := InsertBorrowTbl(bs)
	if err != nil {
		t.Fatalf("TInsertBorrowTbl failed :%v", err)
	}
	t.Logf("TestInsertBorrowTbl last input num:%v", num)
}
