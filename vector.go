package collections

import (
	"math"
)

type binaryCondition func(v1, v2 float64) float64
type Vector []float64

func (v Vector) Max() float64 {
	return matchingValue(math.Max, math.Inf(-1), v)
}

func (v Vector) Min() float64 {
	return matchingValue(math.Min, math.Inf(+1), v)
}

func (v Vector) Empty() bool {
	return len(v) == 0
}

func matchingValue(fn binaryCondition, initial float64, vector Vector) float64 {
	validate(vector)

	current := initial
	for _, value := range vector {
		current = fn(current, value)
	}
	return current
}

func validate(v Vector) {
	if len(v) == 0 {
		panic("empty sample supplyed")
	}
}
