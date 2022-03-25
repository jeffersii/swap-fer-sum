package swapfersum

type Result struct {
	CanSwap     bool
	FirstIndex  int
	SecondIndex int
}

func sum(x []int) int {
	total := 0
	for _, val := range x {
		total += val
	}
	return total
}

func CanSwapForEqualSums(a, b []int) Result {
	sumA := sum(a)
	sumB := sum(b)
	diff := sumA - sumB
	result := Result{}

	if diff%2 != 0 {
		return result
	}

	bLookup := make(map[int]int)
	for secondIndex, val := range b {
		bLookup[val] = secondIndex
	}

	for firstIndex, val := range a {
		if secondIndex, found := bLookup[val-diff/2]; found {
			result.CanSwap = true
			result.FirstIndex = firstIndex
			result.SecondIndex = secondIndex
			break
		}
	}
	return result
}
