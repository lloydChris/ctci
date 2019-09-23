package problem2

import (
	"container/list"
)

func kthElement(l *list.List, k int) list.Element {
	trailing := l.Front()
	for lead := l.Front(); lead != nil; lead = lead.Next() {
		if k > 0 {
			k--
		} else {
			trailing = trailing.Next()
		}
	}
	return *trailing
}

func recursiveKthElement2(e *list.Element, k int, counter *int) *list.Element {
	if e == nil {
		return nil
	}
	newE := recursiveKthElement2(e.Next(), k, counter)
	*counter++
	if *counter == k {
		return e
	}
	return newE
}

func recursiveKthElement(e *list.Element, k int) *list.Element {
	counter := 0
	return recursiveKthElement2(e, k, &counter)
}
