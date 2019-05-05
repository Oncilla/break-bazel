package breaker

import (
	"fmt"
	"testing"
)

func TestBSmash(t *testing.T) {
	if _, err := (Breaker{Smasher: smasher{}}).Break(); err != nil {
		t.Error("Smash should not fail")
	}
}

type smasher struct{}

func (smasher) Smash() (Broken, error) {
	fmt.Println("I love to break stuff")
	return true, nil
}
