package makeborrow

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

func init() {
	orm.RegisterModelWithPrefix("jl_", new(Borrow))
}

type Borrow struct {
	ID              int32  `orm:"column(id)"`
	BorrowType      int32  `orm:"column(borrow_type);-"`
	UserID          int32  `orm:"column(user_id);-"`
	Title           string `orm:"column(title);-`
	Content         string `orm:"column(content);-`
	Litpic          string `orm:"column(litpic);-"`
	BorrowUse       int32  `orm:"column(borrow_use);-"`
	IsDatetype      int32  `orm:"column(is_datetype);-"`
	TimeLimit       int32  `orm:"column(time_limit);-"`
	Style           int32  `orm:"column(style);-"`
	Account         string `orm:"column(account);-"`
	AccountTender   string `orm:"column(account_tender);-"`
	Apr             string `orm:"column(apr);-"`
	AprAdd          string `orm:"column(apr_add);-"`
	MortgageFile    string `orm:"column(mortgage_file);-"`
	IsDxb           int32  `orm:"column(is_dxb);-"`
	Pwd             string `orm:"column(pwd);-"`
	LowestAccount   string `orm:"column(lowest_account);-"`
	MostAccount     string `orm:"column(most_account);-"`
	ValidTime       int32  `orm:"column(valid_time);-"`
	Award           int32  `orm:"column(award);-"`
	Bonus           string `orm:"column(bonus);-"`
	IsFalse         int32  `orm:"column(is_false);-"`
	OpenAccount     int32  `orm:"column(open_account);-"`
	OpenBorrow      int32  `orm:"column(open_borrow);-"`
	OpenTender      int32  `orm:"column(open_tender);-"`
	OpenCredit      int32  `orm:"column(open_credit);-"`
	OpenZiliao      int32  `orm:"column(open_ziliao);-"`
	Material        int32  `orm:"column(material);-"`
	Addtime         int32  `orm:"column(addtime);-"`
	Addip           string `orm:"column(addip);-"`
	Status          int32  `orm:"column(status);-"`
	RutenAllnumber  int32  `orm:"column(ruten_allnumber);-"`
	RutenNumber     int32  `orm:"column(ruten_number);-"`
	VerifyUser      int32  `orm:"column(verify_user);-"`
	VerifyTime      int32  `orm:"column(verify_time);-"`
	VerifyRemark    string `orm:"column(verify_remark)"`
	ReviewUser      int32  `orm:"column(review_user);-"`
	ReviewTimeLocal int32  `orm:"column(review_time_local);-"`
	ReviewTime      int32  `orm:"column(review_time);-"`
	Secured         string `orm:"column(secured);-"`
	Zhuanrangren    string `orm:"column(zhuanrangren);-"`
	Huodong         int32  `orm:"column(huodong);-"`
	SignDate        string `orm:"column(sign_date);-"`
	Subledger       int32  `orm:"column(subledger);-`
	RepaySign       int32  `orm:"column(repay_sign);-"`
	AutoTenderLock  int32  `orm:"column(auto_tender_lock);-"`
	IsAuto          int32  `orm:"column(is_auto);-"`
	IsCheck         int32  `orm:"column(is_check);-"`
	ReviewLock      int32  `orm:"column(review_lock);-"`
	FeeRate         string `orm:"column(fee_rate);-"`
	BorrowName      string `orm:"column(borrow_name);-"`
}

/**
 * [CheckDepositAccount 是否开通徽商存管账户]
 * @param     user_id 用户ID
 * @return    bool ：是存管用户返回true 不是返回false
 * @DateTime 2017-09-01T18:09:08+0800
 */
func CheckDepositAccount(user_id int32) bool {
	Logger.Debugf("CheckDepositAccount input param:%v", user_id)
	o := orm.NewOrm()
	o.Using("default")

	var hsid string
	o.Raw("SELECT hsid FROM jl_user WHERE id=?", user_id).QueryRow(&hsid)
	Logger.Debugf("CheckDepositAccount res", hsid)
	if hsid != "" {
		return true
	}
	return false
}

/**
 * [GetGuarantor 获取担保人]
 * @param    user_id 用户ID
 * @return   string 没有获取到为""
 * @DateTime 2017-09-02T11:45:34+0800
 */
func GetGuarantor(user_id int32) string {
	Logger.Debugf("GetGuarantor input param:%v", user_id)
	o := orm.NewOrm()
	o.Using("default")

	var guarantor string
	err := o.Raw("SELECT guarantor FROM jl_user_field WHERE user_id=?", user_id).QueryRow(&guarantor)
	if err != nil {
		Logger.Error("GetGuarantor query faied")
		return ""
	}
	Logger.Debugf("GetGuarantor res", guarantor)
	return guarantor
}

/**
 * [GetCreditLimit 查询授信额度]
 * @param    user_id   用户ID
 * @return   int32 	   返回授信额度 error查询是否出错
 * @DateTime 2017-09-02T10:29:56+0800
 */
func GetCreditLimit(user_id int32) (string, error) {
	Logger.Debugf("GetCreditLimit input param:%v", user_id)
	o := orm.NewOrm()
	o.Using("default")

	var credit_use string
	err := o.Raw("SELECT credit_use FROM jl_user_credit WHERE user_id=?", user_id).QueryRow(&credit_use)
	if err != nil {
		Logger.Error("GetCreditLimit query faied")
		return "", err
	}
	Logger.Debugf("GetCreditLimit res %v", credit_use)
	return credit_use, nil
}

/**
 * [GetLatestID 得到全局序列号最新值]
 * @param    mbr *MakeBorrowRequest 请求入参
 * @return   int32 最新的标ID
 * @DateTime 2017-11-06T10:19:10+0800
 */
func GetLatestBorrowID() (int32, error) {
	o := orm.NewOrm()
	o.Using("default")

	var ID int32
	err := o.Raw("SELECT next VALUE FOR MYCATSEQ_BORROW").QueryRow(&ID)
	if err != nil {
		Logger.Error("GetLatestID query faied")
		return 0, err
	}
	Logger.Debugf("GetLatestID res %v", ID)
	return ID, nil
}

/**
 * [InsertBorrowTbl 插入一条发标信息]
 * @param    bs *Borrow 发标入参
 * @return   int32返回最新插入的ID，error返回错误信息
 * @DateTime 2017-11-06T12:50:18+0800
 */
func InsertBorrowTbl(bs *Borrow) (int32, error) {
	Logger.Debug("InsertBorrowTbl input param:", bs)
	o := orm.NewOrm()
	o.Using("default")
	o.Begin()

	num, err := o.Insert(bs)
	if err != nil {
		Logger.Errorf("InsertBorrowTbl failed", err)
		return 0, err
	}
	last_insert_num := int32(num)
	Logger.Debugf("InsertBorrowTbl last input num %v", last_insert_num)

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_borrow")

	var str string
	if bs.ID == 0 {
		Err := fmt.Errorf("InsertBorrowTbl input param failed")
		o.Rollback()
		return 0, Err
	}

	if bs.Title != "" {
		str += fmt.Sprintf("title=\"%s\",", bs.Title)
	}

	if bs.Content != "" {
		str += fmt.Sprintf("content=\"%s\",", bs.Content)
	}

	if bs.Litpic != "" {
		str += fmt.Sprintf("litpic=\"%s\",", bs.Litpic)
	}

	if bs.Account != "" {
		str += fmt.Sprintf("account=\"%s\",", bs.Account)
	}

	if bs.AccountTender != "" {
		str += fmt.Sprintf("account_tender=\"%s\",", bs.AccountTender)
	}

	if bs.Apr != "" {
		str += fmt.Sprintf("apr=\"%s\",", bs.Apr)
	}

	if bs.AprAdd != "" {
		str += fmt.Sprintf("apr_add=\"%s\",", bs.AprAdd)
	}

	if bs.MortgageFile != "" {
		str += fmt.Sprintf("mortgage_file=\"%s\",", bs.MortgageFile)
	}

	if bs.Pwd != "" {
		str += fmt.Sprintf("pwd=\"%s\",", bs.Pwd)
	}

	if bs.LowestAccount != "" {
		str += fmt.Sprintf("lowest_account=\"%s\",", bs.LowestAccount)
	}

	if bs.MostAccount != "" {
		str += fmt.Sprintf("most_account=\"%s\",", bs.MostAccount)
	}

	if bs.Bonus != "" {
		str += fmt.Sprintf("bonus=\"%s\",", bs.Bonus)
	}

	if bs.Addip != "" {
		str += fmt.Sprintf("addip=\"%s\",", bs.Addip)
	}

	if bs.VerifyRemark != "" {
		str += fmt.Sprintf("verify_remark=\"%s\",", bs.VerifyRemark)
	}

	if bs.Secured != "" {
		str += fmt.Sprintf("secured=\"%s\",", bs.Secured)
	}

	if bs.Zhuanrangren != "" {
		str += fmt.Sprintf("zhuanrangren=\"%s\",", bs.Zhuanrangren)
	}

	if bs.SignDate != "" {
		str += fmt.Sprintf("sign_date=\"%s\",", bs.SignDate)
	}

	if bs.FeeRate != "" {
		str += fmt.Sprintf("fee_rate=\"%s\",", bs.FeeRate)
	}

	if bs.BorrowName != "" {
		str += fmt.Sprintf("borrow_name=\"%s\",", bs.BorrowName)
	}

	if bs.BorrowType != 0 {
		str += fmt.Sprintf("borrow_type=%d,", bs.BorrowType)
	}

	if bs.UserID != 0 {
		str += fmt.Sprintf("user_id=%d,", bs.UserID)
	}

	if bs.BorrowUse != 0 {
		str += fmt.Sprintf("borrow_use=%d,", bs.BorrowUse)
	}

	if bs.IsDatetype != 0 {
		str += fmt.Sprintf("is_datetype=%d,", bs.IsDatetype)
	}

	if bs.TimeLimit != 0 {
		str += fmt.Sprintf("time_limit=%d,", bs.TimeLimit)
	}

	if bs.Style != 0 {
		str += fmt.Sprintf("style=%d,", bs.Style)
	}

	if bs.IsDxb != 0 {
		str += fmt.Sprintf("is_dxb=%d,", bs.IsDxb)
	}

	if bs.ValidTime != 0 {
		str += fmt.Sprintf("valid_time=%d,", bs.ValidTime)
	}

	if bs.Award != 0 {
		str += fmt.Sprintf("award=%d,", bs.Award)
	}

	if bs.IsFalse != 0 {
		str += fmt.Sprintf("is_false=%d,", bs.IsFalse)
	}

	if bs.OpenAccount != 0 {
		str += fmt.Sprintf("open_account=%d,", bs.OpenAccount)
	}

	if bs.OpenBorrow != 0 {
		str += fmt.Sprintf("open_borrow=%d,", bs.OpenBorrow)
	}

	if bs.OpenTender != 0 {
		str += fmt.Sprintf("open_tender=%d,", bs.OpenTender)
	}

	if bs.OpenCredit != 0 {
		str += fmt.Sprintf("open_credit=%d,", bs.OpenCredit)
	}

	if bs.OpenZiliao != 0 {
		str += fmt.Sprintf("open_ziliao=%d,", bs.OpenZiliao)
	}

	if bs.Material != 0 {
		str += fmt.Sprintf("material=%d,", bs.Material)
	}

	str += fmt.Sprintf("addtime=%d,", time.Now().Unix())

	if bs.Status != 0 {
		str += fmt.Sprintf("status=%d,", bs.Status)
	}

	if bs.RutenAllnumber != 0 {
		str += fmt.Sprintf("ruten_allnumber=%d,", bs.RutenAllnumber)
	}

	if bs.RutenNumber != 0 {
		str += fmt.Sprintf("ruten_number=%d,", bs.RutenNumber)
	}

	if bs.VerifyUser != 0 {
		str += fmt.Sprintf("verify_user=%d,", bs.VerifyUser)
	}

	if bs.VerifyTime != 0 {
		str += fmt.Sprintf("verify_time=%d,", bs.VerifyTime)
	}

	if bs.ReviewUser != 0 {
		str += fmt.Sprintf("review_user=%d,", bs.ReviewUser)
	}

	if bs.ReviewTimeLocal != 0 {
		str += fmt.Sprintf("review_time_local=%d,", bs.ReviewTimeLocal)
	}

	if bs.ReviewTime != 0 {
		str += fmt.Sprintf("review_time=%d,", bs.ReviewTime)
	}

	if bs.Huodong != 0 {
		str += fmt.Sprintf("huodong=%d,", bs.Huodong)
	}

	if bs.Subledger != 0 {
		str += fmt.Sprintf("subledger=%d,", bs.Subledger)
	}

	if bs.RepaySign != 0 {
		str += fmt.Sprintf("repay_sign=%d,", bs.RepaySign)
	}

	if bs.AutoTenderLock != 0 {
		str += fmt.Sprintf("auto_tender_lock=%d,", bs.AutoTenderLock)
	}

	if bs.IsAuto != 0 {
		str += fmt.Sprintf("is_auto=%d,", bs.IsAuto)
	}

	if bs.IsCheck != 0 {
		str += fmt.Sprintf("is_check=%d,", bs.IsCheck)
	}

	if bs.ReviewLock != 0 {
		str += fmt.Sprintf("review_lock=%d,", bs.ReviewLock)
	}

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("id=%d", last_insert_num))
	sql := qb.String()

	Logger.Debug("InsertBorrowTbl sql:", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		o.Rollback()
		Logger.Errorf("InsertBorrowTbl update failed :%v", err)
		return last_insert_num, err
	}
	affectnum, _ := res.RowsAffected()
	Logger.Debugf("InsertBorrowTbl change num :%v", affectnum)
	if affectnum == 0 {
		o.Rollback()
		Err := fmt.Errorf("InsertBorrowTbl failed")
		return 0, Err
	}
	o.Commit()
	return last_insert_num, nil
}
