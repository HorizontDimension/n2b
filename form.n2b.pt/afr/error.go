//package angular form error
//author : thesyncim

package main

import (
	"encoding/json"
	"log"
)

type Error struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Errors []Error

func (e *Errors) Set(key, value string) {
	var found bool
	//check if already exists
	for i := 0; i < len(*e); i++ {
		log.Println("enter range")
		//found it
		if (*e)[i].Key == key {
			(*e)[i].Value = value
			found = true
			break
		}
	}
	log.Println("not found")
	if !found {

		*e = append(*e, Error{Key: key, Value: value})
		log.Println(e)

	}

}
func New() Errors {

	return Errors([]Error{})
}

func main() {
	e := New()
	e.Set("asd", "asd")
	e.Set("adsd", "asd")
	e.Set("adsdf", "asd")
	e.Set("adsdsd", "asd")
	e.Set("adssd", "asd")

	e.Set("asd", "asd")
	r, err := json.Marshal(e)

	log.Println(string(r), err)
}
