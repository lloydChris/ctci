package sumlist

import "container/list"

func getInt(l *list.List) int {
	result := 0
	orderOfMagnitude := 1
	for e := l.Front(); e != nil; e = e.Next() {
		result += e.Value.(int) * orderOfMagnitude
		orderOfMagnitude *= 10
	}
	return result
}

func sumList(l1 *list.List, l2 *list.List) list.List {
	num1 := getInt(l1)
	num2 := getInt(l2)

	sumed := num1 + num2

	result := list.New()
	for sumed > 0 {
		remainder := sumed % 10

		result.PushBack(remainder)
		sumed -= remainder
		sumed /= 10
	}

	return *result
}

func reverseList(l *list.List) *list.List {
	lReversed := list.New()

	for e := l.Front(); e != nil; e = e.Next() {
		lReversed.PushFront(e.Value.(int))
	}
	return lReversed
}

func sumListForward(l1 *list.List, l2 *list.List) *list.List {
	l1Reversed := reverseList(l1)
	l2Reversed := reverseList(l2)

	result := sumList(l1Reversed, l2Reversed)

	return reverseList(&result)
}

func sumlistAsList(l1 *list.List, l2 *list.List) *list.List {
	e1 := l1.Front()
	e2 := l2.Front()
	result := list.New()

	carry := 0

	for e1 != nil && e2 != nil {
		sum := e1.Value.(int) + e2.Value.(int) + carry

		//max possible value is 18
		if sum > 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		result.PushBack(sum)
		e1 = e1.Next()
		e2 = e2.Next()
	}
	return result
}

func sumListRecursive(l1 *list.List, l2 *list.List, carry int) *list.List {
	result := list.New()

	e1 := l1.Front()
	e2 := l2.Front()

	if e1 == nil && e2 == nil {
		result.PushBack(carry)
		return result
	}

	e1Value := 0
	if e1 != nil {
		e1Value = e1.Value.(int)
		l1.Remove(e1)
	}

	e2Value := 0
	if e2 != nil {
		e2Value = e2.Value.(int)
		l2.Remove(e2)
	}

	sum := e1Value + e2Value + carry

	resultElementValue := sum % 10

	nextCarry := (sum - resultElementValue) / 10

	result.PushBack(resultElementValue)

	restOfData := sumListRecursive(l1, l2, nextCarry)
	result.PushBackList(restOfData)

	return result
}
