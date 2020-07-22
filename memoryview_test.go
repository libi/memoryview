package memoryview

import (
	"fmt"
	"testing"
)

func TestMemView(t *testing.T) {
	a := 10
	fmt.Println(MemView(&a))

	b := int32(257)
	fmt.Println(MemView(&b))

	c := [3]int{1, 2, 3}
	fmt.Println(MemView(&c))

	d := c[2:3]
	fmt.Println(MemView(&d))

	e := "abcd"
	fmt.Println(MemView(&e))
}
