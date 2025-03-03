package mathpack_test

import (
	"mytesting/mathpack"
	"testing"
)

//func TestAdd(t *testing.T) {
//	got := mathpack.Add(1, 1)
//	excepted := 3
//
//	if got != excepted {
//		t.Fail()
//	}
//}

type addTestCase struct {
	a, b, expected int
}

var addTestCases = []addTestCase{
	{1, 1, 2},
	{25, 25, 50},
	{2, 1, 3},
	{1, 10, 11},
}

func TestAdd(t *testing.T) {
	for _, tc := range addTestCases {
		got := mathpack.Add(tc.a, tc.b)

		if got != tc.expected {
			t.Errorf("Add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.expected)
		}
	}
}

//func FuzzTestAdd(f *testing.F) {
//	f.Fuzz(func(t *testing.T, a, b int) {
//		mathpack.Add(a, b)
//	})
//}

func AllocateMemory() []int {
	return make([]int, 1000000)
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//mathpack.Add(2, 3)
		_ = AllocateMemory()
	}

	//for _, v := range addTestCases {
	//	b.Run(fmt.Sprintf("%v %v %v", v.a, v.b, v.expected), func(b *testing.B) {
	//		for i := 0; i < b.N; i++ {
	//			mathpack.Add(v.a, v.b)
	//		}
	//	})
	//}
}

func TestFetchUser(t *testing.T) {
	user, err := mathpack.FetchUser(1)
	if err != nil || user != "Alice" {
		t.Errorf("Expected Alice, got %s", user)
	}

	_, err = mathpack.FetchUser(2)
	if err == nil {
		t.Errorf("Expected error for user 2, got nil")
	}
}
