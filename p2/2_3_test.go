package p2

import (
	"fmt"
	"testing"
)

func TestXiaomaoDiaoyu(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8, 9}
	res := XiaomaoDiaoyu(a, b)
	fmt.Println(res)
}
