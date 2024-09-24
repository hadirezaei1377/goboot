// math_test.go
package unit

import "testing"

func TestAddMultipleCases(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{10, -5, 5},
		{-3, 0, -6},
	}

	for _, tt := range tests {
		result := Add(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("For inputs %d and %d, expected %d but got %d", tt.a, tt.b, tt.expected, result)
		}
	}
}
