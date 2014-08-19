package Library

import (
	"encoding/hex"
	"crypto/md5"
	"io"
	"crypto/rand"
	"encoding/base64"
	math_rand "math/rand"
	"time"
)

// 字符串
// md5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
// Guid
func NewGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}
// 后面加个str生成之, 更有保障, 确保唯一
func NewGuidWith(str string) string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString([]byte(string(b) + str)))
}

// 随机密码
// num 几位
func RandomPwd(num int) string {
	chars := make([]byte, 62)
	j := 0
	for i := 48; i <= 57; i++ {
		chars[j] = byte(i)
		j++
	}
	for i := 65; i <= 90; i++ {
		chars[j] = byte(i)
		j++
	}
	for i := 97; i <= 122; i++ {
		chars[j] = byte(i)
		j++
	}
	j--;
	str := ""
	math_rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		x := math_rand.Intn(j)
		str += string(chars[x])
	}
	return str
}
