package initial

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	orm.Debug = true
	user := "cht2_user205"
	passwd := "cht2_user205"
	host := "192.168.8.210"
	port := "8066"
	dbname := "cht2_1"
	// user := "cht"
	// passwd := "cht123456"
	// host := "192.168.10.2"
	// port := "3306"
	// dbname := "chtlocal"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, passwd, host, port, dbname))
}
