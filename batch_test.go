package books_test

import (
	"errors"
	"testing"

	"github.com/vimond/books"
)

//nolint:funlen,gocognit // It doesn't make much sense to split up
func TestBatch(t *testing.T) {
	var (
		errEmptyInput   = errors.New("empty input")   //nolint:goerr113 // this is only needed for this test
		errInvalidInput = errors.New("invalid input") //nolint:goerr113 // this is only needed for this test
	)

	cases := []struct {
		name           string
		input          []int
		expectedOutput [][]int
		batchSize      int
		failOnErr      bool
		expectedErr    []error
	}{
		{
			name:           "Odd number of items",
			input:          []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			expectedOutput: [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}, {12}},
			batchSize:      2,
			failOnErr:      true,
			expectedErr:    []error{nil, nil, nil, nil, nil, nil, nil},
		},
		{
			name:           "Even number of items",
			input:          []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			expectedOutput: [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}, {10, 11}},
			batchSize:      2,
			failOnErr:      true,
			expectedErr:    []error{nil, nil, nil, nil, nil, nil},
		},
		{
			name:           "Wrong order failing on err",
			input:          []int{0, 1, 3, 2, 4, 5},
			expectedOutput: [][]int{{0, 1}, nil},
			batchSize:      2,
			failOnErr:      true,
			expectedErr:    []error{nil, errInvalidInput},
		},
		{
			name:           "Wrong order not failing on err",
			input:          []int{0, 1, 2, 0, 3, 4},
			expectedOutput: [][]int{{0, 1}, nil, {3, 4}},
			batchSize:      2,
			failOnErr:      false,
			expectedErr:    []error{nil, errInvalidInput, nil},
		},
		{
			name:           "Increase batch size",
			input:          []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			expectedOutput: [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9, 10, 11}, {12}},
			batchSize:      3,
			failOnErr:      true,
			expectedErr:    []error{nil, nil, nil, nil, nil},
		},
		{
			name:           "Batch size larger than input",
			input:          []int{0, 1, 2, 3},
			expectedOutput: [][]int{{0, 1, 2, 3}},
			batchSize:      10,
			failOnErr:      true,
			expectedErr:    []error{nil},
		},
		{
			name:           "Batch size larger than input with wrong order",
			input:          []int{0, 1, 5, 0},
			expectedOutput: [][]int{nil},
			batchSize:      10,
			failOnErr:      true,
			expectedErr:    []error{errInvalidInput},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var count int

			// Calling batch with a function iterating through the input and counting each valid input
			r, errs := books.Batch(tc.input, func(d []int) ([]int, error) {
				if len(d) == 0 {
					return nil, errEmptyInput
				}

				for i := range d {
					if d[i] != count {
						return nil, errInvalidInput // wrong order
					}

					count++
				}

				return d, nil
			}, tc.batchSize, tc.failOnErr)

			// Checking errors match
			if len(tc.expectedErr) != len(errs) {
				t.Fatalf("expected error (%+v) did not match returned error (%+v)", tc.expectedErr, errs)
			}

			// Checking that either the both errors are nil, or have the same
			for i := range errs {
				if !errors.Is(errs[i], tc.expectedErr[i]) {
					t.Fatalf("expected error (%+v) did not match returned error (%+v)", tc.expectedErr, errs)
				}
			}

			// Checking output matches
			if len(tc.expectedOutput) != len(r) {
				t.Fatalf("expected output (%+v) did not match returned output (%+v)", tc.expectedOutput, r)
			}

			for i := range r {
				if tc.expectedOutput[i] == nil && r[i] == nil {
					continue
				}

				if len(tc.expectedOutput[i]) != len(r[i]) {
					t.Fatalf("expected output (%+v) did not match returned output (%+v)", tc.expectedOutput, r)
				}

				for j := range r[i] {
					if tc.expectedOutput[i][j] != r[i][j] {
						t.Fatalf("expected output (%+v) did not match returned output (%+v)", tc.expectedOutput, r)
					}
				}
			}
		})
	}
}
