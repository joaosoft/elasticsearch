package elastic

import (
	"encoding/json"

	"bytes"
	"fmt"
	"text/template"

	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

type SearchResponse struct {
	Took     int64  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int64 `json:"total"`
		Successful int64 `json:"successful"`
		Skipped    int64 `json:"skipped"`
		Failed     int64 `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int64 `json:"total"`
		MaxScore int64 `json:"max_score"`
		Hits     []struct {
			Index  string          `json:"_index"`
			Type   string          `json:"_type"`
			ID     string          `json:"_id"`
			Score  int64             `json:"_score"`
			Source json.RawMessage `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
	*OnError
	*OnErrorDocumentNotFound
}

type OnErrorDocumentNotFound struct {
	Index string `json:"_index"`
	Type  string `json:"_type"`
	ID    string `json:"_id"`
	Found bool   `json:"found"`
}

type SearchService struct {
	client *Elastic
	index  string
	typ    string
	id     string
	body   []byte
	object interface{}
	method web.Method
}

func NewSearchService(client *Elastic) *SearchService {
	return &SearchService{
		client: client,
		method: web.MethodGet,
	}
}

func (e *SearchService) Index(index string) *SearchService {
	e.index = index
	return e
}

func (e *SearchService) Type(typ string) *SearchService {
	e.typ = typ
	return e
}

func (e *SearchService) Id(id string) *SearchService {
	e.id = id
	return e
}

func (e *SearchService) Query(query []byte) *SearchService {
	e.body = query
	return e
}

func (e *SearchService) Object(object interface{}) *SearchService {
	e.object = object
	return e
}

type SearchTemplate struct {
	Data interface{} `json:"data,omitempty"`
	From int         `json:"from,omitempty"`
	Size int         `json:"size,omitempty"`
}

func (e *SearchService) Template(path, name string, data *SearchTemplate, reload bool) *SearchService {
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

func (e *SearchService) Execute() (*SearchResponse, error) {

	if e.body != nil {
		e.method = web.MethodPost
	}

	var q string
	if e.typ != "" {
		q += fmt.Sprintf("/%s", e.typ)
	}

	if e.id != "" {
		q += fmt.Sprintf("/%s", e.id)
	} else {
		q += "/_search"
	}

	request, err := e.client.Client.NewRequest(e.method, fmt.Sprintf("%s/%s%s", e.client.config.Endpoint, e.index, q))
	if err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	response, err := request.WithBody(e.body, web.ContentTypeApplicationJSON).Send()
	if err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	elasticResponse := SearchResponse{}
	if err := json.Unmarshal(response.Body, &elasticResponse); err != nil {
		e.client.logger.Error(err)
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	if elasticResponse.OnError != nil {
		return &elasticResponse, nil
	}

	rawHits := make([]json.RawMessage, len(elasticResponse.Hits.Hits))
	for i, rawHit := range elasticResponse.Hits.Hits {
		rawHits[i] = rawHit.Source
	}

	arrayHits, err := json.Marshal(rawHits)
	if err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	if err := json.Unmarshal(arrayHits, e.object); err != nil {
		return nil, errors.New(errors.ErrorLevel, 0, err)
	}

	return &elasticResponse, nil
}
