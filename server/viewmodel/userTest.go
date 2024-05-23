package viewmodel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidate(t *testing.T) {
	genericErrorFormat := "invalid-parameter-%s"

	for _, tc := range []struct {
		name           string
		userModel      UserModel
		expectedErrors []error
	}{
		{
			name: "ShouldSucessfullyValidateAll",
			userModel: UserModel{
				Name: "Andressa",
				Age:  25,
			},
			expectedErrors: []error{},
		},
		{
			name: "ShouldReturnErrorWhenFieldAgeIsNegative",
			userModel: UserModel{
				Name: "Andressa",
				Age:  -1,
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "age"),
			},
		},
		{
			name: "ShouldReturnErrorWhenFielsNameIsEmpty",
			userModel: UserModel{
				Name: "",
				Age:  25,
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "name"),
			},
		},
		{
			name: "ShouldReturnAllErrors",
			userModel: UserModel{
				Name: "",
				Age:  -1,
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "age"),
				fmt.Errorf(genericErrorFormat, "name"),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErrors, tc.userModel.Validate())
		})
	}
}
