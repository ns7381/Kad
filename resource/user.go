package resource

import "github.com/ns7381/Kad/models"

type User struct {
	Name string `json:"name" description:"name of the user" default:"john"`
}

func CreateUser(user *User) (error) {
	return models.InsertUser(user.Name)
}
