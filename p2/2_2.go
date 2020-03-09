package p2

//回文检测 aba abba是回文 abab不是回文

type Stack struct {
	data []interface{}
	head int
}

func NewStack(cap int) *Stack {
	s := &Stack{}
	s.data = make([]interface{}, cap)
	return s
}

//0位置留空 从下标1开始使用
func (s *Stack) Pop() (interface{}, bool) {
	if s.Empty() {
		return 0, false
	}
	d := s.data[s.head]
	s.head--
	return d, true
}

func (s *Stack) Push(d interface{}) {
	s.head++
	s.data[s.head] = d
}

func (s *Stack) Empty() bool {
	if s.head == 0 {
		return true
	}
	return false
}

//回文
func CheckHuiwen(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return false
	}
	if sLen == 1 {
		return true
	}

	stack := NewStack(1000)
	for i, c := range s {
		if sLen%2 == 1 && i == sLen/2 {
			continue
		}
		if i < sLen/2 {
			stack.Push(c)
		}
		if i >= sLen/2 {
			t, _ := stack.Pop()

			if t != c {
				return false
			}
		}
	}
	return true
}
