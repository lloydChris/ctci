package partition

import (
	"container/list"
	"fmt"
	"testing"
)

func print(l list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d,", e.Value)
	}
	fmt.Print("\n")
}
func TestPartition(t *testing.T) {
	l := list.New()
	l.PushBack(3)
	l.PushBack(5)
	l.PushBack(8)
	l.PushBack(5)
	l.PushBack(10)
	l.PushBack(2)
	l.PushBack(1)

	print(*l)
	partition(l, 5)
	print(*l)
}
