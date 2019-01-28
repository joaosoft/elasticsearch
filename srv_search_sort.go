package elastic

import "encoding/json"

type Sort struct {
	mappings []*SortField
}

func NewSort(fields ...*SortField) *Sort {
	new := &Sort{
		mappings: fields,
	}

	return new
}

func (s *Sort) Bytes() []byte {
	data := map[string]interface{}{"sort": s.mappings}
	bytes, _ := json.Marshal(data)

	return bytes
}

func (s *Sort) String() string {
	return string(s.Bytes())
}
