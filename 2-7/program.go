package program

import "container/list"

func recursive(a *list.Element, b *list.Element) *list.Element {
	if a.Next() == nil && b.Next() == nil {
		if a == b {
			return a
		}
	}

	aNext := a.Next()
	if aNext == nil {
		aNext = a
	}

	bNext := b.Next()
	if bNext == nil {
		bNext = b
	}

	intersection := recursive(aNext, bNext)

	if intersection == a {
		return b
	}

	if intersection == b {
		return a
	}

	return intersection
}

func intersection(a list.List, b list.List) *list.Element {
	return recursive(a.Front(), b.Front())
}
