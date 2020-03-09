package p2

import (
	"fmt"
	"testing"
)

func TestCheckHuiwen(t *testing.T) {
	res := CheckHuiwen("abab")
	fmt.Println(res)
}
