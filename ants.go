package main

import "fmt"

type Ant struct {
	name     string
	location *Room
	active   bool
	isDead   bool
}

func startAnts(totalAnts rune, start, end *Room) {
	ants := []*Ant{}

	for i := '1'; i <= totalAnts; i++ {
		ant := &Ant{
			name:     "L" + string(i),
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
			if !room.hasAnt {
				ant.location.hasAnt = false
				ant.location = room
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
