package main

import (
	"fmt"
	"lemin/lemin"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [mazename.txt]")
		return
	}
	lemin.Run(os.Args[1])

}
