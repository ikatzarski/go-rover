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
}

func NewRover() *Rover {
	return &Rover{
		direction: "N",
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
			r.position.y++
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
