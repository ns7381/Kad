package resource

import (
	"github.com/ns7381/Kad/models"
	"github.com/ns7381/Kad/cloud"
	"github.com/ns7381/Kad/build"
)

var RegistryURL = "bigdata.registry.com:5000/library/iop/"

func CreateApp(app *models.Application) (error) {
	err := models.InsertApp(app)

	if err != nil {
		return err
	}

	err = build.CreateJob(app)
	err = build.Build(app.Name, 1)

	deployment := new(models.Deployment)
	deployment.Name = app.Name
	deployment.ContainerImage = RegistryURL + app.Name
	deployment.AppId = app.Id
	err = cloud.Deploy(*deployment)
	if err != nil {
		return err
	}
	err = cloud.NewService(deployment.Name)
	return err
}

func GetApplications() (apps []models.Application, err error) {

	return models.GetApps()
}
