package hw3

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var unsorted = []int{3, 6, 1, 7, 10, 4234, 45, 654}
var sorted = []int{1, 3, 6, 7, 10, 45, 654, 4234}

var GlobalTestSlice = []int{}

func TestBubbleSort(t *testing.T) {
	got := GetBubbleSortedSlice(unsorted)

	if !reflect.DeepEqual(sorted, got) {
		t.Fatalf("expected: %v, got: %v", sorted, got)
	}
}

func TestInsertionSort(t *testing.T) {
	got := GetInsertionSortedSlice(unsorted)

	assert.Equal(t, sorted, got, "they should be equal")
}

func BenchmarkBubbleSort(b *testing.B) {
	x := []int{}

	for i := 0; i < b.N; i++ {
		x = GetBubbleSortedSlice(unsorted)
	}

	GlobalTestSlice = x
}

func BenchmarkInsertionSort(b *testing.B) {
	x := []int{}

	for i := 0; i < b.N; i++ {
		x = GetInsertionSortedSlice(unsorted)
	}

	GlobalTestSlice = x
}
