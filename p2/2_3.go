package p2

//小猫钓鱼游戏逻辑实现
//实现逻辑：用户 A,B的牌是两个队列，存放出牌的地方 使用栈

func XiaomaoDiaoyu(a []int, b []int) bool {
	aQ := NewQueue(200)
	bQ := NewQueue(200)
	for _, k := range a {
		aQ.Add(k)
	}
	for _, k := range b {
		bQ.Add(k)
	}
	outFlag := [15]int{}
	userList := [2]*Queue{aQ, bQ}

	s := NewStack(200)
	for !aQ.Empty() && !bQ.Empty() {
		for _, user := range userList {
			if ai, ok := user.Pop(); ok {
				s.Push(ai)
				//如果标识记录里有值，说明出过这个牌
				if outFlag[ai.(int)] == 1 {
					//赢牌操作
					for !s.Empty() {
						oi, _ := s.Pop()
						//赢牌 一直赢到自己出的牌相同的牌的位置，赢走的牌标识为0
						if oi != ai {
							user.Add(oi)
							outFlag[oi.(int)] = 0
						} else {
							break
						}
					}
					outFlag[ai.(int)] = 0
					//赢牌之后，再出个牌
					oi, _ := user.Pop()
					s.Push(oi)
				} else {
					outFlag[ai.(int)] = 1
				}
			}
		}

	}
	if aQ.Empty() {
		return false
	}
	return true
}
