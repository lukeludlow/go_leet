package goleet

import "testing"

func TestReverse(t *testing.T) {
	in := []int{
		123,
		-123,
		120,
		1534236469,
	}
	expected := []int{
		321,
		-321,
		21,
		0, // detect 32 bit integer overflow
	}
	numCases := 4
	for i := 0; i < numCases; i++ {
		actual := reverse(in[i])
		if actual != expected[i] {
			t.Fatalf("case %d fails: %v\n", i, actual)
		}
	}
}
