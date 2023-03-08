package books

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func TestRemoveDupes(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "consecutive dupes",
			input:    []int{1, 2, 2, 2, 2, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "intermittent dupes",
			input:    []int{1, 2, 2, 3, 3, 4, 5, 6, 7, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "disordered dupes",
			input:    []int{1, 2, 3, 1, 4, 5, 1, 2, 1, 4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range cases {
		out := RemoveDupes(tc.input)
		if len(out) != len(tc.expected) {
			t.Fatalf("%s failed: output did not match length of expected output: (out) %+v != %+v (expected)", tc.name, out, tc.expected)
		}
		for i := range out {
			if out[i] != tc.expected[i] {
				t.Fatalf("%s failed: items in output did not match expected output: (out) %+v != %+v (expected)", tc.name, out, tc.expected)
			}
		}
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		name         string
		inputSlice   []int
		inputElement int
		expected     bool
	}{
		{
			name:         "start",
			inputSlice:   []int{1, 2, 3, 4, 5, 6, 7},
			inputElement: 1,
			expected:     true,
		},
		{
			name:         "end",
			inputSlice:   []int{1, 2, 3, 4, 5, 6, 7},
			inputElement: 7,
			expected:     true,
		},
		{
			name:         "middle",
			inputSlice:   []int{1, 2, 3, 4, 5, 6, 7},
			inputElement: 4,
			expected:     true,
		},
		{
			name:         "missing",
			inputSlice:   []int{1, 2, 3, 4, 5, 6, 7},
			inputElement: 0,
			expected:     false,
		},
		{
			name:         "missing negative",
			inputSlice:   []int{1, 2, 3, 4, 5, 6, 7},
			inputElement: -1,
			expected:     false,
		},
	}

	for _, tc := range cases {
		out := Contains(tc.inputSlice, tc.inputElement)
		if out != tc.expected {
			t.Fatalf("%s failed: output did not expected output: (out) %+v != %+v (expected)", tc.name, out, tc.expected)
		}
	}
}

func TestReverseSlice(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "ordered odd number of elements",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:     "ordered even number of elements",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: []int{8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:     "disordered",
			input:    []int{96, -27349, 0, 1324143, 4, -234222, 23424, 84593, 3424, 7483, -24352},
			expected: []int{-24352, 7483, 3424, 84593, 23424, -234222, 4, 1324143, 0, -27349, 96},
		},
	}

	for _, tc := range cases {
		out := ReverseSlice(tc.input)
		if len(out) != len(tc.expected) {
			t.Fatalf("%s failed: output did not match length of expected output: (out) %+v != %+v (expected)", tc.name, out, tc.expected)
		}
		for i := range out {
			if out[i] != tc.expected[i] {
				t.Fatalf("%s failed: items in output did not match expected output: (out) %+v != %+v (expected)", tc.name, out, tc.expected)
			}
		}
	}
}

func TestFilterSlice(t *testing.T) {
	cases := []struct {
		name     string
		f        func(string) bool
		input    []string
		expected []string
	}{
		{
			name: "Match string",
			f: func(s string) bool {
				return s == "keep"
			},
			input:    []string{"keep", "remove", "keep", "keep", "remove2", "not-keep"},
			expected: []string{"keep", "keep", "keep"},
		},
		{
			name: "Even number",
			f: func(s string) bool {
				i, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return false
				}
				return i%2 == 0
			},
			input:    []string{"1", "2", "6", "not an int", "5", "odd", "0", "-64", "string"},
			expected: []string{"2", "6", "0", "-64"},
		},
	}

	for _, tc := range cases {
		out := FilterSlice(tc.input, tc.f)
		if len(out) != len(tc.expected) {
			t.Fatalf("%s failed: output did not match length of expected output: (out) %+v != %+v (expected)", tc.name, out, tc.expected)
		}
		for i := range out {
			if out[i] != tc.expected[i] {
				t.Fatalf("%s failed: items in output did not match expected output: (out) %+v != %+v (expected)", tc.name, out, tc.expected)
			}
		}
	}
}

// nolint:gosec // this function is only used for testing and does not require cryptographic security
func BenchmarkEqualSlices(b *testing.B) {
	sl1 := make([]int, 100)
	sl2 := make([]int, 100)
	for i := 0; i < 100; i++ {
		sl1[i] = rand.Intn(1000000)
		sl2[i] = rand.Intn(1000000)
	}

	for i := 0; i < b.N; i++ {
		EqualSlices(sl1, sl2)
	}
}

// nolint:gosec // this function is only used for testing and does not require cryptographic security
func BenchmarkDeepEqualSlices(b *testing.B) {
	sl1 := make([]int, 100)
	sl2 := make([]int, 100)
	for i := 0; i < 100; i++ {
		sl1[i] = rand.Intn(1000000)
		sl2[i] = rand.Intn(1000000)
	}

	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(sl1, sl2)
	}
}
