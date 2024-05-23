package viewmodel

import (
	"crud-golang/domain"
	"fmt"
)

type UserAddresModel struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

func (ua UserAddresModel) Validate() []error {
	genericErrorFormat := "invalid-parameter-%s"
	var Errors []error

	if ua.Street == "" {
		Errors = append(Errors, fmt.Errorf(genericErrorFormat, "Street"))
	}

	if ua.City == "" {
		Errors = append(Errors, fmt.Errorf(genericErrorFormat, "City"))
	}

	if ua.State == "" {
		Errors = append(Errors, fmt.Errorf(genericErrorFormat, "State"))
	}

	if len(Errors) > 0 {
		return Errors
	}

	return nil
}

func (ua UserAddresModel) Parse() domain.UserAddres {
	return domain.UserAddres{
		Street: ua.Street,
		City:   ua.City,
		State:  ua.State,
	}
}
