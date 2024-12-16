package main

import (
	"fmt"
	"strconv"
)

type Ant struct {
	name       string
	location   *Room
	active     bool
	isDead     bool
	movesCount int
}

func startAnts(totalAnts int, start, end *Room) {
	ants := []*Ant{}

	for i := 1; i <= totalAnts; i++ {
		ant := &Ant{
			name:     "L" + strconv.Itoa(i),
			location: start,
		}
		ants = append(ants, ant)
	}

	//fmt.Print(ants)
	turns(ants, start, end)
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
		for _, room := range ant.location.connected {
			if !room.hasAnt && room.stepsToEnd <= ant.location.stepsToEnd {
				ant.location.hasAnt = false
				ant.location = room
				ant.movesCount++
				antsMoved = append(antsMoved, (ant.name + "-" + room.name))
				if room == end {
					ant.isDead = true
					break
				}
				room.hasAnt = true
				ant.active = true
				break
			}
			// if !ant.active {
			// 	turns(ants, start, end)
			// }

		}
	}
	for _, a := range antsMoved {
		fmt.Print(a + " ")
	}
	fmt.Println()
	turns(ants, start, end)
}

func findPaths(start, end *Room) [][]*Room {

	var paths [][]*Room

	for range start.connected {
		virtualAnt := &Ant{
			name:     "Bob",
			location: start,
		}
		var path []*Room
		path = walkPath(virtualAnt, start, end, path)
		if path != nil {
			paths = append(paths, path)
		}
	}

	return paths
}

func walkPath(virtualAnt *Ant, start, end *Room, path []*Room) []*Room {

	for _, room := range virtualAnt.location.connected {

		if !room.hasAnt && room != start {
			virtualAnt.location = room
			newPath := append([]*Room(nil), path...)
			newPath = append(newPath, room)

			if room == end {

				return path
			}
			room.hasAnt = true
			nextPath := walkPath(virtualAnt, start, end, newPath)
			if nextPath != nil {
				return nextPath
			}

			// Backtrack by unmarking the room as occupied if no valid path was found
			room.hasAnt = false
		}
	}
	// If no path is found, return nil
	return nil
}
