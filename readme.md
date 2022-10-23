# GoRobot

A robot can walk around in a room where the floor is represented as a number of fields in a wire mesh. Input is first two numbers, which tells the robot how big the room is:

`5 7`

Which means that the room is 5 fields wide and is 7 fields deep.
The size of the room follows two digits and one letter indicating the starting

`3 3 N`

Which means that the robot is in field (3, 3) and faces north. Subsequently, the

* L Turn left
* R Turn right
* F Walk forward

Example:

`LFFRFRFRFF`

After the last command is received, the robot must report which field it is in

Example:

`5 5`



## How to run the programm

Run `go-robot` script or, clone the repo,

`cd go-robot`

`go run main.go`

How to run the tests:

`make test`

# GoRobot API version

[https://github.com/tarcea/go-robot-api](https://github.com/tarcea/go-robot-api)

## Tech stack

* Go