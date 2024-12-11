package main

import (
	"log"
	"strconv"
)

func checkDuplicateCoordinates(rooms map[string]*Room) {
	coordinates := make(map[string]bool)
	for roomName, room := range rooms {
		xy := strconv.Itoa(room.x) + strconv.Itoa(room.y)
		_, exists := coordinates[xy]
		if exists {
			log.Fatal("Room with duplicate coordinates:", roomName, xy)
		}
		coordinates[xy] = true
	}
}
