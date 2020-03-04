package p2

import "fmt"

//解密qq号
//加密规则 先将第一位数字删除，第二位数字移动到尾部，第三位数字删除。。。。。
//12345678
//1 3456782 3 567824 5 78246 7 2468 2 684 6 48 4 8 8
// 1 3 5 7 2 6 4 8
//还原流程,从后往前 将最后一位移动到前面第二位，再把读到的数放在最前面，
// 8  48 684 2468 78246 567824 3456782 12345678

//数组方式实现加密过程
func QQEncryptionWithArr(arr []int) []int {
	l := len(arr)
	if l <= 2 {
		return arr
	}
	res := make([]int, len(arr))
	index := 0
	for {
		top := Rule(arr, l-index)
		res[index] = top
		if index >= l-1 {
			break
		}
		index++
	}
	return res
}

//弹出一个头，拿到第二位的值，全体左移两位，然后把第二位的值放在倒数第二位
func Rule(arr []int, l int) int {
	if l == 1 {
		top := arr[0]
		arr[0] = 0
		fmt.Println(l, "top", top, arr)
		return top
	}
	top := arr[0]
	second := arr[1]
	for i := 0; i < l-2; i++ {
		arr[i] = arr[i+2]
	}
	arr[l-2] = second
	arr[l-1] = 0
	fmt.Println(l, "top", top, arr)
	return top
}
