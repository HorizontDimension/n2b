package main

import (
	"github.com/HorizontDimension/n2b/form.n2b.pt/server/resources"
	"github.com/emicklei/go-restful"
	"labix.org/v2/mgo"
	"log"
	"net/http"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	wsContainer := restful.NewContainer()
	ta := &resources.TransferAgentResource{session}
	ua := &resources.UpgradeAgentResource{}
	files := &resources.File{session}

	ta.Register(wsContainer)
	ua.Register(wsContainer)
	files.Register(wsContainer)

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-N2B"},
		AllowedHeaders: []string{"Content-Type"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs and enter http://localhost:8080/apidocs.json in the api input field.
	//config := swagger.Config{
	//	WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
	//	WebServicesUrl: "http://localhost:8080",
	//	ApiPath:        "/apidocs.json",
	//
	//	// Optionally, specifiy where the UI is located
	//	SwaggerPath:     "/apidocs/",
	//	SwaggerFilePath: "/Users/emicklei/xProjects/swagger-ui/dist"}
	//swagger.RegisterSwaggerService(config, wsContainer)

	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())

}
