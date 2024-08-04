package domain

import (
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
