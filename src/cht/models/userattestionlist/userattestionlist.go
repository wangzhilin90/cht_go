package userattestionlist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type UserAttestionListRequest struct {
	Username             string
	Realname             string
	RealStatus           int32
	EmailStatus          int32
	PhoneStatus          int32
	VideoStatus          int32
	SceneStatus          int32
	LimitOffset          int32
	LimitNum             int32
	ChengHuiTongTraceLog string
}

type UserAttestationDetails struct {
	UserID        int32  `orm:"column(user_id)"`
	CardType      int32  `orm:"column(card_type)"`
	HsCardType    string `orm:"column(hs_card_type)"`
	CardID        string `orm:"column(card_id)"`
	CardPic1      string `orm:"column(card_pic1)"`
	CardPic2      string `orm:"column(card_pic2)"`
	VideoPic      string `orm:"column(video_pic)"`
	RealStatus    int32  `orm:"column(real_status)"`
	RealPasstime  int32  `orm:"column(real_passtime)"`
	EmailStatus   int32  `orm:"column(email_status)"`
	EmailPasstime int32  `orm:"column(email_passtime)"`
	PhoneStatus   int32  `orm:"column(phone_status)"`
	PhonePasstime int32  `orm:"column(phone_passtime)"`
	VideoStatus   int32  `orm:"column(video_status)"`
	VideoPasstime int32  `orm:"column(video_passtime)"`
	SceneStatus   int32  `orm:"column(scene_status)"`
	ScenePasstime int32  `orm:"column(scene_passtime)"`
	VipStatus     int32  `orm:"column(vip_status)"`
	VipPasstime   int32  `orm:"column(vip_passtime)"`
	VipVerifytime int32  `orm:"column(vip_verifytime)"`
}

/**
 * [GetUserAttestionTatalNum 用户认证记录表总数]
 * @param    ualr *UserAttestionListRequest 请求入参
 * @return   int32 返回认证的总记录数
 * @DateTime 2017-10-19T15:01:14+0800
 */
func GetUserAttestionTatalNum(ualr *UserAttestionListRequest) (int32, error) {
	Logger.Debug("GetUserAttestionTatalNum input param:", ualr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("count(1)").
		From("jl_user_attestation UA").
		LeftJoin("jl_user U").On("UA.user_id=U.id").
		LeftJoin("jl_glossary G").On("UA.card_type=G.id").
		Where(fmt.Sprintf("1=1"))

	if strings.TrimSpace(ualr.Username) != "" {
		qb.And(fmt.Sprintf("U.username like \"%%%v%%\"", ualr.Username))
	}

	if strings.TrimSpace(ualr.Realname) != "" {
		qb.And(fmt.Sprintf("U.realname like \"%%%v%%\"", ualr.Realname))
	}

	if ualr.RealStatus != 0 && ualr.RealStatus != -1 {
		qb.And(fmt.Sprintf("UA.real_status=%d", ualr.RealStatus))
	}

	if ualr.EmailStatus != 0 && ualr.EmailStatus != -1 {
		qb.And(fmt.Sprintf("UA.email_status=%d", ualr.EmailStatus))
	}

	if ualr.PhoneStatus != 0 && ualr.PhoneStatus != -1 {
		qb.And(fmt.Sprintf("UA.phone_status=%d", ualr.PhoneStatus))
	}

	if ualr.VideoStatus != 0 && ualr.VideoStatus != -1 {
		qb.And(fmt.Sprintf("UA.video_status=%d", ualr.VideoStatus))
	}

	if ualr.SceneStatus != 0 && ualr.SceneStatus != -1 {
		qb.And(fmt.Sprintf("UA.scene_status=%d", ualr.SceneStatus))
	}

	sql := qb.String()
	Logger.Debug("GetUserAttestionTatalNum sql:", sql)
	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err != nil {
		Logger.Debugf("GetUserAttestionTatalNum query failed:", err)
		return 0, err
	}
	Logger.Debugf("GetUserAttestionTatalNum total num:%v", totalNum)
	return totalNum, nil
}

/**
 * [GetUserAttestionList 用户认证列表]
 * @param    ualr *UserAttestionListRequest 请求入参
 * @return   UserAttestationDetails 列表详情
 * @DateTime 2017-10-19T15:55:41+0800
 */
func GetUserAttestionList(ualr *UserAttestionListRequest) ([]UserAttestationDetails, error) {
	Logger.Debug("GetUserAttestionList input param:", ualr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("UA.*,U.username,U.realname,U.phone,U.email,G.name").
		From("jl_user_attestation UA").
		LeftJoin("jl_user U").On("UA.user_id=U.id").
		LeftJoin("jl_glossary G").On("UA.card_type=G.id").
		Where(fmt.Sprintf("1=1"))

	ualr.Username = strings.TrimSpace(ualr.Username)
	ualr.Realname = strings.TrimSpace(ualr.Realname)

	if ualr.Username != "" {
		qb.And(fmt.Sprintf("U.username like \"%%%v%%\"", ualr.Username))
	}

	if ualr.Realname != "" {
		qb.And(fmt.Sprintf("U.realname like \"%%%v%%\"", ualr.Realname))
	}

	if ualr.RealStatus != 0 && ualr.RealStatus != -1 {
		qb.And(fmt.Sprintf("UA.real_status=%d", ualr.RealStatus))
	}

	if ualr.EmailStatus != 0 && ualr.EmailStatus != -1 {
		qb.And(fmt.Sprintf("UA.email_status=%d", ualr.EmailStatus))
	}

	if ualr.PhoneStatus != 0 && ualr.PhoneStatus != -1 {
		qb.And(fmt.Sprintf("UA.phone_status=%d", ualr.PhoneStatus))
	}

	if ualr.VideoStatus != 0 && ualr.VideoStatus != -1 {
		qb.And(fmt.Sprintf("UA.video_status=%d", ualr.VideoStatus))
	}

	if ualr.SceneStatus != 0 && ualr.SceneStatus != -1 {
		qb.And(fmt.Sprintf("UA.scene_status=%d", ualr.SceneStatus))
	}

	if ualr.LimitNum != 0 {
		qb.Limit(int(ualr.LimitNum))
	}

	if ualr.LimitOffset != 0 {
		qb.Offset(int(ualr.LimitOffset))
	}

	sql := qb.String()
	Logger.Debug("GetUserAttestionList sql:", sql)
	var uad []UserAttestationDetails
	_, err := o.Raw(sql).QueryRows(&uad)
	if err != nil {
		Logger.Errorf("GetUserAttestionList query failed:", err)
		return nil, err
	}
	Logger.Debugf("GetUserAttestionList return value:%v", uad)
	return uad, nil
}
