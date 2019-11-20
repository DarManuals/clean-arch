package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DarManuals/clean-arch/models"
	"github.com/DarManuals/clean-arch/repository/mock"
)

func TestName(t *testing.T) {
	testErr := errors.New("test expectedErr")
	userRepo := mock.UserRepository{
		GetFn: func(id int) (user models.User, e error) {
			if id == 123 {
				return models.User{}, testErr
			}
			return models.User{
				ID:   id,
				Name: "test",
			}, nil
		},
	}

	testCases := []struct {
		name           string
		input          int
		expectedResult models.User
		expectedErr    error
	}{
		{
			name:  "success case",
			input: 1,
			expectedResult: models.User{
				ID:   1,
				Name: "test",
			},
		},
		{
			name:        "error case",
			input:       123,
			expectedErr: testErr,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result, err := NewUserService(userRepo).Retrieve(test.input)

			if !reflect.DeepEqual(test.expectedErr, err) {
				t.Fatalf("\nexpected error:\n %+v\nbut got:\n %+v", test.expectedErr, err)
			}

			if !reflect.DeepEqual(test.expectedErr, err) {
				t.Fatalf("\nexpected result:\n %+v\nbut got:\n %+v", test.expectedResult, result)
			}
		})
	}
}
