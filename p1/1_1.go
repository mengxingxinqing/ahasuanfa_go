package p1

//简单版桶排序算法，需要需要空间跟排序中最大的数一样长
func SimpleSort(arr []int) []int {
	max := 1000
	storeArr := make([]int, max)
	resArr := make([]int, len(arr))
	for _, k := range arr {
		storeArr[k-1] += 1
	}
	index := 0
	for i, k := range storeArr {
		for j := 0; j < k; j++ {
			resArr[index] = i + 1
			index++
		}
	}
	return resArr
}
