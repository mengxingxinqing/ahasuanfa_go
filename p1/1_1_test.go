package p1

import (
	"fmt"
	"testing"
)

func TestSimpleSort(t *testing.T) {
	arr := []int{5, 6, 7, 2, 1, 3}
	res := SimpleSort(arr)
	fmt.Println(res)
}
