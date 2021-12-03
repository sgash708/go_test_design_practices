package repository

import "github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/model"

// User Userモデルへの登録できる構造体
type UserGroup interface {
	Create(usergroup *model.UserGroup) (*model.User, error)
}
