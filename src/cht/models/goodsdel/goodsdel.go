package goodsdel

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type GoodsDeLRequest struct {
	ID                   int32
	ChengHuiTongTraceLog string
}

/**
 * [DelGoods 商品管理---删除商品]
 * @param    gdr *GoodsDeLRequest 请求入参
 * @return   bool true:删除成功 false:删除失败
 * @DateTime 2017-10-23T16:57:20+0800
 */
func DelGoods(gdr *GoodsDeLRequest) bool {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("DelGoods input param:", gdr)

	buf := bytes.Buffer{}
	buf.WriteString("DELETE FROM jl_point_goods WHERE id=? AND sold_num=0")
	sql := buf.String()
	Logger.Debugf("DelGoods sql %v", sql)

	res, err := o.Raw(sql, gdr.ID).Exec()
	if err != nil {
		Logger.Errorf("DelGoods delete failed :%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		Logger.Errorf("DelGoods delete failed :%v", err)
		return false
	}
	return true
}
