package main

import (
	"doorman/internal"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	// Slice of chair width, height or depth
	CharSlice []*float64 `long:"chair" description:"Chair width, height or depth"`
	// Slice of door width or height
	DoorSlice []*float64 `long:"door" description:"Door width or height"`
}

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	var opts Options
	args, err := flags.ParseArgs(&opts, args)

	if err != nil {
		return 1
	}

	char := internal.Chair{
		X: *opts.CharSlice[0],
		Y: *opts.CharSlice[1],
		Z: *opts.CharSlice[2],
	}

	door := internal.Door{
		X: *opts.DoorSlice[0],
		Y: *opts.DoorSlice[1],
	}

	if char.CanThrough(door) {
		fmt.Println("通れます")
		return 0
	}

	fmt.Println("通れません")
	return 1
}
