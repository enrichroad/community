package strcase

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"

	"github.com/PuerkitoBio/goquery"
	uuid "github.com/iris-contrib/go.uuid"
)

/*
IsBlank checks if a string is whitespace or empty (""). Observe the following behavior:
    goutils.IsBlank("")        = true
    goutils.IsBlank(" ")       = true
    goutils.IsBlank("bob")     = false
    goutils.IsBlank("  bob  ") = false
Parameter:
    str - the string to check
Returns:
    true - if the string is whitespace or empty ("")
*/
func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func IsAnyBlank(strs ...string) bool {
	for _, str := range strs {
		if IsBlank(str) {
			return true
		}
	}
	return false
}

func DefaultIfBlank(str, def string) string {
	if IsBlank(str) {
		return def
	} else {
		return str
	}
}

// IsEmpty checks if a string is empty (""). Returns true if empty, and false otherwise.
func IsEmpty(str string) bool {
	return len(str) == 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// 截取字符串
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

// UUID
func UUID() string {
	u, _ := uuid.NewV4()
	return strings.ReplaceAll(u.String(), "-", "")
}

func Equals(a, b string) bool {
	return a == b
}

func EqualsIgnoreCase(a, b string) bool {
	return a == b || strings.ToUpper(a) == strings.ToUpper(b)
}

// RuneLen 字符成长度
func RuneLen(s string) int {
	bt := []rune(s)
	return len(bt)
}

// GetSummary 获取summary
func GetSummary(s string, length int) string {
	s = strings.TrimSpace(s)
	summary := Substr(s, 0, length)
	if RuneLen(s) > length {
		summary += "..."
	}
	return summary
}

// GetHtmlText 获取html文本
func GetHtmlText(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		logrus.Error(err)
		return ""
	}
	return doc.Text()
}

// ToInt64 str to int64，如果转换失败，默认值为0
// str 字符串
func ToInt64(str string) int64 {
	return ToInt64ByDefault(str, 0)
}

// ToInt64ByDefault str to int64
// str 字符串
// def 如果转换失败使用的默认值
func ToInt64ByDefault(str string, def int64) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		val = def
	}
	return val
}

// ToInt str to int，如果转换失败，默认值为0
// str 字符串
func ToInt(str string) int {
	return ToIntByDefault(str, 0)
}

// ToIntByDefault str to int
// str 字符串
// def 如果转换失败使用的默认值
func ToIntByDefault(str string, def int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		val = def
	}
	return val
}
