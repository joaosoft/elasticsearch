package elastic

import "encoding/json"

type Term struct {
	mappings map[string]interface{}
	field    string
}

func NewTerm(field string, value interface{}) *Term {
	new := &Term{
		field: field,
	}

	new.mappings["value"] = value

	return new
}

func (m *Term) Boost(value float64) *Term {
	m.mappings["boost"] = value
	return m
}

func (m *Term) Bytes() []byte {
	data := map[string]map[string]interface{}{"Term": m.mappings}
	bytes, _ := json.Marshal(data)

	return bytes
}

func (m *Term) String() string {
	return string(m.Bytes())
}
