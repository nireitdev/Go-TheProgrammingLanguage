package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string

	//Args[0] = nomnbre comando

	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " - "
	// }

	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}

	s = strings.Join(os.Args[1:], " ")
	fmt.Println(s)
}
