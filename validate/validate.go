/**
 * Created by Goland
 * @file   validate.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/13 08:22
 * @desc   validate.go
 */

package validate

import "regexp"

// CheckIsMobile 手机号码检测
func CheckIsMobile(mobileNum string) bool {
	var regular = "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// IsIdCard 判断是否是18或15位身份证
func IsIdCard(cardNo string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	if m, _ := regexp.MatchString(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`, cardNo); !m {
		return false
	}
	return true
}
