package auth

import (
	"fmt"
	"testing"
)

func Test_encode(t *testing.T) {

	org := "abcdefg"
	key := "123456781234567812345678"

	pass := encode(org, key)
	fmt.Println(pass)

	msg := decode(pass, key)
	fmt.Println(msg)
}
