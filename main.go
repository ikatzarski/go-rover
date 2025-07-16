package main

import "fmt"

func main() {
	println("Mars Rover")
}

var directions = map[int]string{
	0: "N",
	1: "E",
	2: "S",
	3: "W",
}

type Position struct {
	x int
	y int
}

type Rover struct {
	direction int
	position  Position
	gridSize  int
}

func NewRover() Rover {
	return Rover{
		direction: 0,
		position: Position{
			x: 0,
			y: 0,
		},
		gridSize: 10,
	}
}

func (r *Rover) Execute(command string) string {
	for _, c := range command {
		switch c {
		case 'R':
			r.rotateRight()
		case 'L':
			r.rotateLeft()
		case 'M':
			r.moveForward()
		}
	}

	return fmt.Sprintf("%d:%d:%s", r.position.x, r.position.y, directions[r.direction])
}

func (r *Rover) rotateRight() {
	r.direction = (r.direction + 1) % len(directions)
}

func (r *Rover) rotateLeft() {
	r.direction = (r.direction - 1 + len(directions)) % len(directions)
}

func (r *Rover) moveForward() {
	switch directions[r.direction] {
	case "N":
		r.position.y = (r.position.y + 1) % r.gridSize
	case "E":
		r.position.x = (r.position.x + 1) % r.gridSize
	case "S":
		r.position.y = (r.position.y - 1 + r.gridSize) % r.gridSize
	case "W":
		r.position.x = (r.position.x - 1 + r.gridSize) % r.gridSize
	}
}
