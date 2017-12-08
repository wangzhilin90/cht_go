package collection

import (
	_ "cht/initial"
	"testing"
)

func NewCollectionRequest(user_id, start, end, search, status, offset, limit int32, zhuangrangren int32, borrowid string) *CollectionRequest {
	return &CollectionRequest{
		UserID:            user_id,
		Starttime:         start,
		Endtime:           end,
		SearchTime:        search,
		State:             status,
		LimitOffset:       offset,
		LimitNum:          limit,
		CheckZhuanrangren: zhuangrangren,
		Borrowid:          borrowid,
	}
}

func TestGetCollectionInfo(t *testing.T) {
	trr := NewCollectionRequest(2, 0, 0, 0, 0, 1, 1, 2, "CHT00011")
	res, num, err := GetCollectionInfo(trr)
	if err != nil {
		t.Fatalf("TestGetCollectionInfo failed ", err)

	}
	t.Log("TestGetCollectionInfo res ", res, num)
}
