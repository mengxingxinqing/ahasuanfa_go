package p2

import (
	"fmt"
	"testing"
)

func TestQQEncryption(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res := QQEncryptionWithArr(arr)
	fmt.Println(res)
}
