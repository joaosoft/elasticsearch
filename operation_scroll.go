package elastic

import (
	"encoding/json"

	"github.com/joaosoft/web"

	"fmt"

	"github.com/joaosoft/errors"
)

type ScrollResponse struct {
	Index   string `json:"_index,omitempty"`
	Type    string `json:"_type,omitempty"`
	ID      string `json:"_id,omitempty"`
	Version int64  `json:"_version,omitempty"`
	Found   bool   `json:"found"`
	Result  string `json:"result"`
	Created bool   `json:"created"`
}

type ScrollService struct {
	client *Elastic
	index  string
	typ    string
	body   []byte
	method web.Method
}

func NewScrollService(e *Elastic) *ScrollService {
	return &ScrollService{
		client: e,
		method: web.MethodPut,
	}
}

func (e *ScrollService) Index(index string) *ScrollService {
	e.index = index
	return e
}

func (e *ScrollService) Type(typ string) *ScrollService {
	e.typ = typ
	return e
}

func (e *ScrollService) Body(body interface{}) *ScrollService {
	switch v := body.(type) {
	case []byte:
		e.body = v
	default:
		e.body, _ = json.Marshal(v)
	}
	return e
}

func (e *ScrollService) Execute() (*ScrollResponse, error) {

	request, err := e.client.Client.NewRequest(e.method, fmt.Sprintf("%s/%s/%s", e.client.config.Endpoint, e.index, e.typ))
	if err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	response, err := request.WithBody(e.body, web.ContentTypeApplicationJSON).Send()
	if err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	elasticResponse := ScrollResponse{}
	json.Unmarshal(response.Body, &elasticResponse)

	return &elasticResponse, nil
}
