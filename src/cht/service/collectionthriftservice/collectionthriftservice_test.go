package collectionthriftservice

import (
	_ "cht/initial"
	"testing"
)

func NewCollectionRequestStruct(user_id, start, end, search, status, offset, limit int32, borrowid string) *CollectionRequestStruct {
	return &CollectionRequestStruct{
		UserID:      user_id,
		Starttime:   start,
		Endtime:     end,
		SearchTime:  search,
		State:       status,
		LimitOffset: offset,
		LimitNum:    limit,
		Borrowid:    borrowid,
	}
}

func TestGetCollectionList(t *testing.T) {
	tcs := NewCollectionRequestStruct(2, 1376279359, 1376299359, 0, 0, 0, 0, "CHT11")
	cs := collectionservice{}
	res, err := cs.GetCollectionList(tcs)
	if err != nil {
		t.Fatalf("TestGetCollectionList failed:", err)
	}
	t.Log("TestGetCollectionList res:", res)
}
