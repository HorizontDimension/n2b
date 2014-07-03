package models

import (
	"github.com/HorizontDimension/n2b/form.n2b.pt/server/afr"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
)

func AgentTransferCol(s *mgo.Session) *mgo.Collection {
	return s.DB("n2b").C("TransferAgent")
}

type Agent struct {
	Name string
	Nif  string
}

type AgentTransferRequest struct {
	Id       bson.ObjectId `bson:"_id"`
	OldAgent Agent
	NewAgent Agent
	Hardlock string
	Proof    bson.ObjectId
	Created  time.Time
}

func (a *AgentTransferRequest) Save(s *mgo.Session) error {
	//its a new record set the Id
	if a.Id == "" {
		a.Id = bson.NewObjectId()
	}

	a.Created = time.Now()
	_, err := AgentTransferCol(s).Upsert(bson.M{"_id": a.Id}, a)
	if err != nil {
		log.Println("Unable to save user AgentTransferRequest", "AgentTransferRequest", a, "error", err)
		return err
	}
	return nil
}

func (a *AgentTransferRequest) Validate() afr.Errors {
	errors := afr.New()
	log.Println(a)
	if a.OldAgent.Name == "" {
		errors.Set("EmptyOldAgentName", "Agent name cant be empty")
	}
	if a.OldAgent.Nif == "" {
		errors.Set("EmptyOldAgentNif", "Agent nif cant be empty")
	}
	if a.NewAgent.Name == "" {
		errors.Set("EmptyNewAgentName", "Agent name cant be empty")
	}
	if a.NewAgent.Nif == "" {
		errors.Set("EmptyNewAgentNif", "Agent nif cant be empty")
	}
	if a.Hardlock == "" {
		errors.Set("EmptyHardlock", "Hardlock cant be empty")
	}
	return errors
}
