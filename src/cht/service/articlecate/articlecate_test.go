package articlecate

import (
	_ "cht/initial"
	"testing"
)

func NewArticleCateListRequestStruct(id int32, name string, keywords string, status int32) *ArticleCateListRequestStruct {
	return &ArticleCateListRequestStruct{
		ID:       id,
		Name:     name,
		Keywords: keywords,
		Status:   status,
	}
}

func TestGetArticleCateList(t *testing.T) {
	aclr := NewArticleCateListRequestStruct(0, "", "服务中心", -1)
	acs := articlecateservice{}
	res, _ := acs.GetArticleCateList(aclr)
	if res.Status != QUERY_ARTICLE_CATE_SUCCESS {
		t.Fatalf("TestGetArticleCateList failed")
	}
	t.Logf("TestGetArticleCateList return value:%v", res)
}
