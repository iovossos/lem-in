package main

import (
	"fmt"
	"lem-in/funcs"
)

func main() {

	funcs.ParseInput()
	funcs.BuildConnections()
	funcs.VertexDisjointPaths()
	simulationLines := funcs.OptimalAntDistribution()

	fmt.Println("[DEBUG]: simulationLines:", simulationLines)

	for _, line := range simulationLines {
		fmt.Println(line)
	}

}
