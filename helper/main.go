package helper

import (
	"crypto/md5"
	"fmt"
)


func Help(){
	fmt.Println("helper others")
}

func Md5(str string) string{
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
