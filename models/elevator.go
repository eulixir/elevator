package models

import (
	"sort"

	"github.com/asaskevich/govalidator"
)

type State string

const (
	Stopped State = "stopped"
	Moving  State = "moving"
)

type Elevator struct {
	CurrentDirection Direction `json:"current_direction" valid:"required"`
	CurrentFloor     int       `json:"current_floor" valid:"required"`
	ActionQueue      []Action  `json:"action_queue" valid:"required"`
	State            State     `json:"state" valid:"required"`
}

func NewElevator() *Elevator {
	return &Elevator{ActionQueue: []Action{}}
}

func (elevator *Elevator) Validate() error {
	_, err := govalidator.ValidateStruct(elevator)

	if err != nil {
		return err
	}

	return nil
}

func Reorganize(elevator *Elevator) (result [2][]int) {
	result[0] = []int{}
	result[1] = []int{}

	for _, action := range elevator.ActionQueue {
		if action.Direction == elevator.CurrentDirection {
			result[0] = append(result[0], action.Floor)
			sort.Ints(result[0])
		} else {
			result[1] = append(result[1], action.Floor)
			sortDescending(result[0])
		}
	}

	return
}

func sortDescending(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}
