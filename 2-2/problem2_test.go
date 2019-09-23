package problem2

import (
	"container/list"
	"testing"
)

func TestKthElement(t *testing.T) {
	c := list.New()
	c.PushBack(0)
	c.PushBack(1)
	c.PushBack(2)
	c.PushBack(3)
	c.PushBack(4)

	result := kthElement(c, 3)
	expected := 2
	if result.Value != expected {
		t.Fatalf("Expected: %d, Actual: %d", expected, result.Value)
	}
}

func TestRecursiveKthElement(t *testing.T) {
	c := list.New()
	c.PushBack(0)
	c.PushBack(1)
	c.PushBack(2)
	c.PushBack(3)
	c.PushBack(4)

	result := recursiveKthElement(c.Front(), 3)
	expected := 2
	if result.Value != expected {
		t.Fatalf("Expected: %d, Actual: %d", expected, result.Value)
	}
}
