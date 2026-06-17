package main

type Point struct {
	x int
	y int
}

type Game struct {
	snake    []Point
	food     Point
	malware  []Point
	dir      Point
	score    int
	level    int
	gameOver bool
	width    int
	height   int
	quit     chan struct{}
}

func NewGame() *Game {
	return &Game{
		snake:    make([]Point, 0),
		food:     Point{x: 4, y: 4},
		malware:  make([]Point, 0),
		dir:      Point{x: 1, y: 2},
		score:    0,
		level:    1,
		gameOver: false,
		height:   500,
		width:    500,
		quit:     make(chan struct{}),
	}
}

func main() {
}
