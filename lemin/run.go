package lemin

func Run(arg string) error {

	totalAnts, start, end, err := readDataFromFile(arg)
	if err != nil {
		return err
	}

	allPaths, err := findAllPaths(start, end, start, []*Room{})
	if err != nil {
		return err
	}

	allPathSets := findAllPathSets(allPaths, end)

	optimalSet := countTurns(totalAnts, allPathSets)

	ants := spawnAnts(totalAnts, start)

	ants, turnsPerPath := assignPathsToAnts(ants, optimalSet)

	queues := makeQueues(ants, optimalSet)

	startAnts(queues, turnsPerPath, end)

	return nil
}
