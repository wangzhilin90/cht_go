package zkclient

import (
	cf "cht/common/config"
	. "cht/common/logger"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"net"
	"strings"
	"time"
)

var ZkServerAddress = func() []string {
	if cf.BConf.ZkAddress != nil {
		Logger.Debugf("zk server address:%v", cf.BConf.ZkAddress)
		return cf.BConf.ZkAddress
	} else {
		Logger.Fatalf("zk server address is null")
		return nil
	}
}()

//服务表 key值:服务名，对应永久路径 value:服务ip地址，可以有多个
type ServiceMap map[string]interface{}

func createPermanentNode(conn *zk.Conn, path string) error {
	_, err := conn.Create(path, []byte(""), 0, zk.WorldACL(zk.PermAll))
	return err
}

func createTemporaryNode(conn *zk.Conn, path string) error {
	_, err := conn.Create(path, []byte(""), zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return err
}

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		Logger.Errorf("getLocalIP failed", err)
		return "", err
	}

	var ip string
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return ip, nil
}

/**
 * [RegisterNode:注册节点]
 * @param    conn *zk.Conn [zk客户端句柄]
 * @param    path string   [服务名]
 * @param    ip            [服务ip地址]
 * @DateTime 2017-08-23T17:20:11+0800
 */
func RegisterNode(conn *zk.Conn, path string, listenAddr string) error {
	if strings.HasPrefix(path, "/") {
		substr := strings.Split(path, "/")
		var temp string
		for _, v := range substr {
			if v == "" {
				continue
			}
			v := fmt.Sprintf("/%s", v)
			temp = fmt.Sprintf("%s%s", temp, v)
			err := createPermanentNode(conn, temp)
			if err != zk.ErrNodeExists && err != nil {
				fmt.Println("CreatePermanentNode failed", err)
				return err
			}
		}
	}
	path = fmt.Sprintf("%s/%s", path, listenAddr)
	Logger.Debugf("service path", path)
	err := createTemporaryNode(conn, path)
	if err != zk.ErrNodeExists && err != nil {
		fmt.Println("CreateTemporaryNode failed", err)
		return err
	}
	return nil
}

/**
 * [Watch 监听临时节点变化]
 * @param    conn  *zk.Conn客户端句柄
 * @param    key   对应服务名,如/cht/service/provider
 * @param    value 对应ip地址
 * @DateTime 2017-12-25T14:43:40+0800
 */
func WatchNode(conn *zk.Conn, key, value string) error {
	path := fmt.Sprintf("%v/%v", key, value)
	Logger.Debugf("WatchNode listening:%v", path)
	_, _, ch, err := conn.ChildrenW(path)
	if err != nil {
		Logger.Errorf("ChildrenW failed:%v", err)
		return err
	}

	event := <-ch
	//临时节点被删除时，重新注册节点，再重新watch
	if event.Type == zk.EventNodeDeleted {
		err := RegisterNode(conn, key, value)
		if err != nil {
			Logger.Errorf("RegisterNode %v failed:%v", path, err)
			return err
		}
		Logger.Debugf("WatchNode %v again success!", path)
		WatchNode(conn, key, value)
	}
	return nil
}

/**
 * [DeleteNode 删除临时节点]
 * @param    conn       *zk.Conn          [zk客户端句柄]
 * @param    path       string            [服务名]
 * @param    servicemap 				  [服务表]
 * @DateTime 2017-08-23T19:44:31+0800
 */
func DeleteNode(conn *zk.Conn, path string, servicemap ServiceMap) error {
	fmt.Println("...")
	return nil
}

/**
 * [InitSevice 初始化服务表]
 * @param    conn       *zk.Conn    [zk客户端句柄]
 * @param    path       string      [服务名]
 * @param    servicemap ServiceMap	[服务表]
 * @DateTime 2017-08-23T17:27:19+0800
 */
func InitSevice(conn *zk.Conn, path string, servicemap ServiceMap) error {
	hosts, _, err := conn.Children(path)
	if err != nil {
		return err
	}
	servicemap[path] = hosts
	return nil
}

/**
 * [CallService 调用服务]
 * @param    conn *zk.Conn  	    [zk客户端句柄]
 * @param    path string			[服务名]
 * @param    servicemap ServiceMap  [服务表]
 * @DateTime 2017-08-23T17:31:27+0800
 */
func CallService(conn *zk.Conn, path string, servicemap ServiceMap) error {
	for {
		fmt.Println("call service ,", servicemap[path])
		time.Sleep(time.Second * 10)
	}
	return nil
}

/**
 * [ConnectZk description]
 * @param    {[type]}                 zkServers []string) (*zk.Conn [description]
 * @DateTime 2017-08-23T19:26:26+0800
 */
func ConnectZk(zkServers []string) (*zk.Conn, error) {
	conn, _, err := zk.Connect(zkServers, time.Second*10, zk.WithMaxBufferSize(1048576))
	if err != nil {
		fmt.Println("zkConnect faild ", err)
		return nil, err
	}
	return conn, nil
}
