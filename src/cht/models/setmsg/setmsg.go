package setmsg

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type SetMsgDetailsRequest struct {
	UserID               int32
	ChengHuiTongTraceLog string
}

type SetMsgDetailsStruct struct {
	ID            int32 `orm:"column(id)"`
	UserID        int32 `orm:"column(user_id)"`
	Addtime       int32 `orm:"column(addtime)"`
	Status        int32 `orm:"column(status)"`
	TenderStatus  int32 `orm:"column(tender_status)"`
	BorrowStatus  int32 `orm:"column(borrow_status)"`
	ProtoStatus   int32 `orm:"column(proto_status)"`
	StationStatus int32 `orm:"column(station_status)"`
	GuideStatus   int32 `orm:"column(guide_status)"`
	SoundStatus   int32 `orm:"column(sound_status)"`
}

type SetMsgDealRequest struct {
	UserID               int32
	Addtime              int32
	Status               int32
	TenderStatus         int32
	BorrowStatus         int32
	ProtoStatus          int32
	StationStatus        int32
	GuideStatus          int32
	SoundStatus          int32
	ChengHuiTongTraceLog string
}

func GetSetMsgDetails(smdr *SetMsgDetailsRequest) (*SetMsgDetailsStruct, error) {
	Logger.Debugf("GetSetMsgDetails input param:%v", smdr)
	o := orm.NewOrm()
	o.Using("default")

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_set_msg").
		Where(fmt.Sprintf("user_id=%d", smdr.UserID))
	sql := qb.String()
	Logger.Debugf("GetSetMsgDetails sql:%v", sql)

	var sds SetMsgDetailsStruct
	err := o.Raw(sql).QueryRow(&sds)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetSetMsgDetails query failed:%v", sds)
		return nil, err
	}

	Logger.Debugf("GetSetMsgDetails return value:%v", sds)
	return &sds, nil
}

func UpdateSetMsgDetails(smdr *SetMsgDealRequest) bool {
	Logger.Debugf("UpdateSetMsgDetails input param:%v", smdr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_set_msg")

	var str string
	if smdr.Addtime != 0 {
		str += fmt.Sprintf("addtime=%d,", smdr.Addtime)
	}

	if smdr.Status != 0 {
		str += fmt.Sprintf("status=%d,", smdr.Status)
	}

	if smdr.TenderStatus != 0 {
		str += fmt.Sprintf("tender_status=%d,", smdr.TenderStatus)
	}

	if smdr.BorrowStatus != 0 {
		str += fmt.Sprintf("borrow_status=%d,", smdr.BorrowStatus)
	}

	if smdr.ProtoStatus != 0 {
		str += fmt.Sprintf("proto_status=%d,", smdr.ProtoStatus)
	}

	if smdr.StationStatus != 0 {
		str += fmt.Sprintf("station_status=%d,", smdr.StationStatus)
	}

	if smdr.GuideStatus != 0 {
		str += fmt.Sprintf("guide_status=%d,", smdr.GuideStatus)
	}

	if smdr.SoundStatus != 0 {
		str += fmt.Sprintf("sound_status=%d,", smdr.SoundStatus)
	}
	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("user_id=%d", smdr.UserID))
	sql := qb.String()
	Logger.Debugf("UpdateSetMsgDetails sql:%v", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Errorf("UpdateSetMsgDetails update failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	return true
}

func InsertSetMsgDetails(smdr *SetMsgDealRequest) bool {
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("INSERT INTO jl_set_msg (id,user_id,addtime,status,tender_status,borrow_status,proto_status,station_status,guide_status,sound_status) ")
	buf.WriteString("VALUES (next VALUE FOR MYCATSEQ_SET_MSG,?,?,?,?,?,?,?,?,?)")
	sql := buf.String()

	Logger.Debugf("InsertSetMsgDetails sql: %v", sql)

	res, err := o.Raw(sql,
		smdr.UserID,
		time.Now().Unix(),
		smdr.Status,
		smdr.TenderStatus,
		smdr.BorrowStatus,
		smdr.ProtoStatus,
		smdr.StationStatus,
		smdr.GuideStatus,
		smdr.SoundStatus,
	).Exec()
	if err != nil {
		Logger.Errorf("InsertSetMsgDetails insert failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	return true
}
