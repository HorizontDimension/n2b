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

	at := new(models.AgentTransfer)
	at.Hardlock = form.Get("Hardlock")
	at.NewName = form.Get("NewName")
	at.OldName = form.Get("OldName")
	at.NewNif = form.Get("NewNif")
	at.OldNif = form.Get("OldNif")

	at.Validate()

	_, fileheader, err := request.Request.FormFile("file")

	if err != nil {
		log.Println("error GettingFile from request")
	}

	fileId := bson.NewObjectId()

	at.Proof = fileId

	err = models.AddImage(t.Session, fileheader, fileId)
	if err != nil {
		log.Println(err)
	}

	recordId := bson.NewObjectId()
	at.Id = recordId
	at.Save(t.Session)

	// Create a new template and parse the letter into it.
	template := template.Must(template.New("transfer").Parse(TransferAgentTemplate))

	buf := new(bytes.Buffer)
	err = template.Execute(buf, at)
	if err != nil {
		log.Println(err)
	}
	//send the email
	e := email.NewEmail()
	e.From = "IT Department <it@euroneves.pt>"
	e.To = []string{"it@euroneves.pt"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Solicitação de Transferência de Agente |N2B – International, Lda. |"
	//e.Text = []byte("Text Body is, of course, supported!")
	e.Text = buf.Bytes()
	err = e.Send("webmail.euroneves.pt:587", smtp.PlainAuth("", "it@euroneves.pt", "##Hd2013", "webmail.euroneves.pt"))
	if err != nil {
		panic(err)
	}

}

const TransferAgentTemplate = `

Assunto: Solicitação de Transferência de Agente |N2B – International, Lda. |
Anexo: http://192.168.10.110:8080/files/{{.Proof.Hex}}
Mensagem: Venho por este meio Solicitar que os Hardlocks {{.Hardlock}} sejam transferidos do Agente {{.OldName}} com o NIF: {{.OldNif}} para o Agente {{.NewName}} com o NIF: {{.NewNif}} conforme declaração do
cliente em anexo.
Atentamente
{{.NewName}},
{{.NewNif}}

Para Futura Referencia utilize este número {{.Id.Hex}}
`
