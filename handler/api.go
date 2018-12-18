package handler

import (
	"github.com/emicklei/go-restful"
	"net/http"
	"log"
	"github.com/ns7381/Kad/resource"
	"github.com/ns7381/Kad/models"
)

type APIHandler struct {
}

func CreateHttpAPIHandler() (http.Handler, error) {
	apiHandler := APIHandler{}
	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)

	apiV1Ws := new(restful.WebService)

	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	apiV1Ws.Route(
		apiV1Ws.POST("/users").
			To(apiHandler.handleCreateUser).
			Reads(resource.User{}).
			Writes(resource.User{}))

	apiV1Ws.Route(
		apiV1Ws.POST("/apps").
			To(apiHandler.handleCreateApp).
			Reads(models.Application{}).
			Writes(models.Application{}))

	apiV1Ws.Route(
		apiV1Ws.GET("/apps").
			To(apiHandler.handleGetApps).Writes([]models.Application{}))
	return wsContainer, nil
}

func (apiHandler *APIHandler) handleCreateUser(request *restful.Request, response *restful.Response) {
	user := new(resource.User)
	if err := request.ReadEntity(user); err != nil {
		handleInternalError(response, err)
		return
	}
	if err := resource.CreateUser(user); err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusCreated, user)
}

func (apiHandler *APIHandler) handleCreateApp(request *restful.Request, response *restful.Response) {
	app := new(models.Application)
	if err := request.ReadEntity(app); err != nil {
		handleInternalError(response, err)
		return
	}
	if err := resource.CreateApp(app); err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusCreated, app)
}

func (apiHandler *APIHandler) handleGetApps(request *restful.Request, response *restful.Response) {

	result, err := resource.GetApplications()
	if err != nil {
		handleInternalError(response, err)
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

// Handler that writes the given error to the response and sets appropriate HTTP status headers.
func handleInternalError(response *restful.Response, err error) {
	log.Print(err)
	statusCode := http.StatusInternalServerError
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(statusCode, err.Error()+"\n")
}
