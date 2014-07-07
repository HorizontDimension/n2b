package models

import (
	"github.com/HorizontDimension/n2b/form.n2b.pt/server/afr"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"

	"time"
)

func AgentUpgradeCol(s *mgo.Session) *mgo.Collection {
	return s.DB("n2b").C("UpgradeAgent")
}

type AgentUpgradeRequest struct {
	Id          bson.ObjectId `bson:"_id"`
	Agent       Agent
	Software    string
	Proof       bson.ObjectId
	OrderNumber string
	Created     time.Time
}

func (a *AgentUpgradeRequest) Save(s *mgo.Session) error {
	//its a new record set the Id
	if a.Id == "" {
		a.Id = bson.NewObjectId()
	}

	a.Created = time.Now()
	_, err := AgentUpgradeCol(s).Upsert(bson.M{"_id": a.Id}, a)
	if err != nil {
		log.Println("Unable to save  UpgradeAgent", "UpgradeAgent", a, "error", err)
		return err
	}
	return nil
}

func (a *AgentUpgradeRequest) Validate() afr.Errors {

	errors := afr.New()
	log.Println(a)
	if a.Agent.Name == "" {
		errors.Set("EmptyAgentName", "Agent name cant be empty")
	}
	if a.Agent.Nif == "" {
		errors.Set("EmptyAgentNif", "Agent nif cant be empty")
	}
	if a.OrderNumber == "" {
		errors.Set("EmptyOrderNumber", "Software cant be empty")
	}

	if a.Software == "" {
		errors.Set("EmptyHardlock", "Software cant be empty")
	}
	return errors

}
