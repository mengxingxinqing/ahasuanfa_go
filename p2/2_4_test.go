package p2

import (
	"fmt"
	"testing"
)

func TestRevert(t *testing.T) {
	h := &Node{val: 1}
	h2 := &Node{val: 2}
	h3 := &Node{val: 3}
	h4 := &Node{val: 4}
	h.next = h2
	h2.next = h3
	h3.next = h4
	h = revert(h)
	for h != nil {
		fmt.Println(h.val)
		h = h.next
	}
}
