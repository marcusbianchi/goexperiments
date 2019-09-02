package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Years(index)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for index := 0; index < b.N; index++ {
		YearsTwo(index)
	}
}

func ExampleYears() {
	fmt.Println(Years(2))
	// Output:
	// 14
}

func TestYears(t *testing.T) {
	type test struct {
		data   int
		answer int
	}

	tests := []test{
		test{1, 7},
		test{2, 14},
		test{3, 21},
	}

	for _, v := range tests {
		x := Years(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(2))
	// Output:
	// 14
}

func TestYearsTwo(t *testing.T) {
	type test struct {
		data   int
		answer int
	}

	tests := []test{
		test{1, 7},
		test{2, 14},
		test{3, 21},
	}

	for _, v := range tests {
		x := YearsTwo(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
