package lemin

func Run(arg string) {

	totalAnts, rooms, start, end := readDataFromFile(arg)

	calculateDistancesFromEnd(end)

	sortConnectedBySteps(rooms)

	virtualAnt := &Ant{
		name:     "Bob",
		location: start,
	}

	startingPaths := findAllStartingPaths(virtualAnt, start, end, []*Room{})

	allPathSets := findAllPathSets(startingPaths, start, end)

	optimalSet := countTurns(totalAnts, allPathSets)

	ants := spawnAnts(totalAnts, start)

	var turnsPerPath map[int]int
	ants, turnsPerPath = assignPathsToAnts(ants, optimalSet)

	queues := makeQueues(ants, optimalSet)

	startAnts(queues, turnsPerPath, end)
}
