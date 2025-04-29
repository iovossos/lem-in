package funcs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var connections map[string][]string
var tunnels [][2]string
var totalAnts int
var start, end, mode string
var firstLine = true
var maxFlowPaths [][]string

// Edge represents a directed edge in the flow network.
type Edge struct {
	from     string
	to       string
	capacity int
	flow     int
	rev      *Edge
}

// AntAssignment stores which ant is assigned to which path and its order on that path.
type AntAssignment struct {
	antID     int // The global ant number (starting at 1)
	pathIndex int // Index into the paths slice (after sorting)
	order     int // The order of this ant on the path (1 means first ant on that path)
}

func ParseInput() error {

	if len(os.Args) != 2 {
		return errors.New("wrond input syntax")
	}

	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open input file:", err)
		return errors.New("")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if firstLine {
			totalAnts, err = strconv.Atoi(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid ant count %q: %v\n", line, err)
				return errors.New("")
			} else if totalAnts == 0 {
				fmt.Println("Zero ant count")
				return errors.New("")
			}
			firstLine = false
			continue
		}

		fmt.Printf("%s\n", line)

		switch line {
		case "##start":
			mode = "start"
			continue
		case "##end":
			mode = "end"
			continue
		}

		if fields := strings.Fields(line); len(fields) == 3 {

			if mode == "start" {
				start = fields[0]
			} else if mode == "end" {
				end = fields[0]
			}
			mode = ""
		}

		if parts := strings.Split(line, "-"); len(parts) == 2 {
			a, b := parts[0], parts[1]
			tunnels = append(tunnels, [2]string{a, b})
		}
	}

	if start == "" {
		fmt.Println("Start room not found")
		return errors.New("")
	}

	if end == "" {
		fmt.Println("End room not found")
		return errors.New("")
	}

	return nil

}
