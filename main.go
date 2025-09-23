package main

import (
	"fmt"
	"lemin/lemin"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [mazename.txt]")
		return
	}
	err := lemin.Run(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

}
