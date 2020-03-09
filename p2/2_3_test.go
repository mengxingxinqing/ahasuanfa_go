package p2

import (
	"fmt"
	"testing"
)

func TestXiaomaoDiaoyu(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{1, 2, 3, 4, 5, 6}
	res := XiaomaoDiaoyu(a, b)
	fmt.Println(res)
}
