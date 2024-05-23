package viewmodel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAddresValidate(t *testing.T) {
	genericErrorFormat := "invalid-parameter-%s"

	for _, tc := range []struct {
		name            string
		userAddresModel UserAddresModel
		expectedErrors  []error
	}{
		{
			name: "ShouldSucessfullyValidateAll",
			userAddresModel: UserAddresModel{
				Id:     "123",
				Street: "street_mock",
				City:   "city_mock",
				State:  "state_mock",
			},
			expectedErrors: []error{},
		},
		{
			name: "ShouldReturnErrorWhenFielsStreetIsEmpty",
			userAddresModel: UserAddresModel{
				Id:     "123",
				Street: "",
				City:   "city_mock",
				State:  "state_mock",
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "Street"),
			},
		},
		{
			name: "ShouldReturnErrorWhenFielsCityIsEmpty",
			userAddresModel: UserAddresModel{
				Id:     "123",
				Street: "street_mock",
				City:   "",
				State:  "state_mock",
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "City"),
			},
		},
		{
			name: "ShouldReturnErrorWhenFielsStateIsEmpty",
			userAddresModel: UserAddresModel{
				Id:     "123",
				Street: "street_mock",
				City:   "city_mock",
				State:  "",
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "State"),
			},
		},
		{
			name: "ShouldReturnAllErrors",
			userAddresModel: UserAddresModel{
				Id:     "123",
				Street: "",
				City:   "",
				State:  "",
			},
			expectedErrors: []error{
				fmt.Errorf(genericErrorFormat, "Street"),
				fmt.Errorf(genericErrorFormat, "City"),
				fmt.Errorf(genericErrorFormat, "State"),
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedErrors, tc.userAddresModel.Validate())
		})
	}
}
