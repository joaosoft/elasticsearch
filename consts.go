package elastic

const (
	DefaultURL = "http://localhost:9200"

	BulkUpsertHeader = `{ "create" : { "_index" : "test", "_type" : "_doc", "_id" : "%s" } }`
	BulkCreateHeader = `{ "create" : { "_index" : "test", "_type" : "_doc", "_id" : "%s" } }`
	BulkUpdateHeader = `{ "create" : { "_index" : "test", "_type" : "_doc", "_id" : "%s" } }`
	BulkDeleteHeader = `{ "create" : { "_index" : "test", "_type" : "_doc", "_id" : "%s" } }`
)
