package books_test

import (
	"testing"

	"github.com/vimond/books"
)

//nolint:dupl // They are essentially the same test
func TestMin(t *testing.T) {
	smallInt, largeInt := 2, 9
	outputInt := books.Min(smallInt, largeInt)
	outputIntOrder := books.Min(largeInt, smallInt) // passing in the opposite order

	if outputInt != smallInt || outputIntOrder != smallInt {
		t.Fatalf("unexpected output when passing ints")
	}

	smallFloat, largeFloat := 3.2, 6.7
	outputFloat := books.Min(smallFloat, largeFloat)
	outputFloatOrder := books.Min(largeFloat, smallFloat) // passing in the opposite order

	if outputFloat != smallFloat || outputFloatOrder != smallFloat {
		t.Fatalf("unexpected output when passing floats")
	}

	smallStr, largeStr := "aaa", "yyy"
	outputStr := books.Min(smallStr, largeStr)
	outputStrOrder := books.Min(largeStr, smallStr) // passing in the opposite order

	if outputStr != smallStr || outputStrOrder != smallStr {
		t.Fatalf("unexpected output when passing strings")
	}
}

//nolint:dupl // They are essentially the same test
func TestMax(t *testing.T) {
	smallInt, largeInt := 2, 9
	outputInt := books.Max(smallInt, largeInt)
	outputIntOrder := books.Max(largeInt, smallInt) // passing in the opposite order

	if outputInt != largeInt || outputIntOrder != largeInt {
		t.Fatalf("unexpected output when passing ints")
	}

	smallFloat, largeFloat := 3.2, 6.7
	outputFloat := books.Max(smallFloat, largeFloat)
	outputFloatOrder := books.Max(largeFloat, smallFloat) // passing in the opposite order

	if outputFloat != largeFloat || outputFloatOrder != largeFloat {
		t.Fatalf("unexpected output when passing floats")
	}

	smallStr, largeStr := "aaa", "yyy"
	outputStr := books.Max(smallStr, largeStr)
	outputStrOrder := books.Max(largeStr, smallStr) // passing in the opposite order

	if outputStr != largeStr || outputStrOrder != largeStr {
		t.Fatalf("unexpected output when passing strings")
	}
}
