package lemin

import (
	"fmt"
	"strconv"
)

func spawnAnts(totalAnts int, start, end *Room) []*Ant {
	ants := []*Ant{}

	for i := 1; i <= totalAnts; i++ {
		ant := &Ant{
			name:     "L" + strconv.Itoa(i),
			location: start,
		}
		ants = append(ants, ant)
	}

	// //fmt.Print(ants)
	// turns(ants, start, end)
	return ants
}

func turns(ants []*Ant, start, end *Room) {

	aliveFound := false
	for _, a := range ants {
		if !a.isDead {
			aliveFound = true
			break
		}
	}
	if !aliveFound {
		return
	}
	antsMoved := []string{}
	for _, ant := range ants {
		if ant.isDead {
			continue
		}
		if !ant.path[0].hasAnt {

			ant.location.hasAnt = false
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

	// if !ant.active {
	// 	turns(ants, start, end)
	// }

	for _, a := range antsMoved {
		fmt.Print(a + " ")
	}
	fmt.Println()
	turns(ants, start, end)
}

func assignPathsToAnts(ants []*Ant, paths [][]*Room) []*Ant {

	turnsPerPath := make(map[int]int)
	for i, path := range paths {
		turnsPerPath[i] = len(path)
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
		turnsPerPath[bestPath]++

	}
	return ants
}
