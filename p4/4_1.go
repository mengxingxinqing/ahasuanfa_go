package p4

import "fmt"

//1-n张扑克 放到1-n号箱子里多少种放法

func dfs(arr []int, book []int, step int, n int) {
	if step == n {
		fmt.Println(arr)
		return
	}
	for i := 0; i < n; i++ {
		if book[i] == 0 {
			arr[step] = i + 1
			book[i] = 1
			dfs(arr, book, step+1, n)
			book[i] = 0
		}
	}
}
