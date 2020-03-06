package p2

import (
	"fmt"
	"testing"
)

func TestQQEncryption(t *testing.T) {
	arr := []int{7, 8, 9, 5, 6, 11}
	res := QQEncodeWithArr(arr)
	fmt.Println(res)
	res2 := QQDecodeWithArr(res)
	fmt.Println(res2)
}
