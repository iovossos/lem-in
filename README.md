# Lem-in

**Zone01 Curriculum Project**

Lem-in is a pathfinding algorithm project that simulates ant colony movements through a network of rooms.

> **Note**: This project is part of the Zone01 programming curriculum. For the official project instructions and requirements, please refer to `Z01_INSTRUCTIONS.MD` in this directory.

## Project Description

Lem-in reads a map description from a file, which includes the number of ants, room definitions, and connections between rooms. The goal is to find the optimal path(s) to move all ants from the start room to the end room in the least number of turns possible.

## Dependencies
### Go
- Go 1.22.2 or Later
### Python (for visualization)
- Python 3.11 or later
- Pygame library

## Instalation 
1. Install Go from https://golang.org/
2. For visualization:
  - Install Python from https://python.org/
  - Install Pygame:
  ```
    pip install pygame
  ```

## Implemented Features

### Core Algorithm
- **File Parsing**: Complete input file parsing for rooms, connections, and ant count
- **Path Finding**: Recursive pathfinding algorithm to discover all possible routes
- **Path Optimization**: Finds optimal non-overlapping path sets for efficient ant movement
- **Ant Simulation**: Full ant movement simulation with turn-by-turn output
- **Multiple Path Support**: Handles multiple simultaneous paths with intelligent ant distribution

### Data Structures
- **Room System**: Room struct with coordinates, connections, and visited state tracking
- **Ant Management**: Ant struct with location, path assignment, and movement state
- **Path Sets**: Complete path set analysis to find optimal combinations
- **Turn Calculation**: Algorithm to determine minimum turns needed for all ants

## How to Run

1. Ensure you have Go installed on your system (version 1.22.2 or later).
2. Clone this repository.
3. Run the program with a maze file as an argument. All maze files should be in the mazes folder.

e.g go run . example05.txt
e.g (with visualization) go run . example05.txt | python visualizer.py

## Visualization Controls
* SPACE: Play/Pause animation
* LEFT/RIGHT Arrows: Navigate between moves when paused
* ESC: Exit visualization


## Project Structure

- `main.go`: Entry point of the application
- `lemin/`: Package containing the core logic
  - `ants.go`: Ant-related functions
  - `parseinput.go`: Input parsing functions
  - `paths.go`: Path finding and optimization
  - `run.go`: Main execution logic
  - `structs.go`: Data structures used in the project
  - `helper.go`: Helper functions called from both in paths.go and ants.go
- `mazes/`: Directory containing example maze files
- `visualizer.py`: Visualization 

## Algorithm Overview

1. Read and parse the input file
2. Find all possible non-overlapping path sets
3. Calculate the optimal set of paths for the given number of ants
4. Assign ants to paths
5. Simulate ant movement through the colony

## Input File Format

#### The input file should follow this format: 
- The first line is a number indicating the number of ants
- The next lines are the rooms, in this format:
`Name_of_room x_coordinate y_coordinate`      
- There must be at least one start room and one end room, indicated by this comment in the previous line:
`## start` or `##end`
- Immediately after the rooms, the tunnels are listed. These are the names of the rooms connected by the tunnel with a dash between them, like this:
`room1-room2` 

Example:
```
3
##start
0 1 0
##end
1 5 0
2 9 0
3 13 0
0-2
2-3
3-1
```

## Authors

Ioannis Vossos, Uipko Stikker, Katerina Vamvasaki, Arsen Tsntsgoukian