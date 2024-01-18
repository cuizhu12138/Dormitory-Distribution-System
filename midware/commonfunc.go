package midware

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

// 得到str的MD5值
func Token_md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// 得到4位的随机盐
func RandSalt() (salt string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		salt += fmt.Sprintf("%d", rand.Intn(10))
	}
	return
}
