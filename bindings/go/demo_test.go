package demo

import (
	"testing"
)

func TestDemoAdd(t *testing.T) {
	if Add(1, 3) != 4 {
		t.Fatalf(`TestDemoAdd Failed`)
	}
}
