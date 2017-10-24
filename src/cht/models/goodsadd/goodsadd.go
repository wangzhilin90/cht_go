package goodsadd

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type GoodsAddRequest struct {
	ShowTime             int32
	CloseTime            int32
	IsTimer              int32
	Litpic               string
	Name                 string
	Category             int32
	RedbagMoney          string
	OriginalPoint        int32
	CurrentPoint         int32
	TotalNum             int32
	SingleNum            int32
	Content              string
	ChengHuiTongTraceLog string
}

/**
 * [AddGoods 商品管理---添加商品]
 * @param    gar *GoodsAddRequest 请求入参
 * @return   bool true:添加成功 false:添加失败
 * @DateTime 2017-10-23T14:56:16+0800
 */
func AddGoods(gar *GoodsAddRequest) bool {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("AddGoods input param:", gar)

	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_point_goods (id,addtime,show_time, ")
	buf.WriteString("close_time,is_timer,litpic,name, category,redbag_money, ")
	buf.WriteString("original_point,current_point,total_num,single_num,content) ")
	buf.WriteString("values(next VALUE FOR MYCATSEQ_POINT_GOODS,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	sql := buf.String()
	Logger.Debugf("AddGoods sql %v", sql)

	res, err := o.Raw(sql,
		time.Now().Unix(),
		gar.ShowTime,
		gar.CloseTime,
		gar.IsTimer,
		gar.Litpic,
		gar.Name,
		gar.Category,
		gar.RedbagMoney,
		gar.OriginalPoint,
		gar.CurrentPoint,
		gar.TotalNum,
		gar.SingleNum,
		gar.Content,
	).Exec()
	if err != nil {
		Logger.Errorf("AddGoods insert failed :%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	return true
}
