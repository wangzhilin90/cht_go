package borrower

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "time"
)

type BorrowerInfoRequest struct {
	Name                 string
	ChengHuiTongTraceLog string
}

type MaterialInfoStruct struct {
	ID   int32  `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

type BorrowerInfoStruct struct {
	ID           int32  `orm:"column(id)"`
	Realname     string `orm:"column(realname)"`
	IsBorrower   int32  `orm:"column(is_borrower)"`
	CardID       string `orm:"column(card_id)"`
	Credit       string `orm:"column(credit)"`
	Guarantor    string `orm:"column(guarantor)"`
	MaterialList []*MaterialInfoStruct
}

type BorrowerUID struct {
	ID         int32  `orm:"column(id)"`
	Realname   string `orm:"column(realname)"`
	IsBorrower int32  `orm:"column(is_borrower)"`
}

/**
 * [getBorrowerUID 根据username得到用户ID]
 * @param    {[type]}                 birs *BorrowerInfoRequest) (*BorrowerUID, error [description]
 * @return   {[type]}                      [description]
 * @DateTime 2017-09-13T16:30:19+0800
 */
func GetBorrowerUID(birs *BorrowerInfoRequest) (*BorrowerUID, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetBorrowerUID input param:", birs)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id,realname,is_borrower").
		From("jl_user").
		Where(fmt.Sprintf("username=\"%s\"", birs.Name)).
		Limit(1)

	sql := qb.String()
	Logger.Debug("GetBorrowerUID sql:", sql)
	var bu BorrowerUID
	err := o.Raw(sql).QueryRow(&bu)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Error("GetBorrowerUID query failed:", err)
		return nil, err
	}
	Logger.Debugf("GetBorrowerUID res :%v", bu)
	return &bu, nil
}

/**
 * [getCardID 根据用户ID得到身份证信息，得到为空不返回错误信息]
 * @param    user_id 用户ID
 * @return   string 用户身份证号
 * @DateTime 2017-09-13T16:20:24+0800
 */
func GetCardID(user_id int32) (string, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("getCardID input param:", user_id)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("UA.card_id").
		From("jl_user_attestation UA").
		LeftJoin("jl_user U").On("UA.user_id=U.id").
		LeftJoin("jl_glossary G").On("UA.card_type=G.id").
		Where(fmt.Sprintf("UA.user_id=%d", user_id)).
		Limit(1)

	sql := qb.String()
	Logger.Debug("getCardID sql:", sql)
	var card_id string
	o.Raw(sql).QueryRow(&card_id)
	Logger.Debugf("getCardID res :%v", card_id)
	return card_id, nil
}

/**
 * [getCreditUse 得到用户信用额度]
 * @param    user_id 用户ID
 * @return   string  用户信用额度
 * @DateTime 2017-09-13T16:34:41+0800
 */
func GetCreditUse(user_id int32) (string, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("getCardID input param:", user_id)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("credit_use").
		From("jl_user_credit").
		Where(fmt.Sprintf("user_id=%d", user_id)).
		Limit(1)

	sql := qb.String()
	Logger.Debug("getCreditUse sql:", sql)
	var credit_use string
	o.Raw(sql).QueryRow(&credit_use)
	if credit_use == "" {
		return "0", nil
	}
	Logger.Debugf("getCardID res :%v", credit_use)
	return credit_use, nil
}

/**
 * [GetReviewAccount 查询用户发标待审金额]
 * @param    user_id 用户ID
 * @return    string 用户发标待审金额 error 查询是否出错
 * @DateTime 2017-12-04T17:20:11+0800
 */
func GetReviewAccount(user_id int32) (string, error) {
	Logger.Debugf("GetReviewAccount input param: %v", user_id)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT SUM(account)  FROM jl_borrow WHERE user_id = ? AND status = 0 LIMIT 1")
	sql := buf.String()

	Logger.Debugf("GetReviewAccount sql: %v", sql)

	var review_account string
	err := o.Raw(sql, user_id).QueryRow(&review_account)
	if err != nil {
		Logger.Errorf("GetReviewAccount query failed %v", err)
		return "0", err
	}

	if review_account == "" {
		return "0", nil
	}
	Logger.Debugf("GetReviewAccount return value:%v", review_account)
	return review_account, nil
}

/**
 * [getGuarantor 获取担保人]
 * @param    user_id 用户ID
 * @return   string 担保人
 * @DateTime 2017-09-13T16:40:44+0800
 */
func GetGuarantor(user_id int32) (string, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("getGuarantor input param:", user_id)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("UF.guarantor").
		From("jl_user_field UF").
		LeftJoin("jl_user U").On("UF.user_id=U.id").
		Where(fmt.Sprintf("UF.user_id=%d", user_id)).
		Limit(1)

	sql := qb.String()
	Logger.Debug("getGuarantor sql:", sql)
	var guarantor string
	o.Raw(sql).QueryRow(&guarantor)
	Logger.Debugf("getGuarantor res :%v", guarantor)
	return guarantor, nil
}

func GetMaterialInfo(user_id int32) ([]MaterialInfoStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("getMaterialInfo input param:", user_id)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id,name").
		From("jl_material_class").
		Where(fmt.Sprintf("user_id=%d", user_id))

	sql := qb.String()
	Logger.Debug("getMaterialInfo sql:", sql)
	var materialInfo []MaterialInfoStruct
	o.Raw(sql).QueryRows(&materialInfo)
	Logger.Debugf("getMaterialInfo res :%v", materialInfo)
	return materialInfo, nil
}
