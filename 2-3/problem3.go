package problem3

import (
	"container/list"
	"fmt"
)

func deleteMiddle(e *list.Element) {
	// f := e.Next()
	e = e.Next()
	fmt.Printf("%v\n", e.Value)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", *e)
	fmt.Printf("%v\n", &e)
	return
}
