package chop

import (
	"testing"
)

func TestChop(t *testing.T) {

	assertEquals(t, -1, chop(3, []int{}))
	assertEquals(t, -1, chop(3, []int{1}))
	assertEquals(t, 0, chop(1, []int{1}))
	assertEquals(t, 0, chop(1, []int{1, 3, 5, 7}))
	assertEquals(t, 1, chop(3, []int{1, 3, 5, 7}))
	assertEquals(t, -1, chop(8, []int{1, 3, 5, 7}))
	assertEquals(t, 1, chop(3, []int{1, 3, 5, 7, 9}))
	assertEquals(t, 1, chop(3, []int{1, 3, 5, 7, 9, 11}))
}

func assertEquals(t *testing.T, expected int, actual int) {
	if expected != actual {
		t.Errorf("Expected = %v, actual = %v", expected, actual)
	}
}
