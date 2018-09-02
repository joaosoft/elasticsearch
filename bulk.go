package elastic

import (
	"encoding/json"

	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/joaosoft/errors"
)

type BulkResponse struct {
	Took   int  `json:"took"`
	Errors bool `json:"errors"`
	Items  []struct {
		Index struct {
			Index   string `json:"_index"`
			Type    string `json:"_type"`
			ID      string `json:"_id"`
			Version int    `json:"_version"`
			Result  string `json:"result"`
			Shards  struct {
				Total      int `json:"total"`
				Successful int `json:"successful"`
				Failed     int `json:"failed"`
			} `json:"_shards"`
			Status      int `json:"status"`
			SeqNo       int `json:"_seq_no"`
			PrimaryTerm int `json:"_primary_term"`
		} `json:"index,omitempty"`
		Delete struct {
			Index   string `json:"_index"`
			Type    string `json:"_type"`
			ID      string `json:"_id"`
			Version int    `json:"_version"`
			Result  string `json:"result"`
			Shards  struct {
				Total      int `json:"total"`
				Successful int `json:"successful"`
				Failed     int `json:"failed"`
			} `json:"_shards"`
			Status      int `json:"status"`
			SeqNo       int `json:"_seq_no"`
			PrimaryTerm int `json:"_primary_term"`
		} `json:"delete,omitempty"`
		Create struct {
			Index   string `json:"_index"`
			Type    string `json:"_type"`
			ID      string `json:"_id"`
			Version int    `json:"_version"`
			Result  string `json:"result"`
			Shards  struct {
				Total      int `json:"total"`
				Successful int `json:"successful"`
				Failed     int `json:"failed"`
			} `json:"_shards"`
			Status      int `json:"status"`
			SeqNo       int `json:"_seq_no"`
			PrimaryTerm int `json:"_primary_term"`
		} `json:"create,omitempty"`
		Update struct {
			Index   string `json:"_index"`
			Type    string `json:"_type"`
			ID      string `json:"_id"`
			Version int    `json:"_version"`
			Result  string `json:"result"`
			Shards  struct {
				Total      int `json:"total"`
				Successful int `json:"successful"`
				Failed     int `json:"failed"`
			} `json:"_shards"`
			Status      int `json:"status"`
			SeqNo       int `json:"_seq_no"`
			PrimaryTerm int `json:"_primary_term"`
		} `json:"update,omitempty"`
	} `json:"items"`
}
type BulkService struct {
	client *Elastic
	index  string
	typ    string
	id     string
	body   []byte
	method string
	buffer *bytes.Buffer
}

func NewBulkService(e *Elastic) *BulkService {
	return &BulkService{
		buffer: bytes.NewBufferString(""),
		client: e,
		method: http.MethodPost,
	}
}

func (e *BulkService) Index(index string) *BulkService {
	e.index = index
	return e
}

func (e *BulkService) Type(typ string) *BulkService {
	e.typ = typ
	return e
}

func (e *BulkService) Id(id string) *BulkService {
	e.id = id
	return e
}

func (e *BulkService) Body(body interface{}) *BulkService {
	switch v := body.(type) {
	case []byte:
		e.body = v
	default:
		e.body, _ = json.Marshal(v)
	}
	return e
}

func (e *BulkService) Upsert() error {
	e.buffer.WriteString(fmt.Sprintf(BulkUpsertHeader, e.index))
	e.buffer.Write(e.body)
	e.buffer.WriteString("\n")

	return nil
}

func (e *BulkService) Create() error {
	e.buffer.WriteString(fmt.Sprintf(BulkCreateHeader, e.index))
	e.buffer.Write(e.body)
	e.buffer.WriteString("\n")

	return nil
}

func (e *BulkService) Update() error {
	e.buffer.WriteString(fmt.Sprintf(BulkUpdateHeader, e.index))
	e.buffer.Write(e.body)
	e.buffer.WriteString("\n")

	return nil
}

func (e *BulkService) Delete() error {
	e.buffer.WriteString(fmt.Sprintf(BulkDeleteHeader, e.index))
	e.buffer.WriteString("\n")

	return nil
}

func (e *BulkService) Execute() (*BulkResponse, error) {

	// create data on elastic
	reader := bytes.NewReader(e.body)

	request, err := http.NewRequest(e.method, fmt.Sprintf("%s/_bulk", e.client.config.Endpoint), reader)
	if err != nil {
		return nil, errors.New("0", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.New("0", err)
	}
	defer response.Body.Close()

	// unmarshal data
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("0", err)
	}

	elasticResponse := BulkResponse{}
	if err = json.Unmarshal(body, &elasticResponse); err != nil {
		return nil, errors.New("0", err)
	}

	if elasticResponse.Errors {
		return &elasticResponse, errors.New("0", "error executing the request")
	}

	return &elasticResponse, nil
}
