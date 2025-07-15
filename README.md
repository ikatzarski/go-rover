# Mars Rover Kata

## Task

Develop an API for a rover that moves around a 10x10 grid plateau.

The rover has:

- X, Y coordinates representing its position on the grid.
- Direction represented by the letters N, E, S, W (North, East, South, West).
- A starting position: 0:0:N.

## Input

The rover takes as an input a string of commands. Each character is a single move command.

The allowed characters are:

- L and R - rotate the direction of the rover.
- M - move the rover forward 1 square on the grid in the direction it is facing.

If the rover reaches the end of the grid, it wraps around.

## Output

The output is the final position of the rover. The final position is a string like "2:3:W" where:

- 2 is the X coordinate.
- 3 is the Y coordinate.
- W is the position of the rover.

## Examples

- Input: MMRMMLM, Output: 2:3:N
- Input: MMMMMMMMMM, Output: 0:0:N (wrap around)

## Example Interface

```go
type MarsRover interface {
	Execute(command string) string
}
```

## Rules

- The rover receives a string of commands like "RMMLM" and returns its final position - "2:1:N".
- The rover wraps around if it reaches the end of the grid.
