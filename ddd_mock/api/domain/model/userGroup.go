package model

// UserGroup ユーザグループ構造体
type UserGroup struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Private bool   `json:"Private"`
}
