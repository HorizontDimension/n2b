//package angular form error
//author : thesyncim
package afr

import (
//"encoding/json"

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

		//found it"github.com/HorizontDimension/n2b/form.n2b.pt/server/
		if (*e)[i].Key == key {
			(*e)[i].Value = value
			found = true
			break
		}
	}

	if !found {
		*e = append(*e, Error{Key: key, Value: value})

	}
}
func New() Errors {
	return Errors([]Error{})
}
