package collection

import (
	_ "cht/initial"
	"testing"
)

func NewUserCollectionListRequest(user_id, start, end, search, status, offset, limit int32, zhuangrangren int32, borrowid string, tenderID int32, oldUserId int32) *UserCollectionListRequest {
	return &UserCollectionListRequest{
		UserID:            user_id,
		Starttime:         start,
		Endtime:           end,
		SearchTime:        search,
		State:             status,
		LimitOffset:       offset,
		LimitNum:          limit,
		CheckZhuanrangren: zhuangrangren,
		Borrowid:          borrowid,
		TenderID:          tenderID,
		CheckOldUserID:    oldUserId,
	}
}

func TestGetCollectionInfo(t *testing.T) {
	trr := NewUserCollectionListRequest(2, 0, 0, 0, 0, 1, 1, 2, "CHT00011", 3, 20)
	res, num, err := GetCollectionInfo(trr)
	if err != nil {
		t.Fatalf("TestGetCollectionInfo failed ", err)

	}
	t.Log("TestGetCollectionInfo res ", res, num)
}
