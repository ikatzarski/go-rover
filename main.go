package main

import "fmt"

func main() {
	println("Mars Rover")
}

type Position struct {
	y int
}

type Rover struct {
	direction string
	position  Position
	gridSize  int
}

func NewRover() *Rover {
	return &Rover{
		direction: "N",
		position: Position{
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
			r.position.y = (r.position.y + 1) % r.gridSize
		}
	}

	return fmt.Sprintf("0:%d:%s", r.position.y, r.direction)
}

func (r *Rover) rotateRight() {
	switch r.direction {
	case "N":
		r.direction = "E"
	case "E":
		r.direction = "S"
	case "S":
		r.direction = "W"
	case "W":
		r.direction = "N"
	}
}

func (r *Rover) rotateLeft() {
	switch r.direction {
	case "N":
		r.direction = "W"
	case "W":
		r.direction = "S"
	case "S":
		r.direction = "E"
	case "E":
		r.direction = "N"
	}
}
