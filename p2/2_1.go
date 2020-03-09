package p2

import "fmt"

//解密qq号
//加密规则 先将第一位数字删除，第二位数字移动到尾部，第三位数字删除。。。。。
//12345678
//1 3456782 3 567824 5 78246 7 2468 2 684 6 48 4 8 8
// 1 3 5 7 2 6 4 8
//还原流程,从后往前 将最后一位移动到前面第二位，再把读到的数放在最前面，
// 8  48 684 2468 78246 567824 3456782 12345678
// 8  48 684
// 1 3 5 7 2 6 4 8

//数组方式实现加密过程
func QQEncodeWithArr(arr []int) []int {
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

func QQDecodeWithArr(arr []int) []int {
	if len(arr) <= 2 {
		return arr
	}
	for i := len(arr) - 3; i >= 0; i-- {
		DeRule(arr, i)
	}
	return arr
}

//将i位置后面的数据右移一位，然后将原来的最后一位放到i的下一位
func DeRule(arr []int, i int) {
	l := len(arr)
	end := arr[l-1]
	for j := l - 2; j > i; j-- {
		arr[j+1] = arr[j]
	}
	arr[i+1] = end
	fmt.Println(i, arr)
}

type Queue struct {
	h    int
	t    int
	data []interface{}
}

func NewQueue(cap int) *Queue {
	q := &Queue{}
	q.data = make([]interface{}, cap)
	return q
}

func (q *Queue) Add(i interface{}) {
	q.data[q.t] = i
	q.t++
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.Empty() {
		return 0, false
	}
	d := q.data[q.h]
	q.h++
	return d, true
}
func (q *Queue) Empty() bool {
	if q.h == q.t {
		return true
	}
	return false
}

//加密方法 用队列实现
func QQEncodeWithQueue(arr []int) []int {
	q := NewQueue(30)
	for _, d := range arr {
		q.Add(d)
	}
	res := make([]int, 0)
	for !q.Empty() {
		f, _ := q.Pop()
		res = append(res, f.(int))
		if s, ok := q.Pop(); ok {
			q.Add(s)
		}
	}
	return res
}
