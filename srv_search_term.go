package elastic

import "encoding/json"

type term struct {
	mappings map[string]interface{}
	field    string
}

func NewTerm(field string, value interface{}) *term {
	new := &term{
		field: field,
	}

	new.mappings["value"] = value

	return new
}

func (m *term) Boost(value float64) *term {
	m.mappings["boost"] = value
	return m
}

func (m *term) Bytes() []byte {
	data := map[string]map[string]interface{}{"term": m.mappings}
	bytes, _ := json.Marshal(data)

	return bytes
}

func (m *term) String() string {
	return string(m.Bytes())
}
