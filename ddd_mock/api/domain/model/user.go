package model

// User 個人情報構造体
type User struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Address string `json:"Address"`
}
