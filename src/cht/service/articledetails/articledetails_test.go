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

func NewNextRequestStruct(id, cateId, reqType, addtime int32, sort int32, prefix string, isApp int32) *NextRequestStruct {
	return &NextRequestStruct{
		ID:      id,
		Cateid:  cateId,
		Type:    reqType,
		Addtime: addtime,
		Sort:    sort,
		Prefix:  prefix,
		IsApp:   isApp,
	}
}

func TestGetArticleDetails(t *testing.T) {
	adrs := NewArticleDetailsRequestStruct(36)
	ads := articledetailsservice{}
	res, err := ads.GetArticleDetails(adrs)
	if err != nil {
		t.Fatalf("TestGetArticleDetails failed", err)
	}
	t.Logf("TestGetArticleDetails return value:%v", res)
}

func TestUpdateReadNum(t *testing.T) {
	adrs := NewArticleDetailsRequestStruct(641)
	ads := articledetailsservice{}
	res, err := ads.UpdateReadNum(adrs)
	if err != nil {
		t.Fatalf("TestUpdateReadNum failed", err)
	}
	t.Logf("TestUpdateReadNum return value:%v", res)
}

func TestGetPrevArticle(t *testing.T) {
	adrs := NewNextRequestStruct(2589, 5, 1, 1484550650, 1, "wz", 1)
	ads := articledetailsservice{}
	res, err := ads.PrevArticle(adrs)
	if err != nil {
		t.Fatalf("TestGetPrevArticle failed", err)
	}
	t.Logf("TestGetPrevArticle return value:%v", res)
}

func TestGetNextArticle(t *testing.T) {
	adrs := NewNextRequestStruct(2495, 8, 1, 1413881500, 1, "wz", 1)
	ads := articledetailsservice{}
	res, err := ads.NextArticle(adrs)
	if err != nil {
		t.Fatalf("TestGetNextArticle failed", err)
	}
	t.Logf("TestGetNextArticle return value:%v", res)
}
