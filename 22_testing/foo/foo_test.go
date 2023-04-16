package foo_test

import (
	"testing"

	"example/foo"
)

func TestAdd(t *testing.T) {
	result := foo.Add(2, 3)
	if result != 5 {
		t.Logf("Expected 5, got %d", result)
		t.Fail()
	}

	t.Run("case 1", func(t *testing.T) {
		result = foo.Add(3, 7)
		if result != 10 {
			t.Logf("Expected 10, got %d", result)
			t.Fail()
		}
	})
}
