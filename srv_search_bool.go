package elastic

import "encoding/json"

type bool struct {
	mappings  map[string]interface{}
	must      []Query
	filter    []Query
	should    []Query
	shouldNot []Query
}

func NewBool() *bool {
	new := &bool{}

	return new
}

func (m *bool) Must(value Query) *bool {
	m.mappings["must"] = value
	return m
}

func (m *bool) Filter(value Query) *bool {
	m.mappings["filter"] = value
	return m
}

func (m *bool) Should(value Query) *bool {
	m.mappings["should"] = value
	return m
}

func (m *bool) ShouldNot(value Query) *bool {
	m.mappings["should_not"] = value
	return m
}

func (m *bool) Bytes() []byte {
	data := map[string]map[string]interface{}{"bool": m.mappings}
	bytes, _ := json.Marshal(data)

	return bytes
}

func (m *bool) String() string {
	return string(m.Bytes())
}
