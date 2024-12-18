package lemin

func Run(arg string) {

	totalAnts, rooms, start, end := readDataFromFile(arg)

	calculateDistancesFromEnd(end)

	sortConnectedBySteps(rooms)

	virtualAnt := &Ant{
		name:     "Bob",
		location: start,
	}
	startingPaths := walkPath(virtualAnt, start, end, []*Room{})

	allPathSets := findAllPathSets(startingPaths, start, end)

	optimalPaths := countTurns(totalAnts, allPathSets)

	ants := spawnAnts(totalAnts, start)

	var turnsPerPath map[int]int
	ants, turnsPerPath = assignPathsToAnts(ants, optimalPaths)

	queues := makeQueues(ants, optimalPaths)

	turns(queues, turnsPerPath, end)

}
