package usermanager

import "time"

const (
	MaxAuth = 3
	passwordFile = ".password"
	userFile = "users.json"
)

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Birthday time.Time `json:"birthday"`
	Tel string `json:"tel"`
	Addr string `json:"addr"`
	Desc string `json:"desc"`
}