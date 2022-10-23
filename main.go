package main

import (
	"fmt"

	"github.com/tarcea/go-robot/game"
)

func main() {
	var roomX, roomY, initX, initY int
	var initOrientation, command string
	fmt.Println("Please enter the WIDTH and DEEP of the room (5 5): ")
	fmt.Scanln(&roomX, &roomY)
	fmt.Println("Please enter the X Y position and ORIENTATION of the robot (1 2 N): ")
	fmt.Scanln(&initX, &initY, &initOrientation)
	fmt.Println("Please enter COMMAND for the robot (RFRFFRFRF): ")
	fmt.Scanln(&command)
	g := game.NewGame(roomX, roomY, initX, initY)
	g.Orientation = initOrientation
	g.Command = command
	game.RunCommand(g)
	fmt.Println("Report:", g.PositionX, g.PositionY, g.Orientation)
}
