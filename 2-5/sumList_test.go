package sumlist

import (
	"container/list"
	"fmt"
	"testing"
)

func TestSumList(t *testing.T) {
	l1 := list.New()
	l1.PushBack(7)
	l1.PushBack(1)
	l1.PushBack(6)

	l2 := list.New()
	l2.PushBack(5)
	l2.PushBack(9)
	l2.PushBack(2)

	result := sumList(l1, l2)

	fmt.Print("SumList: ")
	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}

func TestSumListForward(t *testing.T) {
	l1 := list.New()
	l1.PushBack(6)
	l1.PushBack(1)
	l1.PushBack(7)

	l2 := list.New()
	l2.PushBack(2)
	l2.PushBack(9)
	l2.PushBack(5)

	result := sumListForward(l1, l2)

	fmt.Print("SumListForward: ")
	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}

func TestSumlistAsList(t *testing.T) {
	l1 := list.New()
	l1.PushBack(7)
	l1.PushBack(1)
	l1.PushBack(6)

	l2 := list.New()
	l2.PushBack(5)
	l2.PushBack(9)
	l2.PushBack(2)

	result := sumlistAsList(l1, l2)

	fmt.Print("SumListAsList: ")
	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}

func TestSumListRecursive(t *testing.T) {
	l1 := list.New()
	l1.PushBack(9)
	l1.PushBack(1)
	l1.PushBack(6)

	l2 := list.New()
	l2.PushBack(9)
	l2.PushBack(9)
	l2.PushBack(2)
	l2.PushBack(9)

	result := sumListRecursive(l1, l2, 0)

	fmt.Print("SumListRecursive: ") //8199
	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()

	l1 = list.New()
	l1.PushBack(9)
	l1.PushBack(7)
	l1.PushBack(8)

	l2 = list.New()
	l2.PushBack(6)
	l2.PushBack(8)
	l2.PushBack(5)

	result = sumListRecursive(l1, l2, 0)

	fmt.Print("SumListRecursive2: ") //8199
	for e := result.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}
