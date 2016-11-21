package main

import (
	"fmt"
)

const (
	BITs = 10
	B    = iota << BITs
	KB   = B << BITs
	MB   = KB << BITs
	GB   = MB << BITs
	TB   = GB << BITs
	PB   = TB << BITs
)

func main() {
	fmt.Println(1 ^ 2)
	fmt.Println(!true)
	fmt.Println(!(10%3 == 0))
	fmt.Println(1 << 10)
	fmt.Println(1 << 10 >> 10)

	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 &^ 11)
	fmt.Println(6 * 11)

	a := 1
	// a = 0
	if a > 0 && (10/a) > 1 {
		fmt.Println("OK")
	}

	fmt.Println(B, KB, MB, GB, TB, PB)
}
