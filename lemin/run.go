package lemin

func Run(arg string) {

	totalAnts, rooms, start, end := readDataFromFile(arg)

	calculateDistancesFromEnd(end)

	sortConnectedBySteps(rooms)

	paths := findPaths(start, end)

	ants := spawnAnts(totalAnts, start, end)

	// //Debug ; Print paths
	// for _, path := range paths {
	// 	for _, room := range path {
	// 		fmt.Print(room.name + " ")
	// 	}
	// 	fmt.Println()
	// }

	ants = assignPathsToAnts(ants, paths)

	turns(ants, start, end)

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
