package collections

import "testing"

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
