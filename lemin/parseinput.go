package lemin

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Validate file, parse ant number, room map, start, end, print everything.
func readDataFromFile(filename string) (ants int, start, end *Room, err error) {
	// Read file contents
	content, err := os.ReadFile("mazes/" + filename)
	if err != nil {
		return 0, nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	// Convert to string and split into lines.
	lines := strings.Split(string(content), "\n")

	//Get number of ants
	lines[0] = strings.TrimSpace(lines[0])
	ants, err = strconv.Atoi(lines[0])
	if err != nil {
		return 0, nil, nil, fmt.Errorf("invalid number of ants: %v", err)
	}
	if ants <= 0 {
		return 0, nil, nil, errors.New("ant number must be higher than 0")
	}

	// Separate rooms and links.
	roomLines, linkLines := splitSections(lines[1:])

	// Parse rooms and links
	var rooms map[string]*Room
	rooms, start, end, err = parseRooms(roomLines)
	if err != nil {
		return 0, nil, nil, err
	}

	err = checkDuplicateCoordinates(rooms)
	if err != nil {
		return 0, nil, nil, err
	}

	err = parseLinks(linkLines, rooms)
	if err != nil {
		return 0, nil, nil, err
	}

	//Print the file
	for i := range lines {
		if lines[i] != "" {
			fmt.Println(lines[i])
		}
	}
	fmt.Println()
	return
}

func checkDuplicateCoordinates(rooms map[string]*Room) error {
	coordinates := make(map[string]bool)
	for roomName, room := range rooms {
		xy := strconv.Itoa(room.x) + "-" + strconv.Itoa(room.y)
		_, exists := coordinates[xy]
		if exists {
			return fmt.Errorf("room %v has duplicate coordinates: %v", roomName, xy)
		}
		coordinates[xy] = true
	}
	return nil
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

func parseRooms(lines []string) (map[string]*Room, *Room, *Room, error) {
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
			return rooms, nil, nil, fmt.Errorf("invalid room format: %v", line)
		}

		// Convert parts to the appropriate types
		name := parts[0]
		if strings.HasPrefix(name, "L") {
			return rooms, nil, nil, fmt.Errorf("invalid room name: %v", name)
		}
		x, err := strconv.Atoi(parts[1])
		if err != nil {
			return rooms, nil, nil, fmt.Errorf("invalid room coordinate: %v", x)
		}
		y, err := strconv.Atoi(parts[2])
		if err != nil {
			return rooms, nil, nil, fmt.Errorf("invalid room coordinate: %v, %v", line, y)
		}

		if x < 0 || y < 0 {
			return rooms, nil, nil, fmt.Errorf("room coordinates cannot be negative numbers: %v", line)
		}

		// Create a new Room and add it to the map
		room := &Room{
			name:      name,
			x:         x,
			y:         y,
			connected: []*Room{},
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
			return rooms, nil, nil, fmt.Errorf("duplicate room name: %v", name)
		}
		rooms[name] = room
	}

	// If start or end is not found, log an error
	if start == nil {
		return rooms, nil, nil, errors.New("start room not defined")
	}
	if end == nil {
		return rooms, nil, nil, errors.New("end room not defined")
	}
	return rooms, start, end, nil
}

func parseLinks(lines []string, rooms map[string]*Room) error {
	for _, line := range lines {
		if line[0] == '#' {
			continue // Skip comments
		}

		// Split the link into two room names (e.g., "0-4")
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return fmt.Errorf("invalid link format: %v", line)
		}

		// Get the corresponding Room objects
		roomFrom, okFrom := rooms[parts[0]]
		roomTo, okTo := rooms[parts[1]]

		// Check if the rooms exist
		if !okFrom || !okTo {
			return fmt.Errorf("invalid link: unknown room(s): %v", line)
		}

		// Checks if the rooms are the same room
		if roomFrom == roomTo {
			return fmt.Errorf("cannot have a link to the same room: %v", line)
		}

		// Add the connection to both rooms
		roomFrom.connected = append(roomFrom.connected, roomTo)
		roomTo.connected = append(roomTo.connected, roomFrom)
	}
	return nil
}
