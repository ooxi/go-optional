// SPDX-License-Identifier: Zlib
package optional

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	emptyInt := Empty[int]()

	if !emptyInt.IsEmpty() {
		t.Fatalf("`emptyInt' must be empty\n")
	}
	if emptyInt.IsPresent() {
		t.Fatalf("`emptyInt' must not be present\n")
	}
}

func TestOf(t *testing.T) {
	presentInt := Of[int](3)

	if presentInt.IsEmpty() {
		t.Fatalf("`presentInt' must be not be empty\n")
	}
	if !presentInt.IsPresent() {
		t.Fatalf("`presentInt' must be present")
	}
}

func TestOfNilable(t *testing.T) {
	emptyInt := OfNilable[int](nil)

	if !emptyInt.IsEmpty() {
		t.Fatalf("`emptyInt' must be empty\n")
	}
	if emptyInt.IsPresent() {
		t.Fatalf("`emptyInt' must not be present\n")
	}

	value := 3
	presentInt := OfNilable(&value)

	if presentInt.IsEmpty() {
		t.Fatalf("`presentInt' must be not be empty\n")
	}
	if !presentInt.IsPresent() {
		t.Fatalf("`presentInt' must be present")
	}
}

func TestOrElsePresent(t *testing.T) {
	presentInt := Of[int](42)

	if value := presentInt.OrElse(3); 42 != value {
		t.Fatalf("Wrong value, expected `42' but got `%d'\n", value)
	}
}

func TestOrElseEmpty(t *testing.T) {
	emptyInt := Empty[int]()

	if value := emptyInt.OrElse(3); 3 != value {
		t.Fatalf("Wrong value, expected `3' but got `%d'\n", value)
	}
}

func TestOrElseThrowPresent(t *testing.T) {
	presentInt := Of[int](42)

	if 42 != presentInt.OrElseThrow() {
		t.Fatalf("Wrong value, expected 3 but got %d\n", presentInt.OrElseThrow())
	}
}

func TestOrElseThrowEmpty(t *testing.T) {
	emptyInt := Empty[int]()

	defer func() {
		if err := recover(); NoSuchElementError != err {
			t.Fatalf("Expected `NoSuchElementError' but got `%v'", err)
		}
	}()

	emptyInt.OrElseThrow()
	t.Fatalf("Expected panic\n")
}
