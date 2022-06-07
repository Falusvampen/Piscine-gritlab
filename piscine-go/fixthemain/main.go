package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(Door *Door) bool {
	PrintStr("Door Opening...\n")
	Door.state = true
	return false
}

func CloseDoor(Door *Door) bool {
	PrintStr("Door Closing...\n")
	Door.state = false
	return true
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?\n")
	Door.state = true
	return false
}

func IsDoorClose(Door *Door) bool {
	PrintStr("is the Door closed ?\n")
	Door.state = true
	return false
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == true {
		CloseDoor(door)
	}
}
