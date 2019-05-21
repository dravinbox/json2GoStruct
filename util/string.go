package util

import "strings"

/**
 * 字符串首字母转化为大写 user_name -> UserName
 */
func StrFirstToUpper(str string) string {
	//假如str=user_name, tmp = [user,name]
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		//遍历每段单词
		for i := 0; i < len(vv); i++ {
			if i == 0 { //单词的首字母

				if 97<=vv[i] && vv[i] <=122 { // 字符在a-z内的，转化为A-Z
					vv[i] -= 32
					upperStr += string(vv[i])
				}else if 65<=vv[i] && vv[i] <=90 {// 字符在A-Z内的，不转化
					upperStr += string(vv[i])

				}

			} else {
				upperStr += string(vv[i])
			}
		}
	}
	//return temp[0] + upperStr
	return upperStr
}
