package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

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
	_, err := AgentTransferCol(s).Upsert(bson.M{"_id": a.Id}, a)
	if err != nil {
		log.Println("Unable to save user AgentTransferRequest", "AgentTransferRequest", a, "error", err)
		return err
	}
	return nil
}

func (a *AgentUpgradeRequest) Validate() {

}
