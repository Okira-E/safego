package safego_test

import (
	safego "github.com/Okira-E/safego/src"
	"strconv"
	"testing"
)

func TestOption_IsNone(t *testing.T) {
	noneVal := safego.None[int]()

	if !noneVal.IsNone() {
		t.Error("Expect None to be none, got: " + strconv.FormatBool(!noneVal.IsNone()))
	}
}

func TestOption_IsSome(t *testing.T) {
	someVal := safego.Some[int](1)

	if !someVal.IsSome() {
		t.Error("Expect Some to be some, got: " + strconv.FormatBool(!someVal.IsSome()))
	}
}

func TestOption_Unwrap(t *testing.T) {
	someVal := safego.Some[int](1)

	if someVal.Unwrap() != 1 {
		t.Error("Expect Some to be 1, got: " + strconv.Itoa(someVal.Unwrap()))
	}
}

func TestOption_UnwrapOr(t *testing.T) {
	someVal := safego.Some[int](1)

	if someVal.UnwrapOr(2) != 1 {
		t.Error("Expect Some to be 1, got: " + strconv.Itoa(someVal.UnwrapOr(2)))
	}

	noneVal := safego.None[int]()

	if noneVal.UnwrapOr(2) != 2 {
		t.Error("Expect None to be 2, got: " + strconv.Itoa(noneVal.UnwrapOr(2)))
	}
}

func TestOption_UnwrapOrElse(t *testing.T) {
	someVal := safego.Some[int](1)

	if someVal.UnwrapOrElse(func() int { return 2 }) != 1 {
		t.Error("Expect Some to be 1, got: " + strconv.Itoa(someVal.UnwrapOrElse(func() int { return 2 })))
	}

	noneVal := safego.None[int]()

	if noneVal.UnwrapOrElse(func() int { return 2 }) != 2 {
		t.Error("Expect None to be 2, got: " + strconv.Itoa(noneVal.UnwrapOrElse(func() int { return 2 })))
	}
}
