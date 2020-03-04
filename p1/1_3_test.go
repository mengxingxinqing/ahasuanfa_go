package p1

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 6, 7, 2, 1, 3, 19, 23}
	fmt.Println(arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
