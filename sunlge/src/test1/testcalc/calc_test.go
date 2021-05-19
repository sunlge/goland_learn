package testcalc

import (
	"sunlge/src/calc"
	"testing"
)

func TestAdd(t *testing.T) {
	if 4 != calc.Add(1, 3 ) {
		t.Errorf("1 + 3 != 2")
	}
}
