package main

import (
	"fmt"
	"testing"
)

var EMPTY_LIST = make([]int, 0, 10)

func TestGenerateOne(t *testing.T) {
	factors := generate(1)

	if ne(factors, EMPTY_LIST) {
		t.Fatalf("Empty list is not equal to %v", factors)
	}
}

func TestGenerateTwo(t *testing.T) {
	factors := generate(2)
	numbers := list(2)
	if ne(factors, numbers) {
		t.Fatalf("%v is not equal to %v", factors, numbers)
	}
}

func TestGenerateThree(t *testing.T) {
	factors := generate(3)
	numbers := list(3)
	if ne(factors, numbers) {
		t.Fatalf("%v is not equal to %v", factors, numbers)
	}
}

func TestGenerateFour(t *testing.T) {
	factors := generate(4)
	numbers := list(2, 2)
	if ne(factors, numbers) {
		t.Fatalf("%v is not equal to %v", factors, numbers)
	}
}

func TestGenerateSix(t *testing.T) {
	factors := generate(6)
	numbers := list(2, 3)
	if ne(factors, numbers) {
		t.Fatalf("%v is not equal to %v", factors, numbers)
	}
}

func TestGenerateEight(t *testing.T) {
	factors := generate(8)
	numbers := list(2, 2, 2)
	if ne(factors, numbers) {
		t.Fatalf("%v is not equal to %v", factors, numbers)
	}
}

func TestGenerateNine(t *testing.T) {
	factors := generate(9)
	numbers := list(3, 3)
	if ne(factors, numbers) {
		t.Fatalf("%v is not equal to %v", factors, numbers)
	}
}

func TestGenerate13195(t *testing.T) {
	factors := generate(13195)
	numbers := list(5, 7, 13, 29)
	if ne(factors, numbers) {
		t.Fatalf("%v is not equal to %v", factors, numbers)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(1232342342)
	}
}

func list(numbers ...int) []int {
	numList := make([]int, 0)

	for _, number := range numbers {
		numList = append(numList, number)
	}

	return numList
}

func ne(a, b []int) bool {
	return fmt.Sprintf("%v", a) != fmt.Sprintf("%v", b)
}

func eq(a, b []int) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}
