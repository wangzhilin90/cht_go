package articledetails

import (
	_ "cht/initial"
	"testing"
)

func NewArticleDetailsRequestStruct(id int32) *ArticleDetailsRequestStruct {
	return &ArticleDetailsRequestStruct{
		ID: id,
	}
}

func NewNextRequestStruct(id, cateId, reqType, addtime int32) *NextRequestStruct {
	return &NextRequestStruct{
		ID:      id,
		Cateid:  cateId,
		Type:    reqType,
		Addtime: addtime,
	}
}

func TestGetArticleDetails(t *testing.T) {
	adrs := NewArticleDetailsRequestStruct(36)
	res, err := GetArticleDetails(adrs)
	if err != nil {
		t.Fatalf("TestGetArticleDetails failed", err)
	}
	t.Logf("TestGetArticleDetails return value:%v", res)
}

func TestUpdateReadNum(t *testing.T) {
	adrs := NewArticleDetailsRequestStruct(641)
	res, err := UpdateReadNum(adrs)
	if err != nil {
		t.Fatalf("TestUpdateReadNum failed", err)
	}
	t.Logf("TestUpdateReadNum return value:%v", res)
}

func TestGetPrevArticle(t *testing.T) {
	// adrs := NewNextRequestStruct(3941, 10, 1, 1413881500)
	adrs := NewNextRequestStruct(2495, 8, 1, 1413881500)
	res, err := GetPrevArticle(adrs)
	if err != nil {
		t.Fatalf("TestGetPrevArticle failed", err)
	}
	t.Logf("TestGetPrevArticle return value:%v", res)
}

func TestGetNextArticle(t *testing.T) {
	adrs := NewNextRequestStruct(2495, 8, 1, 1413881500)
	res, err := GetNextArticle(adrs)
	if err != nil {
		t.Fatalf("TestGetNextArticle failed", err)
	}
	t.Logf("TestGetNextArticle return value:%v", res)
}
