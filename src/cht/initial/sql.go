package initial

import (
	cf "cht/common/config"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	orm.Debug = true
	// user := "cht2_user205"
	// passwd := "cht2_user205"
	// host := "192.168.8.210"
	// port := "8066"
	// dbname := "cht2_1"

	// user := "cht"
	// passwd := "cht123456"
	// host := "192.168.10.2"
	// port := "3306"
	// dbname := "chtlocal"

	user := cf.BConf.Mysql.User
	passwd := cf.BConf.Mysql.Password
	host := cf.BConf.Mysql.Host
	port := cf.BConf.Mysql.Port
	dbname := cf.BConf.Mysql.DbName
	mysql_info := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, passwd, host, port, dbname)
	fmt.Println("init mysql info:", mysql_info)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysql_info)
}
