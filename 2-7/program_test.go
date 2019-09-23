package program

import (
	"container/list"
	"testing"
)

func TestIntersection(t *testing.T) {
	a := list.New()
	a.PushBack("a")
	a.PushBack("b")
	a.PushBack("c")
	a.PushBack("d")
	a.PushBack("d")
	a.PushBack("e")

	b := list.New()
	b.PushBack("z")
	b.PushBack("d")
	b.PushBack("e")
	b.PushBack("f")
}
