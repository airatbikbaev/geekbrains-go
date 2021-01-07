package hw4

import (
	"fmt"
	"testing"
)

var fibNumbers = []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}

func TestFibonacci(t *testing.T) {
	for i, val := range fibNumbers {
		res := getFibonacciNumber(uint64(i))

		if res != val {
			t.Fatalf("expected: %v, got: %v", val, res)
		}
	}
}

func TestTableFibonacci(t *testing.T) {
	tests := []struct {
		name  string
		input uint64
		want  uint64
	}{
		{name: "0", input: 0, want: 0},
		{name: "1", input: 1, want: 1},
		{name: "2", input: 2, want: 1},
		{name: "3", input: 3, want: 2},
		{name: "4", input: 4, want: 3},
		{name: "5", input: 5, want: 5},
		{name: "10", input: 10, want: 55},
		{name: "12", input: 12, want: 144},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			//t.Parallel()

			got := getFibonacciNumber(tc.input)

			if tc.want != got {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func ExampleFibonacci() {
	fmt.Println(getFibonacciNumber(1))
	// Output: 1
}
