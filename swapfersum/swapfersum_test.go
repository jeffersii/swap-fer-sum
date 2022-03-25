package swapfersum

import (
	"math/rand"
	"testing"
	"time"
)

const MaxArrayElemSize = 1000000

func TestDetectsSwappableArrays(t *testing.T) {
	a := []int{1, 3}
	b := []int{2, 4}

	result := CanSwapForEqualSums(a, b)
	expectedResult := Result{
		CanSwap:     true,
		FirstIndex:  0,
		SecondIndex: 0,
	}
	if !result.CanSwap {
		t.Errorf("Swappable arrays %v, %v reported as non-swappable", a, b)
	}
	if result.FirstIndex != expectedResult.FirstIndex || result.SecondIndex != expectedResult.SecondIndex {
		t.Errorf("Swappable arrays %v, %v should have swap indices %d, %d but got %d, %d",
			a, b, expectedResult.FirstIndex, expectedResult.SecondIndex, result.FirstIndex, result.SecondIndex)
	}
}

func TestDetectsNonSwappableArrays(t *testing.T) {
	a := []int{1, 8}   // Sum is 9
	b := []int{10, 11} // Sum is 21, difference is 12 (even)

	result := CanSwapForEqualSums(a, b)
	if result.CanSwap {
		t.Errorf("Non-swappable arrays %v, %v reported as swappable", a, b)
	}
}

func TestDetectsNonSwappableOddDifferenceArrays(t *testing.T) {
	a := []int{1, 3} // Sum is 4
	b := []int{2, 5} // Sum is 7, difference is 3 (odd)

	result := CanSwapForEqualSums(a, b)
	if result.CanSwap {
		t.Errorf("Non-swappable (odd difference) arrays %v, %v reported as swappable", a, b)
	}
}

func randInt() int {
	return -MaxArrayElemSize + rand.Intn(2*MaxArrayElemSize)
}

func FuzzCanSwapForEqualSums(f *testing.F) {
	rand.Seed(time.Now().UnixNano())
	f.Add(uint(2))
	f.Fuzz(func(t *testing.T, n uint) {
		a := make([]int, n)
		b := make([]int, n)
		var i uint
		for i = 0; i < n; i++ {
			a[i] = randInt()
			b[i] = randInt()
		}
		result := CanSwapForEqualSums(a, b)
		if result.CanSwap {
			// Actually swap and sum to verify
			temp := a[result.FirstIndex]
			a[result.FirstIndex] = b[result.SecondIndex]
			b[result.SecondIndex] = temp

			if sum(a) != sum(b) {
				t.Errorf("arrays %v, %v should be swappable but are not", a, b)
			}
		}
	})
}
