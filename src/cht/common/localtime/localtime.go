package localtime

import (
	"time"
)

/*获取当天0点时间*/
func GetLocalZeroTime() int64 {
	year, month, day := time.Now().Date()
	zero := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Unix()
	return zero
}

/*获取当天23:59:59点时间*/
func GetLocal24Time() int64 {
	year, month, day := time.Now().Date()
	end := time.Date(year, month, day, 23, 59, 59, 0, time.Local).Unix()
	return end
}
