package main

import (
	"bufio"
	"fmt"
	"lem-in/funcs"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Wrong input syntax")
		return
	}

	var tunnels [][2]string
	var connections map[string][]string
	var start, end, mode string
	var totalAnts int
	firstLine := true
	var maxFlowPaths [][]string

	inputFile := os.Args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Failed to open input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if firstLine {
			totalAnts, err = strconv.Atoi(line)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid ant count %q: %v\n", line, err)
				return
			} else if totalAnts == 0 {
				fmt.Println("Zero ant count")
				return
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
		return
	}

	if end == "" {
		fmt.Println("End room not found")
		return
	}

	fmt.Println()
	connections = funcs.BuildConnections(tunnels)
	maxFlowPaths, _ = funcs.VertexDisjointPaths(connections, start, end)
	simulationLines := funcs.OptimalAntDistribution(totalAnts, maxFlowPaths)

	for _, line := range simulationLines {
		fmt.Println(line)
	}

}
