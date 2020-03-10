package p4

import "fmt"

//解 xxx + xxx = xxx 有多少种可能的组合方式 1-9选择

func compute(arr []int, book []int, step int, n int, total int) int {
	if step == n {
		if arr[0]*100+arr[1]*10+arr[2]+arr[3]*100+arr[4]*10+arr[5] == arr[7]*100+arr[8]*10+arr[9] {
			fmt.Println(arr[0]*100+arr[1]*10+arr[2], "+", arr[3]*100+arr[4]*10+arr[5], "=", arr[7]*100+arr[8]*10+arr[9])

			return total + 1
		}
		return total
	}
	for i := 1; i <= 9; i++ {
		if book[i] == 0 {
			book[i] = 1
			arr[step-1] = i
			total = compute(arr, book, step+1, n, total)
			book[i] = 0
		}
	}
	return total
}
