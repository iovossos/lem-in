package lemin

func Run(arg string) {

	totalAnts, start, end := readDataFromFile(arg)

	allPaths := findAllPaths(start, end, start, []*Room{})

	allPathSets := findAllPathSets(allPaths, end)

	optimalSet := countTurns(totalAnts, allPathSets)

	ants := spawnAnts(totalAnts, start)

	ants, turnsPerPath := assignPathsToAnts(ants, optimalSet)

	queues := makeQueues(ants, optimalSet)

	startAnts(queues, turnsPerPath, end)

}
