package resources

import (
	"github.com/HorizontDimension/n2b/form.n2b.pt/server/models"
	"github.com/emicklei/go-restful"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"bytes"
	"log"
	"net/smtp"
	"text/template"

	"github.com/jordan-wright/email"
)

type TransferAgentResource struct {
	Session *mgo.Session
}

func (t *TransferAgentResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/transfers").
		Consumes("multipart/form-data").
		Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.POST("/new").To(t.TransferAgent).
		// docs
		Doc("Submit a request for agent transfer ").
		Operation("TransferAgent"))

	container.Add(ws)

}

func (t *TransferAgentResource) TransferAgent(request *restful.Request, response *restful.Response) {
	//err, _ := ioutil.ReadAll(request.Request)
	err := request.Request.ParseMultipartForm(5 << 20)
	if err != nil {
		log.Println(err)
	}
	form := request.Request.Form

	atr := &models.AgentTransferRequest{
		Id:       bson.NewObjectId(),
		Hardlock: form.Get("Hardlock"),
		Proof:    bson.NewObjectId(),
		OldAgent: models.Agent{
			Name: form.Get("OldName"),
			Nif:  form.Get("OldNif"),
		},
		NewAgent: models.Agent{
			Name: form.Get("OldName"),
			Nif:  form.Get("OldNif"),
		},
	}

	atr.Validate()

	_, fileheader, err := request.Request.FormFile("file")

	if err != nil {
		log.Println("error GettingFile from request")
	}

	err = models.AddIFile(t.Session, fileheader, atr.Proof)
	if err != nil {
		log.Println(err)
	}

	atr.Save(t.Session)

	// Create a new template and parse the letter into it.
	template := template.Must(template.New("transfer").Parse(TransferAgentTemplate))

	buf := new(bytes.Buffer)
	err = template.Execute(buf, atr)
	if err != nil {
		log.Println(err)
	}
	//send the email
	e := email.NewEmail()
	e.From = "IT Department <it@euroneves.pt>"
	e.To = []string{"it@euroneves.pt"}
	e.Subject = "Solicitação de Transferência de Agente |N2B – International, Lda. |"
	e.Text = buf.Bytes()
	err = e.Send("webmail.euroneves.pt:587", smtp.PlainAuth("", "it@euroneves.pt", "##Hd2013", "webmail.euroneves.pt"))
	if err != nil {
		panic(err)
	}

}

const TransferAgentTemplate = `

Assunto: Solicitação de Transferência de Agente |N2B – International, Lda. |
Anexo: http://localhost:8080/files/{{.Proof.Hex}}
Mensagem: Venho por este meio Solicitar que os Hardlocks {{.Hardlock}} sejam transferidos do Agente {{.OldAgent.Name}} com o NIF: {{.OldAgent.Nif}} para o Agente {{.NewAgent.Name}} com o NIF: {{.NewAgent.Nif}} conforme declaração do
cliente em anexo.
Atentamente
{{.NewAgent.Name}},
{{.NewAgent.Nif}}

Para Futura Referencia utilize este número {{.Id.Hex}}
`
