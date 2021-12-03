package service

import "github.com/sgash708/go_test_design_mock/ddd_mock/api/domain/repository"

// User 個人情報の構造体
type User struct {
	userRepo      repository.User
	userGroupRepo repository.UserGroup
}

// func
