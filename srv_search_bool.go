package elastic

import "encoding/json"

type Bool struct {
	mappings  map[string]interface{}
	must      []Query
	filter    []Query
	should    []Query
	shouldNot []Query
}

func NewBool() *Bool {
	new := &Bool{}

	return new
}

func (m *Bool) Must(value Query) *Bool {
	m.mappings["must"] = value
	return m
}

func (m *Bool) Filter(value Query) *Bool {
	m.mappings["filter"] = value
	return m
}

func (m *Bool) Should(value Query) *Bool {
	m.mappings["should"] = value
	return m
}

func (m *Bool) ShouldNot(value Query) *Bool {
	m.mappings["should_not"] = value
	return m
}

func (m *Bool) Bytes() []byte {
	data := map[string]map[string]interface{}{"bool": m.mappings}
	bytes, _ := json.Marshal(data)

	return bytes
}

func (m *Bool) String() string {
	return string(m.Bytes())
}
