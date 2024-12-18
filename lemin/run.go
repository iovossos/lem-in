package lemin

import "fmt"

func Run(arg string) {

	totalAnts, rooms, start, end := readDataFromFile(arg)
	fmt.Println(totalAnts)

	calculateDistancesFromEnd(end)

	sortConnectedBySteps(rooms)

	virtualAnt := &Ant{
		name:     "Bob",
		location: start,
	}
	startingPaths := walkPath(virtualAnt, start, end, []*Room{start})

	allPathSets := findAllPathSets(startingPaths, start, end)

	for _, set := range allPathSets {
		for _, path := range set {
			for _, room := range path {
				fmt.Print(room.name + " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}

	// sets := findPaths(start, end)

	// for _, set := range sets {
	// 	for _, path := range set {
	// 		for _, room := range path {
	// 			fmt.Print(room.name + " ")
	// 		}
	// 		fmt.Println()
	// 	}
	// 	fmt.Println()
	// }

	//ants := spawnAnts(totalAnts, start, end)

	// //Debug ; Print paths
	// for _, path := range paths {
	// 	for _, room := range path {
	// 		fmt.Print(room.name + " ")
	// 	}
	// 	fmt.Println()
	// }

	// var turnsPerPath map[int]int
	// ants, turnsPerPath = assignPathsToAnts(ants, paths)

	// queues := makeQueues(ants, paths)

	// for _, queue := range queues {
	// 	for _, ant := range queue {
	// 		fmt.Print(ant.name)
	// 	}
	// 	fmt.Println()
	// }

	//	turns(queues, paths, turnsPerPath, end)

	// Debug: Print parsed data
	/*fmt.Println("Number of ants:", ants)
	for _, room := range rooms {
		fmt.Printf("Room %s (%d, %d): ", room.name, room.x, room.y)
		for _, connected := range room.connected {
			fmt.Printf("%s ", connected.name)
		}
		fmt.Printf(". %d steps from end\n", room.stepsToEnd)
	}*/

	//Debut: Print paths per ant
	// for _, ant := range ants {
	// 	fmt.Println(ant.name, ant.path)
	// }

	//startAnts(ants, start, end)
}
