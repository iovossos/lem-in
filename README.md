# lem-in

A Go implementation of the classic Lem-in problem: finding optimal vertex-disjoint paths through a network of rooms and simulating the movement of ants.

## Team

- **Kostas Apostolou**
- **Yana Kopilova**
- **Vicky Apostolou**


## Features

- **Input Parsing & Validation**: Reads an input file defining the number of ants, rooms, and tunnels. Checks for common errors such as:
  - Invalid or missing ant count
  - Duplicate room names or rooms starting with `L`/`#`
  - Self-linking, duplicate, or reversed tunnels
  - Tunnels referencing undefined rooms
  - Multiple `##start`/`##end` declarations
- **Graph Construction**: Builds an undirected adjacency list of rooms and tunnels.
- **Connectivity Check**: Uses BFS to ensure at least one path exists between the start and end rooms.
- **Vertex-Disjoint Path Extraction**: Implements vertex splitting + Edmonds–Karp max-flow to compute all vertex-disjoint paths from start to end.
- **Ant Distribution & Simulation**: Greedy assignment of ants to minimize total turns, followed by turn-by-turn simulation printing each move.

## Project Structure

```
lem-in/
├── funcs/
│   ├── input.go               # Input parsing and validation
│   ├── buildConnections.go    # Builds adjacency list
│   ├── startEndConnection.go  # Connectivity BFS check
│   ├── maxFlow.go             # Vertex-splitting + Edmonds–Karp algorithm
│   ├── optimalDistribution.go # Path sorting, ant assignment, and simulation
├── main.go                    # Orchestrates parsing, computation, and output
├── README.md                  # Project overview and instructions
```

## Algorithms

- **Parsing**: Validates every line and exits early on errors.
- **Graph Building**: Converts tunnels into an adjacency list for BFS and flow.
- **BFS Check**: Ensures that start and end are connected.
- **Max-Flow (Edmonds–Karp)**: Uses vertex splitting to enforce vertex capacities and finds the maximum number of vertex-disjoint paths.
- **Ant Assignment**: Sorts paths by length, then assigns ants greedily to minimize the completion time.
- **Simulation**: Prints each turn’s moves in the format `L<antID>-<room>`.

## Usage

```bash
go run . <input_file>
```