package p2

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	arr := []int{2, 3, 3, 3, 4, 4, 5, 6, 7, 7, 8}
	fmt.Println(RemoveDuplicates(arr))
}
