package sudoku

type SudokoValueSet interface {
	Reduce([]int) BaseSet
	Keys() []int
	AddKeys(keys []int) BaseSet
}

type BaseSet map[int]bool

func NewSudukoValueSet() (s BaseSet) {
	s = map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
		6: true,
		7: true,
		8: true,
		9: true,
	}

	return s
}

func (s BaseSet) Reduce(vals []int) BaseSet {
	for _, val := range vals {
		delete(s, val)
	}
	return s
}

func (s BaseSet) Keys() (keys []int) {
	for k, _ := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s BaseSet) AddKeys(keys []int) BaseSet {
	for _, key := range keys {
		s[key] = true
	}
	return s
}
