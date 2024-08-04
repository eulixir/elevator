package domain

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Direction string

const (
	Up   Direction = "up"
	Down Direction = "down"
)

type Action struct {
	Direction Direction `json:"direction" valid:"required"`
	Floor     int       `json:"floor" valid:"required"`
}

func NewAction() *Action {
	return &Action{}
}

func (action *Action) Validate() error {
	_, err := govalidator.ValidateStruct(action)

	if err != nil {
		return err
	}

	return nil
}
