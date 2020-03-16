package p2

func RemoveDuplicates(arr []int) []int {
	slow := 0
	fast := 1
	n := len(arr)
	for fast < n {
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
		fast++
	}
	return arr
}
