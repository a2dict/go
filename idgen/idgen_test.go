package idgen

import (
	"fmt"
	"testing"
)

func TestIdgen(t *testing.T) {
	gen := New("order")
	for i := 0; i < 100; i++ {
		fmt.Println(gen.Next())
	}
}
