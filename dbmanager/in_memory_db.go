package dbmanager

import (
	"encoding/json"
	"fmt"
	"os"
	"turns-app-go/model"
)

type InMemoryUsersDBManager struct {
	Users []model.User
}

func (db *InMemoryUsersDBManager) GetUsers() []model.User {
	return db.Users
}

func NewInMemoryUsersDBManager(filePath string) *InMemoryUsersDBManager {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer file.Close()

	// Decode the JSON data into a slice of Person objects
	var people []model.User
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&people); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print the loaded data
	for _, p := range people {
		fmt.Printf("Name: %s, id: %d phone: %s\n", p.Username, p.ID, p.Phone)
	}

	return &InMemoryUsersDBManager{Users: people}
}
