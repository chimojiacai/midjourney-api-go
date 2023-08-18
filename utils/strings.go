// @Author: liyongzhen
// @Description:
// @File: strings
// @Date: 2023/8/18 18:04

package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func Md5(str string) (md5str string) {
	data := []byte(str)
	has := md5.Sum(data)
	md5str = fmt.Sprintf("%x", has)
	return
}

// RandomStr 随机生成字符串
func RandomStr(n int) string {
	var l = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	length := len(l)
	for i := range b {
		b[i] = l[rand.Intn(length)]
	}
	// 增加随机性
	return Md5(fmt.Sprintf("%s_%s", string(b), time.Now()))
}
