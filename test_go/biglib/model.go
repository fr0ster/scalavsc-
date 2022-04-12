package biglib

import (
	"encoding/json"
	"net/http"
)

type Model interface {
	ID() string
}

type DataProvider[MODEL Model] interface {
	FindByID(id string) (MODEL, error)
	List() ([]MODEL, error)
	Update(id string, model MODEL) error
	Insert(model MODEL) error
	Delete(id string) error
}

type HTTPHandler[MODEL Model] struct {
	dataProvider DataProvider[MODEL]
}

func (h HTTPHandler[MODEL]) FindByID(rw http.ResponseWriter, req *http.Request) {
	// validate request here
	id := "my_id" // extract id here
	model, err := h.dataProvider.FindByID(id)
	if err != nil {
		// error handling here
		return
	}
	err = json.NewEncoder(rw).Encode(model)
	if err != nil {
		// error handling here
		return
	}
}
