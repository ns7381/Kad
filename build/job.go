package build

import (
	"fmt"
	"github.com/ns7381/Kad/models"
	"text/template"
	"bytes"
	"github.com/pkg/errors"
)

func CreateJob(app *models.Application) error {
	instanceJob, err := JenkinsClient().CreateJobInFolder(parseJobXML(*app), app.Name)
	if err != nil {
		panic(err)
		return err
	}

	if instanceJob != nil {
		fmt.Println("Job has been created!")
	}
	return err
}

func Build(jobName string, number int64) error {
	_, err := JenkinsClient().BuildJob(jobName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	data, err := JenkinsClient().GetBuild(jobName, number)
	for true {
		data, err = JenkinsClient().GetBuild(jobName, number)
		if "SUCCESS" == data.GetResult() {
			fmt.Println("This build succeeded")
			break
		}
		if "FAILURE" == data.GetResult() {
			fmt.Println("This build failure")
			return errors.New("Job build failure.")
		}
	}
	return err
}

func parseJobXML(app models.Application) string {
	t := template.New("job.xml")
	t, _ = t.ParseFiles("E:\\go_path\\src\\github.com\\ns7381\\Kad\\build\\job.xml")
	buf := new(bytes.Buffer)
	e := t.Execute(buf, app)
	if e != nil {
		fmt.Println(e)
	}
	return buf.String()
}