package idgen

import (
	"fmt"
	"testing"
)

func TestIdgen(t *testing.T) {
	gen := New([]byte("od"))
	for i := 0; i < 100; i++ {
		fmt.Println(gen.Next())
	}
}
