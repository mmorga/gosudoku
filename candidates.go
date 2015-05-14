package gosudoku

import "sort"

func CandidateSetsEqual(set1 []int, set2 []int) bool {
	if len(set1) != len(set2) {
		return false
	}
	sort.Ints(set1)
	sort.Ints(set2)
	for i, val := range set1 {
		if val != set2[i] {
			return false
		}
	}
	return true
}

func CandidatesContainAny(candidates []int, vals []int) bool {
	for _, val := range vals {
		for _, candidate := range candidates {
			if val == candidate {
				return true
			}
		}
	}
	return false
}

func isCompleteSet(vals []int) bool {
	if len(vals) != 9 {
		return false
	}
	sort.Ints(vals)

	for i := 1; i < 10; i++ {
		if i != vals[i-1] {
			return false
		}
	}
	return true
}
