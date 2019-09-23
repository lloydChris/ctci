package partition

import (
	"container/list"
)

func partition(input *list.List, partition int) {
	e := input.Front()
	for e != nil {
		if e.Value.(int) < partition {
			temp := e.Next()
			input.MoveToFront(e)
			e = temp
		} else {
			e = e.Next()
		}
	}
}
