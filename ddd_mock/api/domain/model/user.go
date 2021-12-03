package model

// User 個人情報構造体
type User struct {
	ID      int
	Name    string
	Address string
}

// UserGroup ユーザグループ構造体
type UserGroup struct {
	ID      int
	Name    string
	Private bool
}
