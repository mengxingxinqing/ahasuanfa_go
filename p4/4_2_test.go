package p4

import (
	"fmt"
	"testing"
)

func TestCompute(t *testing.T) {
	arr := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	book := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	res := compute(arr, book, 1, len(arr), 0)
	fmt.Println(res)
}
