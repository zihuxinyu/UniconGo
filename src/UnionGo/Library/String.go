package Library

import (
	"regexp"
	"strings"
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

func StringInSlice(str string, list []string) bool {
	for _, element := range list {
		if element == str {
			return true
		}
	}
	return false
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
