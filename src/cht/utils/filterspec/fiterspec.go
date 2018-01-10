package filterspec

import (
	"reflect"
	"regexp"
)

/**
 * [FiterSpecialCharacters 过滤特殊字符,防止sql注入]
 * @param    input interface{} 请求结构体指针
 * @return   interface{} 	   返回结构体指针
 * @DateTime 2018-01-10T10:19:55+0800
 */
func FiterSpecialCharacters(input interface{}) interface{} {
	str := `select|Update|and|or|delete|insert|trancate| \
			char|into|substr|ascii|declare|exec|count|master|into| \
			drop|execute|\"|%|;|\(|\)|&|\+`
	var re, _ = regexp.Compile(str)
	t := reflect.ValueOf(input).Elem()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		switch f.Kind() {
		case reflect.String:
			if f.CanInterface() {
				if str, ok := f.Interface().(string); ok {
					new1 := re.ReplaceAllString(str, "")
					f.SetString(new1)
				}
			}
		}
	}

	return input
}
