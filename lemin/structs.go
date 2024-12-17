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
	name       string
	location   *Room
	active     bool
	isDead     bool
	movesCount int
	path       []*Room
}
