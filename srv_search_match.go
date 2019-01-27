package elastic

import "encoding/json"

type operator string
type zeroTermsQuery string

const (
	OperatorAnd operator = "and"
	OperatorOr  operator = "or"

	ZeroTermsQueryAll  zeroTermsQuery = "all"
	ZeroTermsQueryNone zeroTermsQuery = "none"
)

type match struct {
	mappings map[string]interface{}
	field    string
}

func NewMatch(field string, query string) *match {
	new := &match{
		field: field,
	}

	new.mappings["query"] = query

	return new
}

func (m *match) Operator(value operator) *match {
	m.mappings["operator"] = value
	return m
}

func (m *match) MinimumShouldMatch(value string) *match {
	m.mappings["minimum_should_match"] = value
	return m
}

func (m *match) Analyzer(value string) *match {
	m.mappings["analyzer"] = value
	return m
}

func (m *match) Lenient(value bool) *match {
	m.mappings["lenient"] = value
	return m
}

func (m *match) Fuzziness(value string) *match {
	m.mappings["fuzziness"] = value
	return m
}

func (m *match) FuzzyRewrite(value string) *match {
	m.mappings["fuzzy_rewrite"] = value
	return m
}

func (m *match) FuzzyTranspositions(value bool) *match {
	m.mappings["fuzzy_transpositions"] = value
	return m
}

func (m *match) PrefixLength(value int64) *match {
	m.mappings["prefix_length"] = value
	return m
}

func (m *match) MaxExpansions(value int64) *match {
	m.mappings["max_expansions"] = value
	return m
}

func (m *match) ZeroTermsQuery(value zeroTermsQuery) *match {
	m.mappings["zero_terms_query"] = value
	return m
}

func (m *match) CutoffFrequency(value float64) *match {
	m.mappings["cutoff_frequency"] = value
	return m
}

func (m *match) AutoGenerateSynonymsPhraseQuery(value bool) *match {
	m.mappings["auto_generate_synonyms_phrase_query"] = value
	return m
}

func (m *match) Bytes() []byte {
	data := map[string]map[string]interface{}{"match": {m.field: m.mappings}}
	bytes, _ := json.Marshal(data)

	return bytes
}

func (m *match) String() string {
	return string(m.Bytes())
}
