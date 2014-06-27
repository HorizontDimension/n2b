package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func main() {
	wsContainer := restful.NewContainer()
	ta := &TransferAgentResource{}
	ua := &UpgradeAgentResource{}

	ta.Register(wsContainer)
	ua.Register(wsContainer)
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
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
