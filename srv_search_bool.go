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

func (m *Bool) Must(value ...Query) *Bool {
	m.must = append(m.must, value...)
	m.mappings["must"] = m.must
	return m
}

func (m *Bool) Filter(value ...Query) *Bool {
	m.filter = append(m.filter, value...)
	m.mappings["filter"] = m.filter
	return m
}

func (m *Bool) Should(value ...Query) *Bool {
	m.should = append(m.should, value...)
	m.mappings["should"] = m.should
	return m
}

func (m *Bool) ShouldNot(value ...Query) *Bool {
	m.shouldNot = append(m.shouldNot, value...)
	m.mappings["should_not"] = m.shouldNot
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
