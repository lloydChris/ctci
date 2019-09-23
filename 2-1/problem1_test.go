package problem1

import (
	"container/list"
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCase := list.New()
	testCase.PushBack(1)
	testCase.PushBack(2)
	testCase.PushBack(2)
	testCase.PushBack(3)

	for e := testCase.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	removeDuplicates(testCase)

	for e := testCase.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
