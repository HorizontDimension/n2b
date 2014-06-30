package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

func AgentTransferCol(s *mgo.Session) *mgo.Collection {
	return s.DB("n2b").C("TransferAgent")
}

type AgentTransfer struct {
	Id       bson.ObjectId `bson:"_id"`
	OldName  string
	NewName  string
	OldNif   string
	NewNif   string
	Hardlock string
	Proof    bson.ObjectId
}

func (a *AgentTransfer) Save(s *mgo.Session) error {
	//its a new record set the Id
	if a.Id == "" {
		a.Id = bson.NewObjectId()
	}

	_, err := AgentTransferCol(s).Upsert(bson.M{"_id": a.Id}, a)
	if err != nil {
		log.Println("Unable to save user account", "user", a, "error", err)
		return err
	}
	return nil
}

func (a *AgentTransfer) Validate() {

}
