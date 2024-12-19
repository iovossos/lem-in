# Lem-in

Lem-in is a pathfinding algorithm project that simulates ant colony movements through a network of rooms.

## Project Description

Lem-in reads a map description from a file, which includes the number of ants, room definitions, and connections between rooms. The goal is to find the optimal path(s) to move all ants from the start room to the end room in the least number of turns possible.

## Features

- Parses input files describing the ant colony structure
- Calculates optimal paths for ant movement
- Simulates ant movement through the colony
- Handles multiple paths and distributes ants efficiently

## How to Run

1. Ensure you have Go installed on your system (version 1.22.2 or later).
2. Clone this repository: https://platform.zone01.gr/git/ivossos/lem-in

3. Run the program with a maze file as an argument:


## Project Structure

- `main.go`: Entry point of the application
- `lemin/`: Package containing the core logic
  - `ants.go`: Ant-related functions
  - `parseinput.go`: Input parsing functions
  - `paths.go`: Path finding and optimization
  - `run.go`: Main execution logic
  - `structs.go`: Data structures used in the project
- `mazes/`: Directory containing example maze files

## Algorithm Overview

1. Read and parse the input file
2. Calculate distances from each room to the end
3. Sort connected rooms by distance to end
4. Find all possible path sets
5. Calculate the optimal set of paths
6. Assign ants to paths
7. Simulate ant movement through the colony

## Input File Format

The input file should follow this format: 
Example:


## Authors

John Vossos, Uipko Stikker, Katerina Vamvasaki, Arsen Tsntsgoukian