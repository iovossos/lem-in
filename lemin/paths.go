package lemin

import "errors"

// Find every possible unique path from the start room to the end room.
func findAllPaths(start, end, current *Room, path []*Room) ([][]*Room, error) {
	var allPaths [][]*Room

	for _, room := range current.connected {
		if !room.visited && room != start {
			if end.visited {
				end.visited = false
			}

			room.visited = true

			newPath := append([]*Room(nil), path...) // Make a copy of the current path
			newPath = append(newPath, room)          // Add the current room to the path

			if room == end {
				allPaths = append(allPaths, newPath)
			} else {
				paths, _ := findAllPaths(start, end, room, newPath)
				allPaths = append(allPaths, paths...)
			}

			// Backtrack by unmarking the room as visited
			room.visited = false
		}
	}
	if len(allPaths) == 0 {
		return allPaths, errors.New("No valid path between Start and End")
	}

	return allPaths, nil
}

// findAllPathSets returns all possible combinations of non-overlapping paths
// that can be formed from the identified paths
func findAllPathSets(allPaths [][]*Room, end *Room) [][][]*Room {
	var allSets [][][]*Room

	// Helper function to check if a path overlaps with any path in the current set
	hasOverlap := func(path []*Room, currentSet [][]*Room) bool {
		visited := make(map[*Room]bool)
		// Mark all rooms in current set as visited
		for _, existingPath := range currentSet {
			for _, room := range existingPath {
				if room != end {
					visited[room] = true
				}
			}
		}
		// Check if new path visits any already visited room
		for _, room := range path {
			if visited[room] && room != end {
				return true
			}
		}
		return false
	}

	// Recursive helper function to build combinations
	var buildSets func(remainingPaths [][]*Room, currentSet [][]*Room)
	buildSets = func(remainingPaths [][]*Room, currentSet [][]*Room) {
		// Add current set to results if it's not empty
		if len(currentSet) > 0 {
			setsCopy := make([][]*Room, len(currentSet))
			copy(setsCopy, currentSet)
			allSets = append(allSets, setsCopy)
		}

		// Try adding each remaining path to the current set
		for i, path := range remainingPaths {
			if !hasOverlap(path, currentSet) {
				// Create new remaining paths slice without current path
				newRemaining := make([][]*Room, 0)
				newRemaining = append(newRemaining, remainingPaths[:i]...)
				newRemaining = append(newRemaining, remainingPaths[i+1:]...)

				// Add current path to set and recurse
				newSet := append(currentSet, path)
				buildSets(newRemaining, newSet)
			}
		}
	}

	// Start the recursive process
	buildSets(allPaths, [][]*Room{})

	return allSets
}

// Count the turns needed for each set based on the number of ants & return the optimal set.
func countTurns(totalAnts int, sets [][][]*Room) [][]*Room {
	turnsPerSet := make(map[int]int)

	for s, set := range sets {

		turnsPerPath := initializeTurnsMap(set)

		for a := 0; a < totalAnts; a++ {
			bestPath := findPathWithFewerTurns(turnsPerPath)

			turnsPerPath[bestPath]++

		}
		turnsPerSet[s] = findMaxTurnsNeeded(turnsPerPath)
	}
	minTurnsNeeded := turnsPerSet[0]
	optimalPath := 0
	for pathIndex, turns := range turnsPerSet {
		if turns < minTurnsNeeded {
			minTurnsNeeded = turns
			optimalPath = pathIndex
		}
	}

	return sets[optimalPath]
}
