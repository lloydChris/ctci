package problem3

import (
	"container/list"
	"fmt"
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	c := list.New()
	c.PushBack(0)
	c.PushBack(1)
	c.PushBack(2)
	c.PushBack(3)
	c.PushBack(4)

	for lead := c.Front(); lead != nil; lead = lead.Next() {
		fmt.Printf("%d, ", lead.Value)
	}
	fmt.Println()

	deleteMiddle(c.Front().Next())

	for lead := c.Front(); lead != nil; lead = lead.Next() {
		fmt.Printf("%d, ", lead.Value)
	}
	fmt.Println()

	// if result.Value != expected {
	// 	t.Fatalf("Expected: %d, Actual: %d", expected, result.Value)
	// }
}
