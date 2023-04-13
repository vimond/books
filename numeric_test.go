package books_test

import (
	"math"
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

func TestClamp(t *testing.T) {
	// x lower than min
	xInt, minInt, maxInt := 1, 5, 15
	outputInt := books.Clamp(xInt, minInt, maxInt)

	if outputInt != minInt {
		t.Fatalf("unexpected output when x is lower than min")
	}

	// x higher than max
	xInt2, minInt2, maxInt2 := 20, 5, 15
	outputInt2 := books.Clamp(xInt2, minInt2, maxInt2)

	if outputInt2 != maxInt2 {
		t.Fatalf("unexpected output when x is larger than max")
	}

	// min is lower than max
	xInt3, minInt3, maxInt3 := 20, 15, 5
	outputInt3 := books.Clamp(xInt3, minInt3, maxInt3)

	if outputInt3 != maxInt3 {
		t.Fatalf("unexpected output when min is lower than max")
	}

	// large bounds
	xInt4, minInt4, maxInt4 := 20, math.MinInt, math.MaxInt
	outputInt4 := books.Clamp(xInt4, minInt4, maxInt4)

	if outputInt4 != xInt4 {
		t.Fatalf("unexpected output when min is lower than max")
	}

	// float is between min and max
	xFloat, minFloat, maxFloat := 4.5, 3.2, 6.7
	outputFloat := books.Clamp(xFloat, minFloat, maxFloat)

	if outputFloat != xFloat {
		t.Fatalf("unexpected output when passing floats")
	}

	xStr, minStr, maxStr := "bbb", "aaa", "yyy"
	outputStr := books.Clamp(xStr, minStr, maxStr)

	if outputStr != xStr {
		t.Fatalf("unexpected output when passing strings")
	}
}
