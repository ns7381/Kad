package models

import (
	"log"

	db "github.com/ns7381/Kad/database"
	"time"
)

type User struct {
	Id        int64     `xorm:"not null pk autoincr INT(11)"`
	Name      string
	CreatedAt time.Time `xorm:"created"`
}

func GetUserByID(Id int64) *User {
	var user User
	has, err := db.ORM.Id(Id).Get(&user)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	if has == false {
		return nil
	}
	return &user
}

func GetUserByName(name string) *User {
	var user User
	has, err := db.ORM.Where("name=?", name).Get(&user)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	if has == false {
		return nil
	}
	return &user
}

func GetUsers() *[]User {
	var users []User
	err := db.ORM.Find(&users)
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	return &users
}

func InsertUser(name string) (error) {
	newUser := new(User)
	newUser.Name = name
	newUser.CreatedAt = time.Now()
	_, err := db.ORM.Insert(newUser)
	return err
}

func RemoveUserByID(ID int64) (error) {
	var user User
	_, err := db.ORM.Id(ID).Delete(&user)
	return err
}
