package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"os"
	"path/filepath"
	// "github.com/liangguangchuan/gobbs/lib"
)

var (
	//基础配置文件
	BConf *Conf
	//项目访问路径
	AppPath string
	//运行模式 dev prod
	RunMode string
)

type MysqlConf struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"dbname"`
}

type RedisConf struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Expire_time int32  `json:"expire_time"`
}

type Conf struct {
	AppName   string    `json:"app_name"`
	RunMode   string    `json:"run_mode"`  //运行模式 分为debug、info、warn、error、fatal 5中模式
	Mysql     MysqlConf `json:"mysql"`     //mysql地址
	Redis     RedisConf `json:"redis"`     //redis地址
	ZkAddress []string  `json:"zkaddress"` //zk地址
	LogPath   string    `json:"logpath"`   //日志路径
}

var Config Conf

func init() {
	var err error
	BConf = NewConfig()
	//获取当前运行的 路径 如果获取失败抛出错误
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	//当前目录
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//拼接 conf 路径
	confPath := filepath.Join(workPath, "conf", "conf.json")
	//如果项目目录拼接conf/conf.xml 不存在对应文件
	if !FileExists(confPath) {
		confPath = filepath.Join(AppPath, "conf", "conf.json")
		if !FileExists(confPath) {
			//兼容测试用例跑的时候提示配置不存在
			confPath = filepath.Join(workPath, "../../conf", "conf.json")
			if !FileExists(confPath) {
				panic("please configure conf.json in current dir,such as `conf/conf.json`")
			}
		}
	}
	fmt.Println("confPath:", confPath)

	//读取文件并赋值 conf
	if err = parseConfig(confPath); err != nil {
		panic(err)
	}
}

/*初始化config,赋默认值*/
func NewConfig() *Conf {
	return &Conf{
		RunMode: "PROD",
		AppName: "cht",
		LogPath: "/var/cht/go_backend_service.log",
		Redis: RedisConf{
			//redis key过期时间60s
			Expire_time: 60,
		},
	}
}

//解析 conf.xml
func parseConfig(confPath string) error {
	fileData, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileData, BConf)
	if err != nil {
		fmt.Errorf("parseConfig error:", err)
		return err
	}

	fmt.Println("parseConfig config:", BConf)
	return err
}

//文件存在检测
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//是否为文件夹
func IsDir(path string) bool {
	if f, err := os.Stat(path); err == nil {
		if f.IsDir() {
			return true
		}
	}
	return false
}
