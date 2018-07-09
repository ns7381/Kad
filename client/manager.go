package client

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"runtime"
	"path/filepath"
	"github.com/emicklei/go-restful"
	"k8s.io/client-go/rest"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func NewClientManager() ClientManager {
	result := &clientManager{
		kubeConfigPath: filepath.Join(basePath, "kube-config.yaml"),
	}

	return result
}

type ClientManager interface {
	Client(req *restful.Request) (kubernetes.Interface, error)
}

type clientManager struct {
	kubeConfigPath string
}

func (self *clientManager) Client(req *restful.Request) (kubernetes.Interface, error) {
	cfg, err := self.Config(req)
	if err != nil {
		return nil, err
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return client, nil
}
func (self *clientManager) Config(req *restful.Request) (*rest.Config, error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", filepath.Join(basePath, "kube-config.yaml"))
	if err != nil {
		panic(err.Error())
	}
	return cfg, nil
}

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(basePath, "kube-config.yaml"))
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(10 * time.Second)
	}
}
