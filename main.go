package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	num       uint
	x         uint
	y         uint
	connected []*Room
	hasAnt    bool
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
	content := make([]byte, 1024*64) // Adjust buffer size as needed
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert to string and split into lines
	lines := strings.Split(string(content[:n]), "\n")

	// Separate rooms and links
	roomLines, linkLines := splitSections(lines)

	// Parse rooms and links
	rooms := parseRooms(roomLines)
	parseLinks(linkLines, rooms)

	// Debug: Print parsed data
	for _, room := range rooms {
		fmt.Printf("Room %d (%d, %d): ", room.num, room.x, room.y)
		for _, connected := range room.connected {
			fmt.Printf("%d ", connected.num)
		}
		fmt.Println()
	}
}

func splitSections(lines []string) ([]string, []string) {
	var roomLines, linkLines []string
	foundLinks := false

	for _, line := range lines {
		line = strings.TrimSpace(line) // Remove leading/trailing whitespace
		if line == "" || strings.HasPrefix(line, "#") {
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

func parseRooms(lines []string) map[uint]*Room {
	rooms := make(map[uint]*Room)

	for _, line := range lines {
		if line == "##start" || line == "##end" || line[0] == '#' {
			continue // Skip special commands or comments
		}

		// Split the line into parts (e.g., "1 23 3")
		parts := strings.Fields(line)
		if len(parts) != 3 {
			fmt.Println("Invalid room format:", line)
			continue
		}

		// Convert parts to the appropriate types
		num, _ := strconv.Atoi(parts[0])
		x, _ := strconv.Atoi(parts[1])
		y, _ := strconv.Atoi(parts[2])

		// Create a new Room and add it to the map
		room := &Room{
			num:       uint(num),
			x:         uint(x),
			y:         uint(y),
			connected: []*Room{},
			hasAnt:    false,
		}
		rooms[uint(num)] = room
	}

	return rooms
}

func parseLinks(lines []string, rooms map[uint]*Room) {
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

		// Convert room names to uint
		from, _ := strconv.Atoi(parts[0])
		to, _ := strconv.Atoi(parts[1])

		// Get the corresponding Room objects
		roomFrom, okFrom := rooms[uint(from)]
		roomTo, okTo := rooms[uint(to)]

		if !okFrom || !okTo {
			fmt.Println("Invalid link: unknown room(s)", line)
			continue
		}

		// Add the connection to both rooms
		roomFrom.connected = append(roomFrom.connected, roomTo)
		roomTo.connected = append(roomTo.connected, roomFrom)
	}
}
