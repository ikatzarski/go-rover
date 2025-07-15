package main

import "fmt"

func main() {
	println("Mars Rover")
}

type Rover struct {
	direction string
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
		}
	}

	return fmt.Sprintf("0:0:%s", r.direction)
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
