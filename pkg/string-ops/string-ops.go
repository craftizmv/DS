package string_ops

import "fmt"

func TestString() {
	s := "hello, world"
	for i := 0; i < len(s); i++ {
		b := s[i] // byte at position i
		fmt.Printf("Byte: %c at index %d\n", b, i)
	}

}
