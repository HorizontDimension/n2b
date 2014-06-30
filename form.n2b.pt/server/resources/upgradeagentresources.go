package resources

import (
	"github.com/emicklei/go-restful"
)

type AgentUpgrade struct {
	Name     string
	Nif      string
	Software string
	Proof    []byte
}

type UpgradeAgentResource struct{}

func (t *UpgradeAgentResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/upgrades").
		Consumes("multipart/form-data").
		Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.POST("/new").To(t.UpgradeAgent).
		// docs
		Doc("Submit a request for agent transfer ").
		Operation("TransferAgent"))

	container.Add(ws)

}

func (t *UpgradeAgentResource) UpgradeAgent(request *restful.Request, response *restful.Response) {

	//blah

}
