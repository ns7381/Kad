package build

import (
	"sync"
	"github.com/bndr/gojenkins"
)

var once sync.Once
var jenkins *gojenkins.Jenkins

func JenkinsClient() *gojenkins.Jenkins {
	once.Do(func() {
		jenkins = gojenkins.CreateJenkins(nil, "http://10.221.129.29:18080", "admin", "123456a?")
		_, err := jenkins.Init()
		if err != nil {
			panic("Jenkins Init Went Wrong!")
		}
	})
	return jenkins
}
