package resources

import (
	"bytes"
	"github.com/HorizontDimension/n2b/form.n2b.pt/server/afr"
	"github.com/HorizontDimension/n2b/form.n2b.pt/server/models"
	"github.com/emicklei/go-restful"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/smtp"
	"reflect"

	"text/template"

	"github.com/jordan-wright/email"
)

type UpgradeAgentResource struct {
	Session *mgo.Session
}

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
	err := request.Request.ParseMultipartForm(5 << 20)
	if err != nil {
		log.Println(err)
	}
	form := request.Request.Form

	aur := &models.AgentUpgradeRequest{
		Id:          bson.NewObjectId(),
		Software:    form.Get("Software"),
		Proof:       bson.NewObjectId(),
		OrderNumber: form.Get("OrderNumber"),
		Agent: models.Agent{
			Name: form.Get("Name"),
			Nif:  form.Get("Nif"),
		},
	}

	errors := aur.Validate()
	if !reflect.DeepEqual(errors, afr.New()) {
		log.Println("we got errors")
		response.WriteAsJson(errors)
		return
	}

	if form.Get("response") == "" {
		errors.Set("InvalidCaptcha", "invalid Captcha")
		response.WriteAsJson(errors)
		return
	}

	challenge := form.Get("challenge")
	resp := form.Get("response")
	if challenge != "" && resp != "" {
		if ok, err := cc.Verify(request.Request.RemoteAddr, challenge, resp); ok {
			log.Println("valid", challenge)

		} else {
			errors.Set("InvalidCaptcha", err.Error())
			response.WriteAsJson(errors)
			return
		}
	}

	_, fileheader, err := request.Request.FormFile("file")

	if err != nil {
		log.Println("error GettingFile from request")
	}

	err = models.AddFile(t.Session, fileheader, aur.Proof)
	if err != nil {
		log.Println(err)
	}

	aur.Save(t.Session)

	// Create a new template and parse the letter into it.
	template := template.Must(template.New("upgrade").Parse(upgradeAgentTemplate))

	buf := new(bytes.Buffer)
	err = template.Execute(buf, aur)
	if err != nil {
		log.Println(err)
	}
	//send the email
	e := email.NewEmail()
	e.From = "IT Department <it@euroneves.pt>"
	e.To = []string{"it@euroneves.pt"}
	e.Subject = "Comprovativo de Upgrade de Concorrência | N2B – International, Lda. |"
	e.Text = buf.Bytes()
	err = e.Send("webmail.euroneves.pt:587", smtp.PlainAuth("", "it@euroneves.pt", "##Hd2013", "webmail.euroneves.pt"))
	if err != nil {
		panic(err)
	}

}

const upgradeAgentTemplate = `
Anexo:
Mensagem: Venho por este meio enviar o comprovativo de upgrade de software {{.Software}} referente à encomenda {{.OrderNumber}}.
Atentamente
{{.Agent.Name}}
{{.Agent.Nif}}
`
