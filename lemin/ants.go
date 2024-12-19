package lemin

import (
	"fmt"
	"strconv"
)

// Initialize a slice of ants with location: start
func spawnAnts(totalAnts int, start *Room) []*Ant {
	ants := []*Ant{}

	for i := 1; i <= totalAnts; i++ {
		ant := &Ant{
			name:     "L" + strconv.Itoa(i),
			location: start,
		}
		ants = append(ants, ant)
	}

	return ants
}

// Makes queues with the paths each ants takes
func makeQueues(ants []*Ant, paths [][]*Room) [][]*Ant {
	queues := make([][]*Ant, len(paths))
	for _, ant := range ants {
		queues[ant.pathIndex] = append(queues[ant.pathIndex], ant)
	}
	return queues
}

// Sending ants turn by turn.
func startAnts(queues [][]*Ant, turnsPerPath map[int]int, end *Room) {
	maxTurns := turnsPerPath[0]

	for _, turnsNeeded := range turnsPerPath {
		if turnsNeeded > maxTurns {
			maxTurns = turnsNeeded
		}
	}
	for i := 0; i < maxTurns; i++ {
		antsMoved := []string{}
		for _, path := range queues {
			for a, ant := range path {
				if a <= i && !ant.isDead {
					ant.location = ant.path[0]
					ant.path = ant.path[1:]
					antsMoved = append(antsMoved, (ant.name + "-" + ant.location.name))
					if ant.location != end {
						ant.location.hasAnt = true
					} else {
						ant.isDead = true
					}
				}
			}
		}

		for _, a := range antsMoved {
			fmt.Print(a + " ")
		}
		fmt.Println()
	}
}

// Assigns paths to ants based on the minimum number of turns needed to reach the end room.
func assignPathsToAnts(ants []*Ant, paths [][]*Room) ([]*Ant, map[int]int) {

	turnsPerPath := make(map[int]int)
	for i, path := range paths {
		turnsPerPath[i] = len(path) - 1
	}

	for _, ant := range ants {
		minTurns := turnsPerPath[0]
		bestPath := 0
		for pathIndex, turnsNeeded := range turnsPerPath {
			if turnsNeeded < minTurns {
				minTurns = turnsNeeded
				bestPath = pathIndex
			}
		}
		ant.path = paths[bestPath]
		ant.pathIndex = bestPath
		turnsPerPath[bestPath]++

	}
	return ants, turnsPerPath
}
