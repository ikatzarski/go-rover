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
	return "0:0:N"
}
