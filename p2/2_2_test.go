package p2

import (
	"fmt"
	"testing"
)

func TestCheckHuiwen(t *testing.T) {
	res := CheckHuiwen("abab")
	fmt.Println(res)
	res = CheckHuiwen("aba")
	fmt.Println(res)
	res = CheckHuiwen("abba")
	fmt.Println(res)
}
