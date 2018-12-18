package models

import (
	"time"
	"github.com/ns7381/Kad/database"
)

type Job struct {
	Id            int64     `xorm:"not null pk autoincr INT(11)"`
	Name          string    `json:"name"`
	AppId         int64     `json:"app_id"`
	CreatedAt     time.Time `xorm:"created"`
}

func InsertJob(job Job) (error) {
	_, err := database.ORM.Insert(&job)
	return err
}