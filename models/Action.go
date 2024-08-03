package models

type Direction int

const (
	Up Direction = iota
	Down
)

type Action struct {
	Direction Direction
	Floor     int
}

func Reorganize(actions *[]Action) []Action {

	return *actions
}
