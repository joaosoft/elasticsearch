package elastic

const (
	DefaultURL = "http://localhost:9200"

	BulkUpsertHeader = `{ "create" : { "_index" : "%s", "_type" : "%s", "_id" : "%s" } }`
	BulkCreateHeader = `{ "create" : { "_index" : "%s", "_type" : "%s", "_id" : "%s" } }`
	BulkUpdateHeader = `{ "create" : { "_index" : "%s", "_type" : "%s", "_id" : "%s" } }`
	BulkDeleteHeader = `{ "create" : { "_index" : "%s", "_type" : "%s", "_id" : "%s" } }`
)
