package lemin

type Room struct {
	name       string
	x          int
	y          int
	connected  []*Room
	hasAnt     bool
	visited    bool
	stepsToEnd int
}

type Ant struct {
	name      string
	location  *Room
	isDead    bool
	pathIndex int
	path      []*Room
}
