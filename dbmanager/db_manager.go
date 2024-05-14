package dbmanager

import (
	"turns-app-go/model"
)

type usersManagerI interface {
	GetUsers() []model.User
}

type DBManager struct {
	UsersManager usersManagerI
}
