package cyclic

import "testing"

func TestCyclic(t *testing.T) {
	type link struct {
		tail *link
	}
	a, b, c := &link{}, &link{}, &link{}
	a.tail, b.tail, c.tail = b, a, c
	d, e, f := &link{}, &link{}, &link{}
	d.tail, e.tail = e, f

	tests := []struct {
		x        interface{}
		expected bool
	}{
		{x: a, expected: true},
		{x: b, expected: true},
		{x: c, expected: true},
		{x: d, expected: false},
		{x: e, expected: false},
		{x: f, expected: false},
	}

	for _, test := range tests {
		got := Cyclic(test.x)
		if got != test.expected {
			t.Errorf("%v unexpected result. expected: %v, but got: %v", test.x, test.expected, got)
		}
	}
}
