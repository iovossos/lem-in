package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	name       string
	x          int
	y          int
	connected  []*Room
	hasAnt     bool
	stepsToEnd int
}

func main() {
	// Open the file
	file, err := os.Open("test1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Read file contents
	content, err := os.ReadFile("test1.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Convert to string and split into lines
	lines := strings.Split(string(content), "\n")

	ants, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal("Invalid number of ants:", err)
	}
	if ants <= 0 {
		log.Fatal("Ant number must be higher than 0.")
	}

	// Separate rooms and links
	roomLines, linkLines := splitSections(lines[1:])

	// Parse rooms and links
	rooms, start, end := parseRooms(roomLines)
	//fmt.Println("Start room:", start.name, ", End room:", end.name)
	checkDuplicateCoordinates(rooms)

	parseLinks(linkLines, rooms)

	calculateDistancesFromEnd(end)

	sortPaths(rooms)

	// Debug: Print parsed data
	/*fmt.Println("Number of ants:", ants)
	for _, room := range rooms {
		fmt.Printf("Room %s (%d, %d): ", room.name, room.x, room.y)
		for _, connected := range room.connected {
			fmt.Printf("%s ", connected.name)
		}
		fmt.Printf(". %d steps from end\n", room.stepsToEnd)
	}*/

	startAnts(rune(ants+'0'), start, end)
}

func splitSections(lines []string) ([]string, []string) {
	var roomLines, linkLines []string
	foundLinks := false

	for _, line := range lines {
		line = strings.TrimSpace(line) // Remove leading/trailing whitespace
		if line == "" {
			// Skip empty lines or comments
			continue
		}

		if strings.Contains(line, "-") {
			foundLinks = true
		}

		if foundLinks {
			linkLines = append(linkLines, line)
		} else {
			roomLines = append(roomLines, line)
		}
	}

	return roomLines, linkLines
}

func parseRooms(lines []string) (map[string]*Room, *Room, *Room) {
	rooms := make(map[string]*Room)
	startFound := false
	endFound := false
	var start, end *Room
	for _, line := range lines {
		if line == "##start" {
			startFound = true
			continue
		} else if line == "##end" {
			endFound = true
			continue
		} else if line[0] == '#' {
			continue
		}

		// Split the line into parts (e.g., "1 23 3")
		parts := strings.Fields(line)
		if len(parts) != 3 {
			log.Fatal("Invalid room format:", line)
		}

		// Convert parts to the appropriate types
		name := parts[0]
		if strings.HasPrefix(name, "L") {
			log.Fatal("Invalid room name:", name)
		}
		x, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("Invalid room coordinate:", x)
		}
		y, err := strconv.Atoi(parts[2])
		if err != nil {
			log.Fatal("Invalid room coordinate:", line, y)
		}

		if x < 0 || y < 0 {
			log.Fatal("Room coordinates cannot be negative numbers:", line)
		}

		// Create a new Room and add it to the map
		room := &Room{
			name:      name,
			x:         x,
			y:         y,
			connected: []*Room{},
			hasAnt:    false,
		}
		if startFound {
			start = room
			startFound = false
		} else if endFound {
			end = room
			endFound = false
		}
		_, exists := rooms[name]
		if exists {
			log.Fatal("Duplicate room name:", name)
		}
		rooms[name] = room
	}

	// If start or end is not found, log an error
	if start == nil {
		log.Fatal("Start room not defined")
	}
	if end == nil {
		log.Fatal("End room not defined")
	}

	return rooms, start, end
}

func parseLinks(lines []string, rooms map[string]*Room) {
	for _, line := range lines {
		if line[0] == '#' {
			continue // Skip comments
		}

		// Split the link into two room names (e.g., "0-4")
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			fmt.Println("Invalid link format:", line)
			continue
		}

		// Get the corresponding Room objects
		roomFrom, okFrom := rooms[parts[0]]
		roomTo, okTo := rooms[parts[1]]

		// Check if the rooms exist
		if !okFrom || !okTo {
			log.Fatal("Invalid link: unknown room(s)", line)
		}

		// Checks if the rooms are the same room
		if roomFrom == roomTo {
			log.Fatal("Cannot have a link to the same room:", line)
		}

		// Add the connection to both rooms
		roomFrom.connected = append(roomFrom.connected, roomTo)
		roomTo.connected = append(roomTo.connected, roomFrom)
	}
}

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

func sortPaths(rooms map[string]*Room) {
	for _, r := range rooms {
		sort.Slice(r.connected, func(i, j int) bool {
			return r.connected[i].stepsToEnd < r.connected[j].stepsToEnd
		})
	}
}
