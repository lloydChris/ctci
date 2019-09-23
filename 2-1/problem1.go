package problem1

import "container/list"

func removeDuplicates(l *list.List) {
	seenList := make(map[interface{}]bool)

	for e := l.Front(); e != nil; e = e.Next() {
		if seenList[e.Value] {
			l.Remove(e)
		} else {
			seenList[e.Value] = true
		}
	}
}
