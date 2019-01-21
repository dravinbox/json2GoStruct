package util

import "strings"

/**
 * 字符串首字母转化为大写 user_name -> UserName
 */
func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		//if y != 0 {
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i]) // + string(vv[i+1])
			} else {
				upperStr += string(vv[i])
			}
		}
		//}
	}
	//return temp[0] + upperStr
	return upperStr
}
