package user

import (
	"fmt"

	"github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/model"
)

// Create
func (u *User) Create(user *model.User) (*model.User, error) {
	// バリデーション
	if user.Name == "" {
		return nil, fmt.Errorf("ユーザ名が抜けています。")
	}

	createUser, err := u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return createUser, nil
}
