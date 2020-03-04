package p1

import "fmt"

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	fmt.Println("left", left, "right", right)
	flag := arr[left]
	leftIndex := left
	rightIndex := right
	for leftIndex != rightIndex {
		//找到右边比基准数小的值的位置，如果比基准值大，左移
		for arr[rightIndex] >= flag && leftIndex < rightIndex {
			rightIndex--
		}
		//找到左边比基准数大的值的位置，如果比基准值小，右移
		for arr[leftIndex] <= flag && leftIndex < rightIndex {
			leftIndex++
		}
		//找到交换的位置，交换数值
		if leftIndex < rightIndex {
			fmt.Println("交换", leftIndex, rightIndex, arr[leftIndex], arr[rightIndex])
			arr[leftIndex], arr[rightIndex] = arr[rightIndex], arr[leftIndex]
		}

	}
	//将基准数交换位置
	arr[leftIndex], arr[left] = arr[left], arr[leftIndex]
	fmt.Println("leftIndex", leftIndex, "rightIndex", rightIndex)
	fmt.Println(arr)
	QuickSort(arr, left, leftIndex-1)
	QuickSort(arr, rightIndex+1, right)
	return
}
