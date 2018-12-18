package models

import "time"

type Deployment struct {
	Id            int64     `xorm:"not null pk autoincr INT(11)"`
	Name          string    `json:"name"`
	ContainerImage string `json:"containerImage"`
	AppId         int64     `json:"app_id"`
	CreatedAt     time.Time `xorm:"created"`
}
