package repository

import "github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/model"

// User Userモデルへの登録できる構造体
type User interface {
	/* Infra層で使用する関数を記述する */
	Create(u *model.User) (*model.User, error)
}
