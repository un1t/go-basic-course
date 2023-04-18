package foo_test

import (
	"testing"

	"github.com/un1t/example/foo"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b, result int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{5, 5, 10},
	}

	for _, tt := range tests {
		result := foo.Add(tt.a, tt.b)
		if result != tt.result {
			t.Errorf("Expected %v, got %v", tt.result, result)
		}
	}
}
