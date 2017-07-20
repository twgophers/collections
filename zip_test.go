package collections

import (
	"reflect"
	"testing"
)

func TestZip(t *testing.T) {
	cases := []struct {
		v1   []float64
		v2   []float64
		want []Pair
	}{
		{
			[]float64{1.0, 2.0, 3.0}, []float64{1.0, 3.0, 5.0},
			[]Pair{Pair{A: 1.0, B: 1.0}, Pair{A: 2.0, B: 3.0}, Pair{A: 3.0, B: 5.0}},
		},
		{
			[]float64{1.0, 2.0}, []float64{1.0, 3.0, 5.0},
			[]Pair{Pair{A: 1.0, B: 1.0}, Pair{A: 2.0, B: 3.0}},
		},
		{
			[]float64{1.0}, []float64{1.0, 3.0, 5.0},
			[]Pair{Pair{A: 1.0, B: 1.0}},
		},
		{
			[]float64{1.0, 2.0}, []float64{1.0},
			[]Pair{Pair{A: 1.0, B: 1.0}},
		},
		{
			[]float64{}, []float64{1.0, 3.0, 5.0}, []Pair{},
		},
		{
			[]float64{1.0, 3.0, 5.0}, []float64{}, []Pair{},
		},
	}
	for _, c := range cases {
		gotZip := Zip(c.v1, c.v2)
		if !reflect.DeepEqual(gotZip, c.want) {
			t.Errorf("Zip(%v, %v) want: %v; got: %v", c.v1, c.v2, c.want, gotZip)
		}
	}
}
