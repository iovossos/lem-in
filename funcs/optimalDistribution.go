// optimalDistribution.go
package funcs

import (
	"fmt"
	"sort"
	"strings"
)

// AntAssignment stores which ant is assigned to which path and its order on that path.
type AntAssignment struct {
	antID     int // The global ant number (starting at 1)
	pathIndex int // Index into the paths slice (after sorting)
	order     int // The order of this ant on the path (1 means first ant on that path)
}

// sortPaths sorts the available paths in ascending order by their length.
func sortPaths(paths [][]string) [][]string {
	// Make a copy so that the original order remains unchanged.
	sorted := make([][]string, len(paths))
	copy(sorted, paths)
	sort.Slice(sorted, func(i, j int) bool {
		return len(sorted[i]) < len(sorted[j])
	})
	return sorted
}

// assignAnts distributes the ants among the available paths using a greedy strategy.
// It returns a slice of AntAssignment.
func assignAnts(totalAnts int, paths [][]string) []AntAssignment {
	// Check if we have any valid paths. If not, return an empty assignment.
	if len(paths) == 0 {
		return []AntAssignment{}
	}

	assignments := make([]AntAssignment, 0, totalAnts)
	// assignedCounts[i] holds the number of ants already assigned to paths[i]
	assignedCounts := make([]int, len(paths))

	for ant := 1; ant <= totalAnts; ant++ {
		// Find the path where the candidate finish time is minimal.
		bestIdx := 0
		// Candidate finish time = (assignedCounts + 1) + (len(path) - 1) = assignedCounts + len(path)
		bestFinish := assignedCounts[0] + len(paths[0])
		for i := 1; i < len(paths); i++ {
			candidate := assignedCounts[i] + len(paths[i])
			if candidate < bestFinish {
				bestFinish = candidate
				bestIdx = i
			}
		}
		// Assign this ant to the chosen path.
		assignedCounts[bestIdx]++
		assignments = append(assignments, AntAssignment{
			antID:     ant,
			pathIndex: bestIdx,
			order:     assignedCounts[bestIdx], // The order number on the chosen path.
		})
	}
	return assignments
}

// SimulateAnts runs a turn-by-turn simulation of ant movement along their assigned paths.
// It returns a slice of strings, where each string is one turn's moves in the format "L<antID>-<room>".
func SimulateAnts(totalAnts int, paths [][]string, assignments []AntAssignment) []string {
	// Determine the maximum number of turns required.
	maxTurn := 0
	for _, a := range assignments {
		finish := a.order + len(paths[a.pathIndex]) - 1
		if finish > maxTurn {
			maxTurn = finish
		}
	}
	// Prepare a map for quick lookup of an ant's assignment by its antID.
	antAssignMap := make(map[int]AntAssignment)
	for _, a := range assignments {
		antAssignMap[a.antID] = a
	}
	// We'll simulate from turn 1 to maxTurn.
	resultLines := make([]string, 0)
	// To keep moves in a consistent order, we'll iterate over ants in order of their ID.
	antIDs := make([]int, totalAnts)
	for i := 1; i <= totalAnts; i++ {
		antIDs[i-1] = i
	}
	// Simulate each turn.
	for t := 1; t <= maxTurn; t++ {
		moves := make([]string, 0)
		for _, antID := range antIDs {
			a := antAssignMap[antID]
			// An ant can only start moving from its launch turn (which is its order).
			if t >= a.order {
				// Its position index on the path (starting at 0 for the start room) is:
				// pos = t - a.order + 1, because on its launch turn it moves from start to the first room.
				pos := t - a.order + 1
				// Only print a move if the ant is still en route.
				// When pos equals len(path), the ant has reached the end.
				if pos < len(paths[a.pathIndex]) {
					room := paths[a.pathIndex][pos]
					moves = append(moves, fmt.Sprintf("L%d-%s", antID, room))
				}
			}
		}
		// If any moves occurred during this turn, record them.
		if len(moves) > 0 {
			resultLines = append(resultLines, strings.Join(moves, " "))
		}
	}
	return resultLines
}

// OptimalAntDistribution combines the above steps:
// It sorts the given paths, assigns ants to them using a greedy strategy,
// simulates their movements turn by turn, and returns the list of turn strings.
func OptimalAntDistribution(totalAnts int, paths [][]string) []string {
	sortedPaths := sortPaths(paths)
	// Check if there are any valid paths.
	if len(sortedPaths) == 0 {
		fmt.Println("Valid paths not found")
		return []string{}
	}
	assignments := assignAnts(totalAnts, sortedPaths)
	return SimulateAnts(totalAnts, sortedPaths, assignments)
}
