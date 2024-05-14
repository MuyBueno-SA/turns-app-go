package dbmanager

import (
	"turns-app-go/model"
)

type InMemoryUsersDBManager struct {
	Users []model.User
}

func (db *InMemoryUsersDBManager) GetUsers() []model.User {
	return db.Users
}

var UsersSlice = []model.User{
	{
		ID:       0,
		Username: "Virginia D'Espósito",
		Email:    "vir@test.com",
		Phone:    "123456789",
		Activity: "Psicopedagoga",
	},
	{
		ID:       2,
		Username: "Federico Bogado",
		Email:    "fico@test.com",
		Phone:    "123456789",
		Activity: "Developer",
	},
	{
		ID:       4,
		Username: "Susana Horia",
		Email:    "susana@other.com",
		Phone:    "987654321",
		Activity: "Reiki",
	},
}

func DefaultInMemoryUsersDBManager() *InMemoryUsersDBManager {
	return &InMemoryUsersDBManager{Users: UsersSlice}
}
