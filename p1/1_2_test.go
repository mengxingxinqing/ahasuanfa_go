package p1

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 6, 7, 2, 1, 3}
	res := BubbleSort(arr)
	fmt.Println(res)
}
