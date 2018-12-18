package models

import (
	"time"
	db "github.com/ns7381/Kad/database"
	"log"
)

type Application struct {
	Id            int64     `xorm:"not null pk autoincr INT(11)"`
	Name          string    `json:"name"`
	SourceControl string    `json:"source_control"`
	Repository    string    `json:"repository"`
	Branch        string    `json:"branch"`
	Build         string    `json:"build"`
	PkgInclude    string    `json:"pkg_include"`
	ImageType     string    `json:"image_type"`
	CreatedAt     time.Time `xorm:"created"`
}

func InitDB() {
	err := db.ORM.Sync2(new(Application))
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func InsertApp(app *Application) (error) {
	_, err := db.ORM.Insert(app)
	return err
}

func GetApps() (as []Application, err error) {
	err = db.ORM.Desc("created_at").Find(&as)
	return as, err
}
