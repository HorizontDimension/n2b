package main

import (
	"github.com/emicklei/go-restful"
	//"github.com/emicklei/go-restful/swagger"
	"github.com/gorilla/schema"
	"net/http"
	//"github.com/jordan-wright/email"
)

var decoder *schema.Decoder

type TransferAgentResource struct{}

func (t *TransferAgentResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/transfers").
		Consumes("application/x-www-form-urlencoded").
		Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.GET("/new").To(t.TransferAgent).
		// docs
		Doc("Submit a request for agent transfer ").
		Operation("TransferAgent"))

	container.Add(ws)

}

func (t *TransferAgentResource) TransferAgent(request *restful.Request, response *restful.Response) {

	at := new(AgentTransfer)
	err := decoder.Decode(at, request.Request.PostForm)
	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

}

type UpgradeAgentResource struct{}

func (t *UpgradeAgentResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/upgrades").
		Consumes("application/x-www-form-urlencoded").
		Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.GET("/new").To(t.UpgradeAgent).
		// docs
		Doc("Submit a request for agent transfer ").
		Operation("TransferAgent"))

	container.Add(ws)

}

func (t *UpgradeAgentResource) UpgradeAgent(request *restful.Request, response *restful.Response) {

	//blah

}
