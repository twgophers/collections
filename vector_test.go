package collections

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	cases := []struct {
		Vector
		want float64
	}{
		{Vector{11.0}, 11.0},
		{Vector{11.0, 12.0}, 12.0},
		{Vector{11.0, 20.0, 12.0}, 20.0},
	}

	for _, c := range cases {
		gotMax := c.Max()

		if gotMax != c.want {
			t.Errorf("Expected max (%v) in (%v) but got (%v)", c.want, c.Vector, gotMax)
		}
	}
}

func TestMaxFail(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected panic when empty sample")
		}
	}()

	Vector{}.Max()
}

func TestMin(t *testing.T) {
	cases := []struct {
		Vector
		want float64
	}{
		{Vector{13.0}, 13.0},
		{Vector{12.0, 13.0}, 12.0},
		{Vector{12.0, 11.0, 13.0}, 11.0},
	}
	for _, c := range cases {
		gotMin := c.Min()
		if gotMin != c.want {
			t.Errorf("Expected max (%v) in (%v) but got (%v)", c.want, c.Vector, gotMin)
		}
	}
}

func TestMinFail(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Expected panic when empty sample")
		}
	}()

	Vector{}.Min()
}

func TestEmpty(t *testing.T) {
	cases := []struct {
		Vector
		want bool
	}{
		{Vector{1.0}, false},
		{Vector{}, true},
	}

	for _, c := range cases {
		got := c.Empty()
		if got != c.want {
			t.Errorf("(%v).Empty() want: %v but got: %v",
				c.Vector, c.want, got)
		}
	}
}

func TestLess(t *testing.T) {
	cases := []struct {
		Vector
		i, j int
		want bool
	}{
		{Vector{1.0, 2.0}, 0, 1, true},
		{Vector{1.0, 2.0}, 1, 0, false},
		{Vector{1.0, 1.0}, 0, 1, false},
		{Vector{1.0, 1.0}, 1, 0, false},
		{Vector{1.0}, 0, 0, false},
	}

	for _, c := range cases {
		got := c.Less(c.i, c.j)
		if got != c.want {
			t.Errorf("(%v).Less(%d, %d) want: %v but got: %v",
				c.Vector, c.i, c.j, c.want, got)
		}
	}
}

func TestLess_WhenTheValuesAreGreaterThanVector(t *testing.T) {
	cases := []struct {
		Vector
		i, j int
	}{
		{Vector{}, 0, 1},
		{Vector{1.0}, 0, 2},
	}
	for _, c := range cases {
		defer func() {
			if recover() == nil {
				t.Error("Expected panic when the parameters are greater than Vector size")
			}
		}()

		c.Less(c.i, c.j)
	}
}

func TestSwap(t *testing.T) {
	cases := []struct {
		before Vector
		i, j   int
		want   Vector
	}{
		{Vector{1.0, 2.0, 3.0}, 0, 2, Vector{3.0, 2.0, 1.0}},
		{Vector{1.0, 2.0}, 1, 0, Vector{2.0, 1.0}},
		{Vector{1.0, 1.0}, 0, 1, Vector{1.0, 1.0}},
		{Vector{1.0}, 0, 0, Vector{1.0}},
	}

	for _, c := range cases {
		c.before.Swap(c.i, c.j)
		if !reflect.DeepEqual(c.before, c.want) {
			t.Errorf("(%v).Swap(%d, %d) want: %v; got: %v",
				c.want, c.i, c.j, c.want, c.before)
		}
	}
}

func TestSwap_WhenTheValuesAreGreaterThanVector(t *testing.T) {
	cases := []struct {
		Vector
		i, j int
	}{
		{Vector{1.0, 2.0, 3.0}, 0, 4},
		{Vector{}, 0, 0},
	}

	for _, c := range cases {
		defer func() {
			if recover() == nil {
				t.Error("Expected panic when the parameters are greater than Vector size")
			}
		}()
		c.Swap(c.i, c.j)
	}
}
