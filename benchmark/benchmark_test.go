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

func TestDayOfWeek(t *testing.T) {
	type test struct {
		Input          int
		ExpectedResult string
	}

	var testCases = []test{ // pass data to it
		{Input: 1, ExpectedResult: "شنبه"},
		{Input: 2, ExpectedResult: "یکشنبه"},
		{Input: 3, ExpectedResult: "دوشنبه"},
		{Input: 4, ExpectedResult: "سه شنبه"},
		{Input: 5, ExpectedResult: "چهارشنبه"},
		{Input: 6, ExpectedResult: "پنجشنبه"},
		{Input: 7, ExpectedResult: "جمعه"},
		{Input: 8, ExpectedResult: ""},   // corner case
		{Input: 0, ExpectedResult: ""},   // corner case
		{Input: -15, ExpectedResult: ""}, // corner case
	}

	for _, c := range testCases {
		result := dayOfWeek(s.Input)
		if result != c.ExpectedResult {
			t.Errorf("expested: %s, result: %s\n", c, c.ExpectedResult, result)
		}
	}

}
