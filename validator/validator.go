/**
 * Created by Goland
 * @file   validate.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/6/13 08:22
 * @desc   validate.go
 */

package validator

import (
	"encoding/json"
	"github.com/x-module/helper/function"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// CheckIsMobile 手机号码检测
func CheckIsMobile(mobileNum string) bool {
	var regular = "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// IsIdCard 判断是否是18或15位身份证
func IsIdCard(cardNo string) bool {
	// 18位身份证 ^(\d{17})([0-9]|X)$
	if m, _ := regexp.MatchString(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`, cardNo); !m {
		return false
	}
	return true
}

// IsURL 是否为URL地址
func IsURL(url string) bool {
	match, _ := regexp.MatchString(`^(http|https)://[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?$`, url)
	return match
}

// IsQQ 验证是否为QQ号
func IsQQ(qq string) bool {
	match, _ := regexp.MatchString(`^[1-9][0-9]{4,12}$`, qq)
	return match
}

// IsWeChat 验证是否为微信号
func IsWeChat(wechat string) bool {
	match, _ := regexp.MatchString(`^[a-zA-Z][-_a-zA-Z0-9]{6,20}$`, wechat)
	return match
}

// IsWeibo 验证是否为微博ID
func IsWeibo(weibo string) bool {
	if len(weibo) < 6 || len(weibo) > 20 {
		return false
	}

	if matched, _ := regexp.MatchString(`^[a-zA-Z][\w-]*$`, weibo); !matched {
		return false
	}

	return true
}

// IsPassword 验证密码是否合法
// 密码长度在6-20个字符之间，必须包含数字、字母和特殊符号
func IsPassword(password string) bool {
	if len(password) < 6 || len(password) > 20 {
		return false
	}

	if matched, _ := regexp.MatchString(`[a-zA-Z]`, password); !matched {
		return false
	}

	if matched, _ := regexp.MatchString(`\d`, password); !matched {
		return false
	}

	if matched, _ := regexp.MatchString(`[^a-zA-Z\d]`, password); !matched {
		return false
	}

	return true
}

// IsBankCardNo 验证是否为大陆银行卡号
func IsBankCardNo(cardNumber string) bool {
	if len(cardNumber) != 16 && len(cardNumber) != 19 {
		return false
	}
	var cardArr []int
	for _, c := range cardNumber {
		if c < '0' || c > '9' {
			return false
		}
		cardArr = append(cardArr, int(c-'0'))
	}
	if len(cardArr) == 16 {
		sum := 0
		for i := len(cardArr) - 1; i >= 0; i-- {
			if i%2 == 0 {
				cardArr[i] *= 2
				if cardArr[i] > 9 {
					cardArr[i] -= 9
				}
			}
			sum += cardArr[i]
		}
		return sum%10 == 0
	} else {
		sum := 0
		for i := len(cardArr) - 1; i >= 0; i-- {
			if (len(cardArr)-i)%2 == 0 {
				cardArr[i] *= 2
				if cardArr[i] > 9 {
					cardArr[i] -= 9
				}
			}
			sum += cardArr[i]
		}
		return sum%10 == 0
	}
}

// IsTime 验证是否为时间格式（HH:mm:ss）
func IsTime(str string) bool {
	reg := regexp.MustCompile(`^(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d$`)
	return reg.MatchString(str)
}

// IsDate 验证是否为日期格式（yyyy-MM-dd）
func IsDate(str string) bool {
	reg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !reg.MatchString(str) {
		return false
	}
	parts := strings.Split(str, "-")
	year, _ := strconv.Atoi(parts[0])
	month, _ := strconv.Atoi(parts[1])
	day, _ := strconv.Atoi(parts[2])
	return isValidMonth(month, year) && isValidDay(day, month, year)
}

func isValidMonth(month, year int) bool {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return true
	case 4, 6, 9, 11:
		return true
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			return true
		}
		return false
	}
	return false
}

func isValidDay(day, month, year int) bool {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day >= 1 && day <= 31 {
			return true
		}
	case 4, 6, 9, 11:
		if day >= 1 && day <= 30 {
			return true
		}
	case 2:
		if year%4 == 0 && year%100 != 0 || year%400 == 0 {
			if day >= 1 && day <= 29 {
				return true
			}
		} else {
			if day >= 1 && day <= 28 {
				return true
			}
		}
	}
	return false
}

// IsDateTime 验证是否为日期时间格式（yyyy-MM-dd HH:mm:ss）
func IsDateTime(str string) bool {
	reg := regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`)
	if !reg.MatchString(str) {
		return false
	}
	if !IsDate(str[0:10]) || !IsTime(str[11:]) {
		return false
	}
	return true
}

// IsIDCard 验证身份证号(18或15位)
func IsIDCard(str string) bool {
	if len(str) != 15 && len(str) != 18 {
		return false
	}
	if len(str) == 18 {
		return IsIDCard18(str)
	} else {
		return IsIDCard15(str)
	}
}

// IsIDCard18 验证18位身份证号
func IsIDCard18(id string) bool {
	// 18位身份证号码正则表达式，根据规则来编写
	regExp := "^[1-9]\\d{5}(19|20)\\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$"
	// 利用正则表达式匹配身份证号码
	match, err := regexp.MatchString(regExp, id)
	if err != nil {
		// 匹配过程出错，返回false
		return false
	}
	if !match {
		// 身份证号码不符合规则，返回false
		return false
	}
	// 解析身份证号码中的年、月、日
	year, _ := strconv.Atoi(id[6:10])
	month, _ := strconv.Atoi(id[10:12])
	day, _ := strconv.Atoi(id[12:14])
	// 判断年份是否合法
	if year < 1900 || year > time.Now().Year() {
		return false
	}
	// 判断月份是否合法
	if month < 1 || month > 12 {
		return false
	}
	// 判断日期是否合法
	if day < 1 || day > 31 {
		return false
	}
	// 对身份证号码的最后一位进行校验
	// 根据身份证号码的规则，最后一位可能是数字0-9，也可能是字符X（表示10）
	// 将字符X转换成数字10进行校验
	lastChar := id[len(id)-1]
	var lastNum int
	if lastChar == 'X' || lastChar == 'x' {
		lastNum = 10
	} else {
		lastNum, _ = strconv.Atoi(string(lastChar))
	}
	// 对身份证号码的前17位进行加权和校验
	// 加权系数，根据规则固定
	weights := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 计算加权和
	sum := 0
	for i := 0; i < len(weights); i++ {
		num, _ := strconv.Atoi(string(id[i]))
		sum += num * weights[i]
	}

	// 计算校验码
	checkCode := sum % 11
	checkCodeMap := map[int]string{
		0:  "1",
		1:  "0",
		2:  "10", // 身份证最后一位是X，加权求和是10
		3:  "9",
		4:  "8",
		5:  "7",
		6:  "6",
		7:  "5",
		8:  "4",
		9:  "3",
		10: "2",
	}
	// 校验身份证号码的最后一位
	return checkCodeMap[checkCode] == strconv.Itoa(lastNum)
}

// IsIDCard15 验证15位身份证号
func IsIDCard15(idCard string) bool {
	// 验证是否为15位数字
	if match, _ := regexp.MatchString(`^\d{15}$`, idCard); !match {
		return false
	}

	// 将身份证号前两位转换成省份代码
	provinceCode, err := strconv.Atoi(idCard[:2])
	if err != nil || provinceCode < 11 || provinceCode > 91 {
		return false
	}

	// 验证生日是否正确
	year := strconv.Itoa(1900 + int(idCard[6]-'0')*10 + int(idCard[7]-'0'))
	month := string(idCard[8:10])
	day := string(idCard[10:12])
	if match, _ := regexp.MatchString(`^(19|20)\d{2}(0[1-9]|1[0-2])(0[1-9]|[12]\d|3[01])$`, year+month+day); !match {
		return false
	}

	return true
}

// IsIPv4 是否为ipv4地址
func IsIPv4(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() != nil
}

// IsIPv6 是否为ipv6地址
func IsIPv6(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() == nil
}

// IsAllChinese 验证给定的字符串全部为中文
func IsAllChinese(input string) bool {
	for _, r := range input {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}
	return true
}

// IsContainChinese 验证给定的字符串包含中文
func IsContainChinese(input string) bool {
	for _, r := range input {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// IsChineseName 验证是否为中文名
func IsChineseName(name string) bool {
	pattern := "^[\u4E00-\u9FA5]{2,6}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(name)
}

// IsEnglishName 验证是否为英文名
func IsEnglishName(name string) bool {
	match, _ := regexp.MatchString(`^([a-zA-Z]+\s)*[a-zA-Z]+$`, name)
	return match
}

// IsNumber 验证是否全部为数字
func IsNumber(input string) bool {
	reg := regexp.MustCompile("^[0-9]+$")
	return reg.MatchString(input)
}

// IsPostalCode 验证是否为邮编号码
func IsPostalCode(str string) bool {
	reg := regexp.MustCompile(`^[1-9]\d{5}$`)
	return reg.MatchString(str)
}

// IsTelephone 验证是否为座机号码
func IsTelephone(telephone string) bool {
	match, _ := regexp.MatchString(`^0\d{2,3}-?\d{7,8}$`, telephone)
	return match
}

var (
	alphaMatcher         *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z]+$`)
	letterRegexMatcher   *regexp.Regexp = regexp.MustCompile(`[a-zA-Z]`)
	intStrMatcher        *regexp.Regexp = regexp.MustCompile(`^[\+-]?\d+$`)
	urlMatcher           *regexp.Regexp = regexp.MustCompile(`^((ftp|http|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]+([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`)
	dnsMatcher           *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`)
	emailMatcher         *regexp.Regexp = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	chineseMobileMatcher *regexp.Regexp = regexp.MustCompile(`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
	chineseIdMatcher     *regexp.Regexp = regexp.MustCompile(`^[1-9]\d{5}(18|19|20|21|22)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	chineseMatcher       *regexp.Regexp = regexp.MustCompile("[\u4e00-\u9fa5]")
	chinesePhoneMatcher  *regexp.Regexp = regexp.MustCompile(`\d{3}-\d{8}|\d{4}-\d{7}`)
	creditCardMatcher    *regexp.Regexp = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$`)
	base64Matcher        *regexp.Regexp = regexp.MustCompile(`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`)
)

// IsAlpha 检查字符串是否只包含字母(a-zA-Z)。
func IsAlpha(str string) bool {
	return alphaMatcher.MatchString(str)
}

// IsAllUpper 检查字符串是否全是大写字母A-Z。
func IsAllUpper(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return str != ""
}

// IsAllLower 检查字符串是否全是小写字母a-z。
func IsAllLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return str != ""
}

// ContainUpper 检查字符串是否至少包含一个大写字母A-Z。
func ContainUpper(str string) bool {
	for _, r := range str {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// ContainLower 检查字符串是否至少包含一个小写字母a-z。
func ContainLower(str string) bool {
	for _, r := range str {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// ContainLetter 检查字符串是否至少包含一个字母。
func ContainLetter(str string) bool {
	return letterRegexMatcher.MatchString(str)
}

// IsJSON 检查字符串是否为有效的JSON。
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsNumberStr 检查字符串是否可以转换为数字。
func IsNumberStr(s string) bool {
	return IsIntStr(s) || IsFloatStr(s)
}

// IsFloatStr 检查字符串是否可以转换为浮点数。
func IsFloatStr(str string) bool {
	_, e := strconv.ParseFloat(str, 64)
	return e == nil
}

// IsIntStr 检查字符串是否可以转换为整数。
func IsIntStr(str string) bool {
	return intStrMatcher.MatchString(str)
}

// IsIp 检查字符串是否是有效ip地址
func IsIp(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	return ip != nil
}

// IsPort 检查字符串是否是有效端口
// Play:
func IsPort(str string) bool {
	if i, err := strconv.ParseInt(str, 10, 64); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

// IsUrl 检查字符串是否是url
func IsUrl(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return urlMatcher.MatchString(str)
}

// IsDns 检查字符串是否为DNS
func IsDns(dns string) bool {
	return dnsMatcher.MatchString(dns)
}

// IsEmail 检查字符串是否为电子邮件地址。
func IsEmail(email string) bool {
	return emailMatcher.MatchString(email)
}

// IsChineseMobile 检查字符串是否为中国手机号码。
func IsChineseMobile(mobileNum string) bool {
	return chineseMobileMatcher.MatchString(mobileNum)
}

// IsChineseIdNum 检查字符串是否为中国身份证。
func IsChineseIdNum(id string) bool {
	return chineseIdMatcher.MatchString(id)
}

// ContainChinese 检查字符串是否中文。
func ContainChinese(s string) bool {
	return chineseMatcher.MatchString(s)
}

// IsChinesePhone 检查字符串是否为中文电话号码。
// Valid chinese phone is xxx-xxxxxxxx or xxxx-xxxxxxx.
func IsChinesePhone(phone string) bool {
	return chinesePhoneMatcher.MatchString(phone)
}

// IsCreditCard 检查字符串是否为信用卡。
func IsCreditCard(creditCart string) bool {
	return creditCardMatcher.MatchString(creditCart)
}

// IsBase64 检查字符串是否为base64字符串。
func IsBase64(base64 string) bool {
	return base64Matcher.MatchString(base64)
}

// IsRegexMatch 检查字符串是否与regexp匹配。
func IsRegexMatch(str, regex string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(str)
}

// IsStrongPassword 检查字符串是否为强密码，如果len(password)小于length参数，返回false 强密码:α(低+上)+数字+特殊字符 (!@#$%^&*()?><).
func IsStrongPassword(password string, length int) bool {
	if len(password) < length {
		return false
	}
	var num, lower, upper, special bool
	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsLower(r):
			lower = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}

	return num && lower && upper && special
}

// IsWeakPassword 检查字符串是否为弱密码
func IsWeakPassword(password string) bool {
	var num, letter, special bool
	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsLetter(r):
			letter = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}
	return (num || letter) && !special
}

// IsGBK 判断是否是GBK编码
func IsGBK(data []byte) bool {
	i := 0
	for i < len(data) {
		if data[i] <= 0xff {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}

	return true
}

// IsPointer 检查变量是否指针类型;
// notNil 是否检查变量非nil.
func IsPointer(val any, notNil bool) (res bool) {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		if !notNil || (notNil && val != nil) {
			res = true
		}
	}
	return
}

// isBool 是否布尔值.
func isBool(val any) bool {
	return val == true || val == false
}

// isMap 检查变量是否字典.
func isMap(val any) bool {
	return reflect.ValueOf(val).Kind() == reflect.Map
}

// isString 变量是否字符串.
func isString(val any) bool {
	return function.GetVariateType(val) == "string"
}

// isByte 变量是否字节切片.
func isByte(val any) bool {
	return function.GetVariateType(val) == "[]uint8"
}

// isBinary 字符串是否二进制.
func isBinary(s string) bool {
	for _, b := range s {
		if 0 == b {
			return true
		}
	}
	return false
}

// isInt 变量是否整型数值.
func isInt(val any) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		_, err := strconv.Atoi(str)
		return err == nil
	}

	return false
}

// isNumeric 变量是否数值(不包含复数).
func isNumeric(val any) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		_, err := strconv.ParseFloat(str, 64)
		return err == nil
	}

	return false
}
