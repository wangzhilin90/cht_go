package articlecate

import (
	_ "cht/initial"
	"testing"
)

func NewArticleCateListRequest(id int32, name string, keywords string, status int32) *ArticleCateListRequest {
	return &ArticleCateListRequest{
		ID:       id,
		Name:     name,
		Keywords: keywords,
		Status:   status,
	}
}

func TestGetArticleCateList(t *testing.T) {
	aclr := NewArticleCateListRequest(0, "", "服务中心", -1)
	res, err := GetArticleCateList(aclr)
	if err != nil {
		t.Fatalf("TestGetArticleCateList failed:%v", err)
	}
	t.Logf("TestGetArticleCateList return value:%v", res)
}
