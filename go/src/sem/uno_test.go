package sem

import (
	"testing"
)

const I = -1

func TestUno(t *testing.T) {
	s := NewSem(I)

	if s == nil {
		t.Error()
	}
}
