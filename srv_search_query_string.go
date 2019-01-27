package elastic

import "encoding/json"

type QueryString struct {
	mappings map[string]interface{}
	field    string
}

func NewQueryString(field string, value interface{}) *QueryString {
	new := &QueryString{
		field: field,
	}

	new.mappings["value"] = value

	return new
}

func (m *QueryString) DefaultField(value float64) *QueryString {
	m.mappings["default_field"] = value
	return m
}

func (m *QueryString) DefaultOperator(value operator) *QueryString {
	m.mappings["default_operator"] = value
	return m
}

func (m *QueryString) Analyser(value string) *QueryString {
	m.mappings["analyser"] = value
	return m
}

func (m *QueryString) QuoteAnalyzer(value string) *QueryString {
	m.mappings["quote_analyzer"] = value
	return m
}

func (m *QueryString) AllowLeadingWildcard(value bool) *QueryString {
	m.mappings["allow_leading_wildcard"] = value
	return m
}

func (m *QueryString) EnablePositionIncrements(value bool) *QueryString {
	m.mappings["enable_position_increments"] = value
	return m
}

func (m *QueryString) FuzzyMaxExpansions(value int64) *QueryString {
	m.mappings["fuzzy_max_expansions"] = value
	return m
}

func (m *QueryString) Fuzziness(value string) *QueryString {
	m.mappings["fuzziness"] = value
	return m
}

func (m *QueryString) FuzzyPrefixLength(value int64) *QueryString {
	m.mappings["fuzzy_prefix_length"] = value
	return m
}

func (m *QueryString) FuzzyTranspositions(value bool) *QueryString {
	m.mappings["fuzzy_transpositions"] = value
	return m
}

func (m *QueryString) PhraseSlop(value int64) *QueryString {
	m.mappings["phrase_slop"] = value
	return m
}

func (m *QueryString) Boost(value int64) *QueryString {
	m.mappings["boost"] = value
	return m
}

func (m *QueryString) AutoGeneratePhraseQueries(value bool) *QueryString {
	m.mappings["auto_generate_phrase_queries"] = value
	return m
}

func (m *QueryString) AnalyzeWildcard(value bool) *QueryString {
	m.mappings["analyze_wildcard"] = value
	return m
}

func (m *QueryString) MaxDeterminizedStates(value bool) *QueryString {
	m.mappings["max_determinized_states"] = value
	return m
}

func (m *QueryString) MinimumShouldMatch(value string) *QueryString {
	m.mappings["minimum_should_match"] = value
	return m
}

func (m *QueryString) Lenient(value bool) *QueryString {
	m.mappings["lenient"] = value
	return m
}

func (m *QueryString) TimeZone(value string) *QueryString {
	m.mappings["time_zone"] = value
	return m
}

func (m *QueryString) QuoteFieldSuffix(value string) *QueryString {
	m.mappings["quote_field_suffix"] = value
	return m
}

func (m *QueryString) AutoGenerateSynonymsPhraseQuery(value bool) *QueryString {
	m.mappings["auto_generate_synonyms_phrase_query"] = value
	return m
}

func (m *QueryString) AllFields(value string) *QueryString {
	m.mappings["all_fields"] = value
	return m
}

func (m *QueryString) TieBreaker(value string) *QueryString {
	m.mappings["tie_breaker"] = value
	return m
}

func (m *QueryString) Query(value string) *QueryString {
	m.mappings["query"] = value
	return m
}

func (m *QueryString) Fields(values ...string) *QueryString {
	if _, ok := m.mappings["fields"]; !ok {
		m.mappings["fields"] = make([]string, 0)
	}

	m.mappings["fields"] = append(m.mappings["fields"].([]string), values...)
	return m
}

func (m *QueryString) Bytes() []byte {
	data := map[string]map[string]map[string]interface{}{"query": {"query_string": m.mappings}}
	bytes, _ := json.Marshal(data)

	return bytes
}

func (m *QueryString) String() string {
	return string(m.Bytes())
}
