package example

import (
	"github.com/ooxi/go-optional/optional"
	"testing"
)

func TestUsage(t *testing.T) {
	optionalValue := optional.Empty[int]()

	if !optionalValue.IsEmpty() {
		t.Fatalf("`optionalValue' must be empty\n")
	}

	optionalValue = optional.Of[int](42)

	if !optionalValue.IsPresent() {
		t.Fatalf("`optionalValue' must be present\n")
	}
	if 42 != optionalValue.OrElseThrow() {
		t.Fatalf("Wrong optional value `%d'\n", optionalValue.OrElseThrow())
	}
}
