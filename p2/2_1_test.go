package p2

import (
	"fmt"
	"testing"
)

func TestQQEncryption(t *testing.T) {
	arr := []int{6, 1, 5, 9, 4, 7, 2, 8, 3}
	res := QQEncodeWithArr(arr)
	fmt.Println(res)
	res2 := QQDecodeWithArr(res)
	fmt.Println(res2)

	arr3 := []int{6, 3, 1, 7, 5, 8, 9, 2, 4}
	res3 := QQEncodeWithQueue(arr3)
	fmt.Println("res3", res3)
}
