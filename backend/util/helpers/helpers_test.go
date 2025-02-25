package helpers

import (
	"testing"
)

func TestPtr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		val := 42
		ptr := Ptr(val)
		if *ptr != val {
			t.Errorf("expected %d, got %d", val, *ptr)
		}
	})

	t.Run("string", func(t *testing.T) {
		val := "hello"
		ptr := Ptr(val)
		if *ptr != val {
			t.Errorf("expected %s, got %s", val, *ptr)
		}
	})

	t.Run("struct", func(t *testing.T) {
		type sample struct {
			Field string
		}
		val := sample{Field: "test"}
		ptr := Ptr(val)
		if *ptr != val {
			t.Errorf("expected %+v, got %+v", val, *ptr)
		}
	})
}
