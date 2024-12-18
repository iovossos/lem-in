package lemin

import "sort"

func calculateDistancesFromEnd(end *Room) {
	// Initialize BFS queue and visited set
	queue := []*Room{end}           // Start BFS from the end room
	visited := make(map[*Room]bool) // Track visited rooms
	visited[end] = true
	end.stepsToEnd = 0 // End room has 0 steps to itself

	// BFS Loop
	for len(queue) > 0 {
		// Dequeue the first room
		currentRoom := queue[0]
		queue = queue[1:]

		// Traverse all connected rooms (neighbors)
		for _, neighbor := range currentRoom.connected {
			if !visited[neighbor] {
				// Mark the neighbor as visited
				visited[neighbor] = true

				// Set the steps to the end for this neighbor
				neighbor.stepsToEnd = currentRoom.stepsToEnd + 1

				// Enqueue the neighbor for further exploration
				queue = append(queue, neighbor)
			}
		}
	}
}

func sortConnectedBySteps(rooms map[string]*Room) {
	for _, r := range rooms {
		sort.Slice(r.connected, func(i, j int) bool {
			return r.connected[i].stepsToEnd < r.connected[j].stepsToEnd
		})
	}
}

// func findAllStartingPaths(current, end *Room) [][]*Room {
// 	start := current
// 	var startingPaths [][]*Room
// 	var startingPath []*Room
// 	for _, connected := range current.connected {
// 		startingPath = append(startingPath, connected)
// 		current = connected
// 		findAllStartingPaths(current, end)
// 	}

// }

func findAllPathSets(startingPaths [][]*Room, start, end *Room) [][][]*Room {
	var sets [][][]*Room

	for _, startingPath := range startingPaths {
		var set [][]*Room
		set = append(set, startingPath)
		for _, room := range startingPath {

			room.visited = true

		}

		for range start.connected {
			virtualAnt := &Ant{
				name:     "Bob",
				location: start,
			}
			var path []*Room
			path = walkNonOverlappingPaths(virtualAnt, start, end, path)
			if path != nil {
				set = append(set, path)

			}

		}
		sets = append(sets, set)

		for _, path := range set {
			for _, room := range path {
				room.visited = false
			}
		}

	}

	return sets
}

func walkPath(virtualAnt *Ant, start, end *Room, path []*Room) [][]*Room {
	var allPaths [][]*Room
	for _, room := range virtualAnt.location.connected {

		if !room.visited && room != start {
			if end.visited {
				end.visited = false
			}

			virtualAnt.location = room
			room.visited = true

			newPath := append([]*Room(nil), path...) // Make a copy of the current path
			newPath = append(newPath, room)          // Add the current room to the path

			if room == end {
				allPaths = append(allPaths, newPath)

			} else {

				paths := walkPath(virtualAnt, start, end, newPath)
				allPaths = append(allPaths, paths...)
			}

			// Backtrack by unmarking the room as visited if no valid path was found
			room.visited = false
		}
	}
	return allPaths
}

func walkNonOverlappingPaths(virtualAnt *Ant, start, end *Room, path []*Room) []*Room {

	for _, room := range virtualAnt.location.connected {

		if !room.visited && room != start {
			if end.visited {
				end.visited = false
			}
			virtualAnt.location = room
			newPath := append([]*Room(nil), path...) // Make a copy of the current path
			newPath = append(newPath, room)          // Add the current room to the path
			room.visited = true
			if room == end {

				return newPath
			}

			nextPath := walkNonOverlappingPaths(virtualAnt, start, end, newPath)
			if nextPath != nil {
				return nextPath
			}

			// Backtrack by unmarking the room as visited if no valid path was found
			room.visited = false
		}
	}

	// If no path is found, return nil
	return nil
}
