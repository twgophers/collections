package collections

type Pair struct {
	A, B float64
}

func Zip(a, b []float64) []Pair {

	if len(a) == 0 || len(b) == 0 {
		return []Pair{}
	}

	if len(a) < len(b) {
		return createPairs(a, b)
	}
	return createPairs(a[:len(b)], b)

}

func createPairs(a, b []float64) []Pair {
	Pairs := make([]Pair, len(a), len(a))

	for i, e := range a {
		Pairs[i] = Pair{e, b[i]}
	}
	return Pairs

}
