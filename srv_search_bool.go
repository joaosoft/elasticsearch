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
	new := &Bool{
		mappings: make(map[string]interface{}),
	}

	return new
}

func (b *Bool) Must(value ...Query) *Bool {
	if len(value) > 0 {
		b.must = append(b.must, value...)
		b.mappings["must"] = b.must
	}
	return b
}

func (b *Bool) Filter(value ...Query) *Bool {
	if len(value) > 0 {
		b.filter = append(b.filter, value...)
		b.mappings["filter"] = b.filter
	}
	return b
}

func (b *Bool) Should(value ...Query) *Bool {
	if len(value) > 0 {
		b.should = append(b.should, value...)
		b.mappings["should"] = b.should
	}
	return b
}

func (b *Bool) ShouldNot(value ...Query) *Bool {
	if len(value) > 0 {
		b.shouldNot = append(b.shouldNot, value...)
		b.mappings["should_not"] = b.shouldNot
	}
	return b
}

func (b *Bool) Data() interface{} {
	data := map[string]interface{}{"bool": b.mappings}
	return data
}

func (b *Bool) Bytes() []byte {
	bytes, _ := json.Marshal(b.Data())

	return bytes
}

func (b *Bool) String() string {
	return string(b.Bytes())
}
