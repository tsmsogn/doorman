package main

import (
	"fmt"
	"os"
)

func main() {
	char := Chair{x: 1, y: 20, z: 15}
	door := Door{x: 10, y: 1}

	if char.canThrough(door) {
		fmt.Println("通れます")
		os.Exit(0)
	}

	fmt.Println("通れません")
	os.Exit(1)
}
