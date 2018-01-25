package articlecate

import (
	. "cht/common/logger"
	"cht/common/zkclient"
	ac "cht/models/articlecate"
	"cht/utils/filterspec"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

const (
	QUERY_ARTICLE_CATE_SUCCESS = 1000
	QUERY_ARTICLE_CATE_FAILED  = 1001
	QUERY_ARTICLE_CATE_EMPTY   = 1002
)

var Stat = map[int]string{
	QUERY_ARTICLE_CATE_SUCCESS: "查询文章分类表详情成功",
	QUERY_ARTICLE_CATE_FAILED:  "查询文章分类表详情失败",
	QUERY_ARTICLE_CATE_EMPTY:   "查询文章分类表详情为空",
}

type articlecateservice struct{}

func (acs *articlecateservice) GetArticleCateList(requestObj *ArticleCateListRequestStruct) (r *ArticleCateListResponseStruct, err error) {
	Logger.Infof("GetArticleCateList requestObj:%v", requestObj)
	requestObj = filterspec.FiterSpecialCharacters(requestObj).(*ArticleCateListRequestStruct)
	aclr := new(ac.ArticleCateListRequest)
	aclr.ID = requestObj.GetID()
	aclr.Name = requestObj.GetName()
	aclr.Keywords = requestObj.GetKeywords()
	aclr.Description = requestObj.GetDescription()
	aclr.Pid = requestObj.GetPid()
	aclr.Status = requestObj.GetStatus()
	aclr.ChengHuiTongTraceLog = requestObj.GetChengHuiTongTraceLog()
	res, err := ac.GetArticleCateList(aclr)
	if err != nil {
		Logger.Errorf("GetArticleCateList failed:%v", err)
		return &ArticleCateListResponseStruct{
			Status: QUERY_ARTICLE_CATE_FAILED,
			Msg:    Stat[QUERY_ARTICLE_CATE_FAILED],
		}, nil
	}

	if res == nil {
		Logger.Debugf("GetArticleCateList query empty")
		return &ArticleCateListResponseStruct{
			Status: QUERY_ARTICLE_CATE_EMPTY,
			Msg:    Stat[QUERY_ARTICLE_CATE_EMPTY],
		}, nil
	}

	var response ArticleCateListResponseStruct
	for _, v := range res {
		acds := new(ArticleCateDetailsStruct)
		acds.ID = v.ID
		acds.Name = v.Name
		acds.Keywords = v.Keywords
		acds.Description = v.Description
		acds.Pid = v.Pid
		acds.Status = v.Status
		acds.ImgURL = v.ImgURL
		acds.Sort = v.Sort
		acds.Addtime = v.Addtime
		response.ArticleCateList = append(response.ArticleCateList, acds)
	}

	response.Status = QUERY_ARTICLE_CATE_SUCCESS
	response.Msg = Stat[QUERY_ARTICLE_CATE_SUCCESS]
	Logger.Debugf("GetArticleCateList response:%v", response)
	return &response, nil
}

/**
 * [StartArticleCateServer 开启文章分类表服务]
 * @DateTime 2017-12-13T14:43:01+0800
 */
func StartArticleCateServer() {
	zkServers := zkclient.ZkServerAddress
	conn, err := zkclient.ConnectZk(zkServers)
	if err != nil {
		Logger.Fatalf("connect zk failed %v ", err)
	}
	defer conn.Close()

	port := "30069"
	ip, _ := zkclient.GetLocalIP()
	listenAddr := fmt.Sprintf("%s:%s", ip, port)

	servicename := "/cht/ArticleCateThriftService/providers"
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

	handler := &articlecateservice{}
	processor := NewArticleCateThriftServiceProcessor(handler)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	server.Serve()
}
