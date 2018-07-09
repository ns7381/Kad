package handler

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"github.com/ns7381/Kad/client"
	"log"
	errorsK8s "k8s.io/apimachinery/pkg/api/errors"
	"github.com/ns7381/Kad/resource/deployment"
)

type APIHandler struct {
	cManager client.ClientManager
}

func CreateHttpAPIHandler(cManager client.ClientManager) (http.Handler, error){
	apiHandler := APIHandler{cManager: cManager}
	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)

	apiV1Ws := new(restful.WebService)

	//InstallFilters(apiV1Ws, cManager)

	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	apiV1Ws.Route(
		apiV1Ws.POST("/appdeployment").
			To(apiHandler.handleDeploy).
			Reads(deployment.AppDeploymentSpec{}).
			Writes(deployment.AppDeploymentSpec{}))
	return wsContainer, nil
}

func (apiHandler *APIHandler) handleDeploy(request *restful.Request, response *restful.Response) {
	k8sClient, err := apiHandler.cManager.Client(request)
	if err != nil {
		handleInternalError(response, err)
		return
	}

	appDeploymentSpec := new(deployment.AppDeploymentSpec)
	if err := request.ReadEntity(appDeploymentSpec); err != nil {
		handleInternalError(response, err)
		return
	}
	if err := deployment.DeployApp(appDeploymentSpec, k8sClient); err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusCreated, appDeploymentSpec)
}

// Handler that writes the given error to the response and sets appropriate HTTP status headers.
func handleInternalError(response *restful.Response, err error) {
	log.Print(err)
	statusCode := http.StatusInternalServerError
	statusError, ok := err.(*errorsK8s.StatusError)
	if ok && statusError.Status().Code > 0 {
		statusCode = int(statusError.Status().Code)
	}
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(statusCode, err.Error()+"\n")
}