package Library

import (
	"regexp"
	"strings"
	"strconv"
)
var (
	lowerRe, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	brRe, _ = regexp.Compile("<br.*?>")
	tagRe, _ = regexp.Compile("<.*?>")
)
func StripTags(html string) string {
	html = lowerRe.ReplaceAllStringFunc(html, strings.ToLower)
	html = strings.Replace(html, "\n", " ", -1)
	html = strings.Replace(html, "\r", "", -1)
	html = strings.Replace(html, "&nbsp;", " ", -1)
	html = brRe.ReplaceAllString(html, "")
	html = tagRe.ReplaceAllString(html, "")
	return html
}
// Simplify HTML text by removing tags
func RemoveFormatting(html string) string {
	return StripTags(html)
}



//列表是否包含给定项
func ListContains(list []interface{}, key interface{}) (finded bool) {
	for _, v := range list {
		if v == key {
			finded = true
			break
		}
	}
	return
}
//字符串数组中是否包含给定项
func StringsContains(list []string, key string) (finded bool) {
	for _, v := range list {
		if v == key {
			finded = true
			break
		}
	}
	return
}


func StringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, e := range a {
		if e != b[i] {
			return false
		}
	}
	return true
}

//字符串转长整型
func Str2int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
//字符串转整形
func Str2int(s string) (int, error) {
	return strconv.Atoi(s)
}
//整形转字符串
func Int2str(i int) string {
	return strconv.Itoa(i)
}
