package user

import (
	"github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/model"
	"github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/repository"
)

// User 個人情報の構造体
type User struct {
	userRepo repository.User
}

type UserService interface {
	Create(*model.User) (*model.User, error)
}

// NewUser ユーザ構造体を初期化
func NewUser(userRepo repository.User) UserService {
	return &User{userRepo: userRepo}
}
