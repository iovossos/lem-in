package lemin

func initializeTurnsMap(paths [][]*Room) map[int]int {
	turnsPerPath := make(map[int]int)
	for i, path := range paths {
		turnsPerPath[i] = len(path) - 1
	}
	return turnsPerPath
}

func findPathWithFewerTurns(turnsPerPath map[int]int) int {
	minTurns := turnsPerPath[0]
	bestPath := 0
	for pathIndex, turnsNeeded := range turnsPerPath {
		if turnsNeeded < minTurns {
			minTurns = turnsNeeded
			bestPath = pathIndex
		}
	}
	return bestPath
}

func findMaxTurnsNeeded(turnsPerPath map[int]int) int {
	maxTurns := turnsPerPath[0]
	for _, turnsNeeded := range turnsPerPath {
		if turnsNeeded > maxTurns {
			maxTurns = turnsNeeded
		}
	}
	return maxTurns
}
