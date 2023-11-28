package auth

import (
	"errors"
	"fmt"
	"regexp"

	encrypt "github.com/wumansgy/goEncrypt/aes"
)

func CheckPassword(in string) error {
	if len(in) < 8 || len(in) > 32 {
		return errors.New("密码长度应该在8-32之间")
	}
	re := "^(?=.*\\d)(?=.*[a-z])(?=.*[A-Z]).{8,32}$"
	b, err := regexp.MatchString(re, in)
	if !b || err != nil {
		return fmt.Errorf("强密码校验失败: %v", err)
	}
	return nil
}

func encode(msg, key string) string {
	out, err := encrypt.AesCbcEncryptBase64([]byte(msg), []byte(key), nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return out
}

func decode(base64Hex, key string) string {
	out, err := encrypt.AesCbcDecryptByBase64(base64Hex, []byte(key), nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(out)
}
