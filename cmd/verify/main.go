package main

import (
	"doorman/internal"
	"fmt"
	"os"
)

func main() {
	char := internal.Chair{X: 1, Y: 20, Z: 15}
	door := internal.Door{X: 10, Y: 1}

	if char.CanThrough(door) {
		fmt.Println("通れます")
		os.Exit(0)
	}

	fmt.Println("通れません")
	os.Exit(1)
}
