package wordcount

import "testing"

// unit testing
func TestCountWords(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"Hello world", 2},
		{"Go is awesome!", 3},
		{"", 0},
		{"   ", 0},
		{"OneWord", 1},
		{"Hello, Go!", 2},
	}

	for _, tc := range testCases {
		result := CountWords(tc.input)
		if result != tc.expected {
			t.Errorf("CountWords(%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}

// benchmark testing
func BenchmarkCountWords(b *testing.B) {
	text := "Go is a statically typed, compiled programming language designed at Google."

	for i := 0; i < b.N; i++ {
		CountWords(text)
	}
}

// fuzz testing
func FuzzCountWords(f *testing.F) {
	testCases := []string{"Hello world", " ", "OneWord", "12345", "!@#$%^&*()", "Go\tis\tawesome"}

	for _, tc := range testCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, input string) {
		_ = CountWords(input)
	})
}
