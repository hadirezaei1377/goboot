package benchmark_test

import "testing"

func TestSquare(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := square(i)
		if result/i != i {
			t.Errorf("expected: %d", "result: %d\n", i*i, result)
		}
	}
}
