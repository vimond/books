package books_test

import (
	"testing"

	"github.com/vimond/books"
)

func TestPtr(t *testing.T) {
	str := "string test"
	ps := books.Ptr(str)

	if ps == nil || *ps != str {
		t.Fatalf("string pointers did not match")
	}

	i := 2
	pi := books.Ptr(i)

	if pi == nil || *pi != i {
		t.Fatalf("int pointers did not match")
	}

	b := true
	pb := books.Ptr(b)

	if pb == nil || *pb != b {
		t.Fatalf("bool pointers did not match")
	}
}

func TestDeref(t *testing.T) {
	var nilStrPtr *string
	emptyStr := books.Deref(nilStrPtr)

	if emptyStr != "" {
		t.Fatalf("unexpected output for nil string")
	}

	var nilIntPtr *int
	zeroInt := books.Deref(nilIntPtr)

	if zeroInt != 0 {
		t.Fatalf("unexpected output for nil int")
	}

	str := "string test"
	derefString := books.Deref(&str)

	if str != derefString {
		t.Fatalf("unexpected string output")
	}

	i := 42
	derefInt := books.Deref(&i)

	if i != derefInt {
		t.Fatalf("unexpected int output")
	}

	b := true
	derefBool := books.Deref(&b)

	if b != derefBool {
		t.Fatalf("unexpected bool output")
	}
}
