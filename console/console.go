package main

import "os"

func main() {
	if len(os.Args) > 1 {
		println("You executed", os.Args[2])
	} else {
		println("You didn't pass any argument")
	}
}
