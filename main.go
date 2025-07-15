package main

func main() {
	println("Mars Rover")
}

type Rover struct {
}

func NewRover() *Rover {
	return &Rover{}
}

func (r *Rover) Execute(command string) string {
	if command == "R" {
		return "0:0:E"
	}

	return "0:0:N"
}
