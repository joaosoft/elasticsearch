package elastic

import "encoding/json"

type queryString struct {
	mappings map[string]interface{}
	field    string
}

func NewQueryString(field string, value interface{}) *queryString {
	new := &queryString{
		field: field,
	}

	new.mappings["value"] = value

	return new
}

func (m *queryString) DefaultField(value float64) *queryString {
	m.mappings["default_field"] = value
	return m
}

func (m *queryString) DefaultOperator(value operator) *queryString {
	m.mappings["default_operator"] = value
	return m
}

func (m *queryString) Analyser(value string) *queryString {
	m.mappings["analyser"] = value
	return m
}

func (m *queryString) QuoteAnalyzer(value string) *queryString {
	m.mappings["quote_analyzer"] = value
	return m
}

func (m *queryString) AllowLeadingWildcard(value bool) *queryString {
	m.mappings["allow_leading_wildcard"] = value
	return m
}

func (m *queryString) EnablePositionIncrements(value bool) *queryString {
	m.mappings["enable_position_increments"] = value
	return m
}

func (m *queryString) FuzzyMaxExpansions(value int64) *queryString {
	m.mappings["fuzzy_max_expansions"] = value
	return m
}

func (m *queryString) Fuzziness(value string) *queryString {
	m.mappings["fuzziness"] = value
	return m
}

func (m *queryString) FuzzyPrefixLength(value int64) *queryString {
	m.mappings["fuzzy_prefix_length"] = value
	return m
}

func (m *queryString) FuzzyTranspositions(value bool) *queryString {
	m.mappings["fuzzy_transpositions"] = value
	return m
}

func (m *queryString) PhraseSlop(value int64) *queryString {
	m.mappings["phrase_slop"] = value
	return m
}

func (m *queryString) Boost(value int64) *queryString {
	m.mappings["boost"] = value
	return m
}

func (m *queryString) AutoGeneratePhraseQueries(value bool) *queryString {
	m.mappings["auto_generate_phrase_queries"] = value
	return m
}

func (m *queryString) AnalyzeWildcard(value bool) *queryString {
	m.mappings["analyze_wildcard"] = value
	return m
}

func (m *queryString) MaxDeterminizedStates(value bool) *queryString {
	m.mappings["max_determinized_states"] = value
	return m
}

func (m *queryString) MinimumShouldMatch(value string) *queryString {
	m.mappings["minimum_should_match"] = value
	return m
}

func (m *queryString) Lenient(value bool) *queryString {
	m.mappings["lenient"] = value
	return m
}

func (m *queryString) TimeZone(value string) *queryString {
	m.mappings["time_zone"] = value
	return m
}

func (m *queryString) QuoteFieldSuffix(value string) *queryString {
	m.mappings["quote_field_suffix"] = value
	return m
}

func (m *queryString) AutoGenerateSynonymsPhraseQuery(value bool) *queryString {
	m.mappings["auto_generate_synonyms_phrase_query"] = value
	return m
}

func (m *queryString) AllFields(value string) *queryString {
	m.mappings["all_fields"] = value
	return m
}

func (m *queryString) TieBreaker(value string) *queryString {
	m.mappings["tie_breaker"] = value
	return m
}

func (m *queryString) Query(value string) *queryString {
	m.mappings["query"] = value
	return m
}

func (m *queryString) Fields(values ...string) *queryString {
	if _, ok := m.mappings["fields"]; !ok {
		m.mappings["fields"] = make([]string, 0)
	}

	m.mappings["fields"] = append(m.mappings["fields"].([]string), values...)
	return m
}

func (m *queryString) Bytes() []byte {
	data := map[string]map[string]map[string]interface{}{"query": {"query_string": m.mappings}}
	bytes, _ := json.Marshal(data)

	return bytes
}

func (m *queryString) String() string {
	return string(m.Bytes())
}
