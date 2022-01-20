package faker

import (
	"encoding/json"
	"github.com/wander4747/faker-go/locale"
)

type nameStruct struct {
	FirstName []string `json:"first_name"`
	LastName  []string `json:"last_name"`
}
type name struct {
	*Fake
	File []byte
	data *nameStruct
}

type NameInterface interface {
	FirstName() string
	LastName() string
	FullName() string
}

func (f *Fake) Name() NameInterface {
	loader := locale.Loader(f.Locale, locale.NAME)

	data, err := json.Marshal(loader)
	if err != nil || loader == nil {
		panic("error converter struct")
	}

	name := &name{f, data, nil}
	name.data, err = name.getData()
	return name
}

func (n *name) FirstName() string {
	i := random(len(n.data.FirstName))
	return n.data.FirstName[i]
}

func (n *name) LastName() string {
	i := random(len(n.data.LastName))
	return n.data.LastName[i]
}

func (n *name) FullName() string {
	return n.FirstName() + " " + n.LastName()
}

func (n *name) getData() (*nameStruct, error) {
	var names nameStruct

	err := json.Unmarshal(n.File, &names)
	if err != nil {
		return nil, err
	}

	return &names, nil
}
