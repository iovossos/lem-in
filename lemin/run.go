package lemin

func Run(arg string) {

	totalAnts, rooms, start, end := readDataFromFile(arg)

	calculateDistancesFromEnd(end)

	sortConnectedBySteps(rooms)

	startingPaths := findAllStartingPaths(start, end, []*Room{})

	allPathSets := findAllPathSets(startingPaths, start, end)

	optimalSet := countTurns(totalAnts, allPathSets)

	ants := spawnAnts(totalAnts, start)

	ants, turnsPerPath := assignPathsToAnts(ants, optimalSet)

	queues := makeQueues(ants, optimalSet)

	startAnts(queues, turnsPerPath, end)

}
