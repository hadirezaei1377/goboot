package mathematics_test

import "testing"

func TestPowerH(t *testing.M) {

	type test struct {
		Base           float64
		Power          float64
		ExpectedResult float64
	}

	testCases := []test{
		{Base: 2, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: -1, ExpectedResult: 0.5},
		{Base: 2, Power: 1, ExpectedResult: 2},
		{Base: 2, Power: -2, ExpectedResult: 0.25},
		{Base: 2, Power: 2, ExpectedResult: 4},
		{Base: 2, Power: 5, ExpectedResult: 32},
	}

	for _, c := range testCases {
		res := PowerH(c.Base, c.Power)
		if res != c.ExpectedResult {
			t.Errorf("expected: %f, result: %f\n", c.ExpectedResult, res)
		}
	}
}

// testing:
// go test /. mathematics
// output: ok

func TestPowerMath(t *testing.M) {

	type test struct {
		Base           float64
		Power          float64
		ExpectedResult float64
	}

	testCases := []test{
		{Base: 2, Power: 0, ExpectedResult: 1},
		{Base: 2, Power: -1, ExpectedResult: 0.5},
		{Base: 2, Power: 1, ExpectedResult: 2},
		{Base: 2, Power: -2, ExpectedResult: 0.25},
		{Base: 2, Power: 2, ExpectedResult: 4},
		{Base: 2, Power: 5, ExpectedResult: 32},
	}

	for _, c := range testCases {
		res := PowerMath(c.Base, c.Power)
		if res != c.ExpectedResult {
			t.Errorf("expected: %f, result: %f\n", c.ExpectedResult, res)
		}
	}
}

var result float64

func BenchmarkPowerH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := PowerH(2, 4)
		result = res
	}
}

// benchmarking:
// go test -bench=./mathematics

//output:
/*
your os is windows
your windows architecture is 64 bit                 ----> os considers int as int64
your package is mathematics
your os cpu is 2.5 GHZ
benchmark-PowerH-8     479866887   2.97 ns/op       ----> 8 is number of cpu cores,  479866887 is number of test cases used in benchmark 2.7 ns for each
pass
ok       2.1 s                                      ----> duration of benchmarking

*/

func BenchmarkPowerMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := PowerMath(2, 4)
		result = res
	}
}

// benchmarking:
// go test -bench=./mathematics

//output:
/*
your os is windows
your windows architecture is 64 bit                 ----> os considers int as int64
your package is mathematics
your os cpu is 2.5 GHZ
benchmark-PowerH-8      415584756   2.8 ns/op       ----> 8 is number of cpu cores,  479866887 is number of test cases used in benchmark 2.7 ns for each
benchmark-PowerMath     49238748    21.12 ns/op
pass
ok       3.2 s                                      ----> duration of benchmarking

*/

// number of testcases decreased cause of increasing ns/op

// consider memory (not just cpu)
// go test -bench=./mathematics

/*
your os is windows
your windows architecture is 64 bit
your package is mathematics
your os cpu is 2.5 GHZ
benchmark-PowerH-8      419187595   2.84 ns/op      0 B/op    0 allocs/op
benchmark-PowerMath     52710456    21.07 ns/op     0 B/op    0 allocs/op
pass
ok       3.25 s


0 B/op       is number of bytes which allocated to function per each operation
0 allocs/op  is number of allocating
*/
