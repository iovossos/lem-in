package main

import (
	"lemin/lemin"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	lemin.Run(os.Args[1])

}
