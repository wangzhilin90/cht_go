package articledetails

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	"cht/models/articledetails"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_ARTICLE_DETAILS_SUCCESS = 1000
	QUERY_ARTICLE_DETAILS_FAILED  = 1001
)

var ArticleStatus = map[int]string{
	QUERY_ARTICLE_DETAILS_SUCCESS: "查询文章详情成功",
	QUERY_ARTICLE_DETAILS_FAILED:  "查询文章详情失败",
}

const (
	QUERY_PREV_ARTICLE_SUCCESS = 1000
	QUERY_PREV_ARTICLE_FAILED  = 1001
)

var PrevArticleStatus = map[int]string{
	QUERY_PREV_ARTICLE_SUCCESS: "查询上一篇文章成功",
	QUERY_PREV_ARTICLE_FAILED:  "查询下一篇文章失败",
}

const (
	QUERY_NEXT_ARTICLE_SUCCESS = 1000
	QUERY_NEXT_ARTICLE_FAILED  = 1001
)

var NextArticleStatus = map[int]string{
	QUERY_NEXT_ARTICLE_SUCCESS: "查询下一篇文章成功",
	QUERY_NEXT_ARTICLE_FAILED:  "查询下一篇文章失败",
}

type articledetailsservice struct{}

func (ads *articledetailsservice) GetArticleDetails(requestObj *ArticleDetailsRequestStruct) (r *ArticleDetailsResponseStruct, err error) {
	Logger.Infof("GetArticleDetails requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*ArticleDetailsRequestStruct)
	adrs := new(articledetails.ArticleDetailsRequestStruct)
	adrs.ID = requestObj.GetID()
	adrs.Status = requestObj.GetStatus()
	adrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, err := articledetails.GetArticleDetails(adrs)
	if err != nil {
		Logger.Errorf("GetArticleDetails get details failed:%v", err)
		return &ArticleDetailsResponseStruct{
			Status: QUERY_ARTICLE_DETAILS_FAILED,
			Msg:    ArticleStatus[QUERY_ARTICLE_DETAILS_FAILED],
		}, nil
	}
	var response ArticleDetailsResponseStruct
	if res != nil {
		adst := new(ArticleDetailsStruct)
		adst.ID = res.ID
		adst.Cateid = res.Cateid
		adst.Title = res.Title
		adst.Content = res.Content
		adst.Keywords = res.Keywords
		adst.Description = res.Description
		adst.ImgURL = res.ImgURL
		adst.Sort = res.Sort
		adst.Status = res.Status
		adst.Addtime = res.Addtime
		adst.BannerURL = res.BannerURL
		adst.Isbanner = res.Isbanner
		adst.Type = res.Type
		adst.Name = res.Name
		adst.ReadNum = res.ReadNum
		response.ArticleDetails = adst
	}
	response.Status = QUERY_ARTICLE_DETAILS_SUCCESS
	response.Msg = ArticleStatus[QUERY_ARTICLE_DETAILS_SUCCESS]
	Logger.Debugf("GetArticleDetails response:%v", response)
	return &response, nil
}

func (ads *articledetailsservice) UpdateReadNum(requestObj *ArticleDetailsRequestStruct) (r int32, err error) {
	Logger.Infof("UpdateReadNum requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*ArticleDetailsRequestStruct)
	adrs := new(articledetails.ArticleDetailsRequestStruct)
	adrs.ID = requestObj.GetID()
	adrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, _ := articledetails.UpdateReadNum(adrs)

	return res, nil
}

func (ads *articledetailsservice) PrevArticle(requestObj *NextRequestStruct) (r *ArticleDetailsResponseStruct, err error) {
	Logger.Infof("PrevArticle requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*NextRequestStruct)
	nrs := new(articledetails.NextRequestStruct)
	nrs.ID = requestObj.GetID()
	nrs.Cateid = requestObj.GetCateid()
	nrs.Type = requestObj.GetType()
	nrs.Addtime = requestObj.GetAddtime()
	nrs.Sort = requestObj.GetSort()
	nrs.Prefix = requestObj.GetPrefix()
	nrs.IsApp = requestObj.GetIsApp()
	nrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, err := articledetails.GetPrevArticle(nrs)
	if err != nil {
		Logger.Errorf("PrevArticle failed:%v", err)
		return &ArticleDetailsResponseStruct{
			Status: QUERY_PREV_ARTICLE_FAILED,
			Msg:    PrevArticleStatus[QUERY_PREV_ARTICLE_FAILED],
		}, nil
	}

	var response ArticleDetailsResponseStruct
	if res != nil {
		adst := new(ArticleDetailsStruct)
		adst.ID = res.ID
		adst.Cateid = res.Cateid
		adst.Title = res.Title
		adst.Content = res.Content
		adst.Keywords = res.Keywords
		adst.Description = res.Description
		adst.ImgURL = res.ImgURL
		adst.Sort = res.Sort
		adst.Status = res.Status
		adst.Addtime = res.Addtime
		adst.BannerURL = res.BannerURL
		adst.Isbanner = res.Isbanner
		adst.Type = res.Type
		adst.Name = res.Name
		adst.ReadNum = res.ReadNum
		response.ArticleDetails = adst
	}
	response.Status = QUERY_PREV_ARTICLE_SUCCESS
	response.Msg = PrevArticleStatus[QUERY_PREV_ARTICLE_SUCCESS]
	Logger.Debugf("PrevArticle response:%v", response)
	return &response, nil
}

func (ads *articledetailsservice) NextArticle(requestObj *NextRequestStruct) (r *ArticleDetailsResponseStruct, err error) {
	Logger.Infof("NextArticle requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*NextRequestStruct)
	nrs := new(articledetails.NextRequestStruct)
	nrs.ID = requestObj.GetID()
	nrs.Cateid = requestObj.GetCateid()
	nrs.Type = requestObj.GetType()
	nrs.Addtime = requestObj.GetAddtime()
	nrs.Sort = requestObj.GetSort()
	nrs.Prefix = requestObj.GetPrefix()
	nrs.IsApp = requestObj.GetIsApp()
	nrs.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, err := articledetails.GetNextArticle(nrs)
	if err != nil {
		Logger.Errorf("NextArticle failed:%v", err)
		return &ArticleDetailsResponseStruct{
			Status: QUERY_NEXT_ARTICLE_FAILED,
			Msg:    NextArticleStatus[QUERY_NEXT_ARTICLE_FAILED],
		}, nil
	}

	var response ArticleDetailsResponseStruct
	if res != nil {
		adst := new(ArticleDetailsStruct)
		adst.ID = res.ID
		adst.Cateid = res.Cateid
		adst.Title = res.Title
		adst.Content = res.Content
		adst.Keywords = res.Keywords
		adst.Description = res.Description
		adst.ImgURL = res.ImgURL
		adst.Sort = res.Sort
		adst.Status = res.Status
		adst.Addtime = res.Addtime
		adst.BannerURL = res.BannerURL
		adst.Isbanner = res.Isbanner
		adst.Type = res.Type
		adst.Name = res.Name
		adst.ReadNum = res.ReadNum
		response.ArticleDetails = adst
	}
	response.Status = QUERY_NEXT_ARTICLE_SUCCESS
	response.Msg = NextArticleStatus[QUERY_NEXT_ARTICLE_SUCCESS]
	Logger.Debugf("NextArticle response:%v", response)
	return &response, nil
}

/*开启获取文章详情服务*/
func StartArticleDetailsServer() {
	zkServers := zkclient.ZkServerAddress
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
		Logger.Fatalf("RegisterNode %v failed", servicename, err)
	}

	go func() {
		time.Sleep(time.Second * 2)
		err = zkclient.WatchNode(conn, servicename, listenAddr)
		if err != nil {
			Logger.Fatalf("WatchNode %v failed:%v", servicename, err)
		}
	}()

	serverTransport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		Logger.Fatal("NewTServerSocket failed", err)
	}

	handler := &articledetailsservice{}
	processor := NewArticleDetailsThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
