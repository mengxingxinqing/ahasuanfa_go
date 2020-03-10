package p4

import "testing"

func TestDfs(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	book := []int{0, 0, 0, 0}
	n := len(arr)
	dfs(arr, book, 0, n)
}
