package elastic

import (
	"encoding/json"

	"bytes"
	"fmt"
	"text/template"

	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

type CountResponse struct {
	Count  int64 `json:"count"`
	Shards struct {
		Total      int64 `json:"total"`
		Successful int64 `json:"successful"`
		Skipped    int64 `json:"skipped"`
		Failed     int64 `json:"failed"`
	} `json:"_shards"`
	*OnError
	*OnErrorDocumentNotFound
}

type CountService struct {
	client *Elastic
	index  string
	typ    string
	body   []byte
	method web.Method
}

func NewCountService(client *Elastic) *CountService {
	return &CountService{
		client: client,
		method: web.MethodGet,
	}
}

func (e *CountService) Index(index string) *CountService {
	e.index = index
	return e
}

func (e *CountService) Type(typ string) *CountService {
	e.typ = typ
	return e
}

func (e *CountService) Body(body []byte) *CountService {
	e.body = body
	return e
}

type CountTemplate struct {
	Data interface{} `json:"data,omitempty"`
}

func (e *CountService) Template(path, name string, data *CountTemplate, reload bool) *CountService {
	key := fmt.Sprintf("%s/%s", path, name)

	var result bytes.Buffer
	var err error

	if _, found := templates[key]; !found {
		e.client.mux.Lock()
		defer e.client.mux.Unlock()
		templates[key], err = ReadFile(key, nil)
		if err != nil {
			e.client.logger.Error(err)
			return e
		}
	}

	t := template.New(name)
	t, err = t.Parse(string(templates[key]))
	if err == nil {
		if err := t.ExecuteTemplate(&result, name, data); err != nil {
			e.client.logger.Error(err)
			return e
		}

		e.body = result.Bytes()
	} else {
		e.client.logger.Error(err)
		return e
	}

	return e
}

func (e *CountService) Execute() (*CountResponse, error) {

	if e.body != nil {
		e.method = web.MethodPost
	}

	var q string
	if e.typ != "" {
		q += fmt.Sprintf("/%s", e.typ)
	}

	q += "/_count"

	request, err := e.client.Client.NewRequest(e.method, fmt.Sprintf("%s/%s%s", e.client.config.Endpoint, e.index, q))
	if err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	response, err := request.WithBody(e.body, web.ContentTypeApplicationJSON).Send()
	if err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	elasticResponse := CountResponse{}
	if err := json.Unmarshal(response.Body, &elasticResponse); err != nil {
		e.client.logger.Error(err)
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	return &elasticResponse, nil
}
