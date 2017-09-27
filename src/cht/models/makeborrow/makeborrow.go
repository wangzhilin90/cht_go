package makeborrow

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MakeBorrowRequest struct {
	ID         int32  `orm:"column(id)`
	BorrowType int32  `orm:"column(borrow_type)`
	UserID     int32  `orm:"column(user_id)`
	Title      string `orm:"column(title)`
	Content    string `orm:"column(content)`
	Litpic     string `orm:"-;column(litpic)"`
	// Litpic     string `orm:"-"`
	BorrowUse  int32  `orm:"column(borrow_use)`
	IsDatetype int32  `orm:"column(is_datetype)`
	TimeLimit  int32  `orm:"column(time_limit)`
	Style      int32  `orm:"column(style)`
	Account    string `orm:"column(account)`
	// AccountTender   string `orm:"column(account_tender)`
	AccountTender   string `orm:"-;column(account_tender)"`
	Apr             string `orm:"column(apr)`
	AprAdd          string `orm:"column(apr_add)`
	MortgageFile    string `orm:"column(mortgage_file)`
	IsDxb           int32  `orm:"column(is_dxb)`
	Pwd             string `orm:"column(pwd);default("")`
	LowestAccount   string `orm:"column(lowest_account)`
	MostAccount     string `orm:"column(most_account)`
	ValidTime       int32  `orm:"column(valid_time)`
	Award           int32  `orm:"column(award)`
	Bonus           string `orm:"column(bonus)`
	IsFalse         int32  `orm:"column(is_false)`
	OpenAccount     int32  `orm:"column(open_account)`
	OpenBorrow      int32  `orm:"column(open_borrow)`
	OpenTender      int32  `orm:"column(open_tender)`
	OpenCredit      int32  `orm:"column(open_credit)`
	OpenZiliao      int32  `orm:"column(open_ziliao)`
	Material        int32  `orm:"column(material)`
	Addtime         int32  `orm:"column(addtime)`
	Addip           string `orm:"column(addip);`
	Status          int32  `orm:"column(status)`
	RutenAllnumber  int32  `orm:"column(ruten_allnumber)`
	RutenNumber     int32  `orm:"column(ruten_number)`
	VerifyUser      int32  `orm:"column(verify_user)`
	VerifyTime      int32  `orm:"column(verify_time)`
	VerifyRemark    string `orm:"column(verify_remark)`
	ReviewUser      int32  `orm:"column(review_user)`
	ReviewTimeLocal int32  `orm:"column(review_time_local)`
	ReviewTime      int32  `orm:"column(review_time)`
	// Secured         string `orm:"column(secured)"`
	Secured        string `"column(secured);"-""`
	Zhuanrangren   string `orm:"column(zhuanrangren)`
	Huodong        int32  `orm:"column(huodong)`
	SignDate       string `orm:"column(sign_date)`
	Subledger      int32  `orm:"column(subledger)`
	RepaySign      int32  `orm:"column(repay_sign)`
	AutoTenderLock int32  `orm:"column(auto_tender_lock)`
	IsAuto         int32  `orm:"column(is_auto)`
	IsCheck        int32  `orm:"column(is_check)`
	ReviewLock     int32  `orm:"column(review_lock)`
	FeeRate        string `orm:"column(fee_rate);`
	BorrowName     string `orm:"column(borrow_name)`
}

type Borrow struct {
	Id         int32  `orm:"column(id)`
	BorrowType int32  `orm:"column(borrow_type)`
	UserId     int32  `orm:"column(user_id)`
	Title      string `orm:"column(title)`
	Content    string `orm:"column(content)`
	// Litpic     string `orm:"column(litpic);"`
	Litpic          string `orm:"-"`
	BorrowUse       int32  `orm:"column(borrow_use)`
	IsDatetype      int32  `orm:"column(is_datetype)`
	TimeLimit       int32  `orm:"column(time_limit)`
	Style           int32  `orm:"column(style)`
	Account         string `orm:"column(account)`
	AccountTender   string `orm:"-"`
	Apr             string `orm:"column(apr)`
	AprAdd          string `orm:"column(apr_add)`
	MortgageFile    string `orm:"column(mortgage_file)`
	IsDxb           int32  `orm:"column(is_dxb)`
	Pwd             string `orm:"column(pwd);default("")`
	LowestAccount   string `orm:"column(lowest_account)`
	MostAccount     string `orm:"column(most_account)`
	ValidTime       int32  `orm:"column(valid_time)`
	Award           int32  `orm:"column(award)`
	Bonus           string `orm:"column(bonus)`
	IsFalse         int32  `orm:"column(is_false)`
	OpenAccount     int32  `orm:"column(open_account)`
	OpenBorrow      int32  `orm:"column(open_borrow)`
	OpenTender      int32  `orm:"column(open_tender)`
	OpenCredit      int32  `orm:"column(open_credit)`
	OpenZiliao      int32  `orm:"column(open_ziliao)`
	Material        int32  `orm:"column(material)`
	Addtime         int32  `orm:"column(addtime)`
	Addip           string `orm:"column(addip);`
	Status          int32  `orm:"column(status)`
	RutenAllnumber  int32  `orm:"column(ruten_allnumber)`
	RutenNumber     int32  `orm:"column(ruten_number)`
	VerifyUser      int32  `orm:"column(verify_user)`
	VerifyTime      int32  `orm:"column(verify_time)`
	VerifyRemark    string `orm:"column(verify_remark)`
	ReviewUser      int32  `orm:"column(review_user)`
	ReviewTimeLocal int32  `orm:"column(review_time_local)`
	ReviewTime      int32  `orm:"column(review_time)`
	// Secured         string `orm:"column(secured)"`
	Secured        string `"column(secured);"-""`
	Zhuanrangren   string `orm:"column(zhuanrangren)`
	Huodong        int32  `orm:"column(huodong)`
	SignDate       string `orm:"column(sign_date)`
	Subledger      int32  `orm:"column(subledger)`
	RepaySign      int32  `orm:"column(repay_sign)`
	AutoTenderLock int32  `orm:"column(auto_tender_lock)`
	IsAuto         int32  `orm:"column(is_auto)`
	IsCheck        int32  `orm:"column(is_check)`
	ReviewLock     int32  `orm:"column(review_lock)`
	FeeRate        string `orm:"-;size(64)"`
	BorrowName     string `orm:"column(borrow_name)`
}

func init() {
	orm.RegisterModelWithPrefix("jl_", new(Borrow))
}

/**
 * [CheckDepositAccount 是否开通徽商存管账户]
 * @param     user_id 用户ID
 * @return    bool ：是存管用户返回true 不是返回false
 * @DateTime 2017-09-01T18:09:08+0800
 */
func CheckDepositAccount(user_id int32) bool {
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
 * [InsertBorrowTbl 插入一条发标信息]
 * @param    *MakeBorrowRequest   发标入参
 * @DateTime 2017-09-02T10:30:36+0800
 */
func InsertBorrowTbl(mbr *MakeBorrowRequest) error {
	Logger.Debug("InsertBorrowTbl input param:", mbr)
	o := orm.NewOrm()
	o.Using("default")

	Logger.Debug("mbr.Secured:", mbr.Secured)
	sql := bytes.Buffer{}
	sql.WriteString("insert into jl_borrow ")
	sql.WriteString("(id,borrow_type,user_id,title,content,litpic,borrow_use,is_datetype,time_limit,")
	sql.WriteString("style,account,account_tender,apr,apr_add,mortgage_file,is_dxb,pwd,lowest_account,")
	sql.WriteString("most_account,valid_time,award,bonus,is_false,open_account,open_borrow,open_tender,")
	sql.WriteString("open_credit,open_ziliao,material,addtime,addip,status,ruten_allnumber,ruten_number,")
	sql.WriteString("verify_user,verify_time,verify_remark,review_user,review_time_local,review_time,secured,")
	sql.WriteString("zhuanrangren,huodong,sign_date,subledger,repay_sign,auto_tender_lock,")
	sql.WriteString("is_auto,is_check,review_lock,fee_rate,borrow_name) ")
	sql.WriteString(" values (next VALUE FOR MYCATSEQ_BORROW,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	Logger.Debug("InsertBorrowTbl sql:", sql.String())

	_, err := o.Raw(sql.String(),
		mbr.BorrowType,
		mbr.UserID,
		mbr.Title,
		mbr.Content,
		mbr.Litpic,
		mbr.BorrowUse,
		mbr.IsDatetype,
		mbr.TimeLimit,
		mbr.Style,
		mbr.Account,
		mbr.AccountTender,
		mbr.Apr,
		mbr.AprAdd,
		mbr.MortgageFile,
		mbr.IsDxb,
		mbr.Pwd,
		mbr.LowestAccount,
		mbr.MostAccount,
		mbr.ValidTime,
		mbr.Award,
		mbr.Bonus,
		mbr.IsFalse,
		mbr.OpenAccount,
		mbr.OpenBorrow,
		mbr.OpenTender,
		mbr.OpenCredit,
		mbr.OpenZiliao,
		mbr.Material,
		time.Now().Unix(),
		mbr.Addip,
		mbr.Status,
		mbr.RutenAllnumber,
		mbr.RutenNumber,
		mbr.VerifyUser,
		mbr.VerifyTime,
		mbr.VerifyRemark,
		mbr.ReviewUser,
		mbr.ReviewTimeLocal,
		mbr.ReviewTime,
		mbr.Secured,
		mbr.Zhuanrangren,
		mbr.Huodong,
		mbr.SignDate,
		mbr.Subledger,
		mbr.RepaySign,
		mbr.AutoTenderLock,
		mbr.IsAuto,
		mbr.IsCheck,
		mbr.ReviewLock,
		mbr.FeeRate,
		mbr.BorrowName).Exec()
	if err != nil {
		Logger.Error("InsertBorrowTbl insert failed", err)
		return err
	}
	return nil
}

func NewTMakeBorrowRequest(userID int32, borrowtype int32, borrowUse int32) *Borrow {
	return &Borrow{
		// Id:         next VALUE FOR MYCATSEQ_BORROR
		BorrowType: borrowtype,
		UserId:     userID,
		BorrowUse:  borrowUse,
		Title:      "biaoti",
		// Title:         ",",
		Content:       ",",
		Litpic:        "",
		TimeLimit:     1,
		Account:       "1000000.00",
		AccountTender: "0.00",
		Apr:           "0.0000",
		AprAdd:        "0.0000",
		MortgageFile:  ",",
		VerifyRemark:  ",",
		Pwd:           ",",
		LowestAccount: "50.00",
		MostAccount:   "0.00",
		ValidTime:     1,
		Bonus:         "0.00",
		OpenAccount:   1,
		OpenBorrow:    1,
		OpenTender:    1,
		OpenCredit:    1,
		OpenZiliao:    1,
		Addip:         ",",
		Secured:       "241234",
		Zhuanrangren:  ",",
		SignDate:      ",",
		FeeRate:       "20.00",
		BorrowName:    ",",
	}
}

func TInsertBorrowTbl() error {
	o := orm.NewOrm()
	o.Using("default")

	mbr := NewTMakeBorrowRequest(30, 5, 0)
	num, err := o.Insert(mbr)
	if err != nil {
		Logger.Errorf("TInsertBorrowTbl failed", err)
		return err
	}
	Logger.Debugf("TInsertBorrowTbl res %v", num)
	return nil
}
