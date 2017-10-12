package articledetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/articledetails"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type articledetailsservice struct{}

func (ads *articledetailsservice) GetArticleDetails(requestObj *ArticleDetailsRequestStruct) (r *ArticleDetailsResultStruct, err error) {
	adrs := new(articledetails.ArticleDetailsRequestStruct)
	adrs.ID = requestObj.GetID()
	adrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, _ := articledetails.GetArticleDetails(adrs)

	var response ArticleDetailsResultStruct
	response.ID = res.ID
	response.Cateid = res.Cateid
	response.Title = res.Title
	response.Content = res.Content
	response.Keywords = res.Keywords
	response.Description = res.Description
	response.ImgURL = res.ImgURL
	response.Sort = res.Sort
	response.Status = res.Status
	response.Addtime = res.Addtime
	response.BannerURL = res.BannerURL
	response.Isbanner = res.Isbanner
	response.Type = res.Type
	response.Name = res.Name
	return &response, nil
}

func (ads *articledetailsservice) UpdateReadNum(requestObj *ArticleDetailsRequestStruct) (r int32, err error) {
	adrs := new(articledetails.ArticleDetailsRequestStruct)
	adrs.ID = requestObj.GetID()
	adrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, _ := articledetails.UpdateReadNum(adrs)
	return res, nil
}

func (ads *articledetailsservice) PrevArticle(requestObj *NextRequestStruct) (r *ArticleDetailsResultStruct, err error) {
	nrs := new(articledetails.NextRequestStruct)
	nrs.ID = requestObj.GetID()
	nrs.Cateid = requestObj.GetCateid()
	nrs.Type = requestObj.GetType()
	nrs.Addtime = requestObj.GetAddtime()
	nrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, _ := articledetails.GetPrevArticle(nrs)

	var response ArticleDetailsResultStruct
	response.ID = res.ID
	response.Cateid = res.Cateid
	response.Title = res.Title
	response.Content = res.Content
	response.Keywords = res.Keywords
	response.Description = res.Description
	response.ImgURL = res.ImgURL
	response.Sort = res.Sort
	response.Status = res.Status
	response.Addtime = res.Addtime
	response.BannerURL = res.BannerURL
	response.Isbanner = res.Isbanner
	response.Type = res.Type
	response.Name = res.Name
	return &response, nil
}

func (ads *articledetailsservice) NextArticle(requestObj *NextRequestStruct) (r *ArticleDetailsResultStruct, err error) {
	nrs := new(articledetails.NextRequestStruct)
	nrs.ID = requestObj.GetID()
	nrs.Cateid = requestObj.GetCateid()
	nrs.Type = requestObj.GetType()
	nrs.Addtime = requestObj.GetAddtime()
	nrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, _ := articledetails.GetNextArticle(nrs)

	var response ArticleDetailsResultStruct
	response.ID = res.ID
	response.Cateid = res.Cateid
	response.Title = res.Title
	response.Content = res.Content
	response.Keywords = res.Keywords
	response.Description = res.Description
	response.ImgURL = res.ImgURL
	response.Sort = res.Sort
	response.Status = res.Status
	response.Addtime = res.Addtime
	response.BannerURL = res.BannerURL
	response.Isbanner = res.Isbanner
	response.Type = res.Type
	response.Name = res.Name
	return &response, nil
}

/*开启获取文章详情服务*/
func StartArticleDetailsServer() {
	zkServers := []string{"192.168.8.208:2181"}
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30031"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/ArticleDetailsThriftService/providers"
	err = zkclient.RegisterNode(conn, servicename, listenAddr)
	if err != nil {
		Logger.Fatalf("RegisterNode failed", err)
	}

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &articledetailsservice{}
	processor := NewArticleDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
