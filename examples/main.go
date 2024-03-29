package main

import (
	"fmt"
	"github.com/joaosoft/elasticsearch"
	"time"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// var client = elastic.NewElastic()
// you can define the configuration without having a configuration file
var config, _, _ = elasticsearch.NewConfig("localhost:9200")
var client, _ = elasticsearch.NewElastic(elasticsearch.WithConfiguration(config.Elastic))

func main() {

	// delete indexes
	fmt.Println(":: DELETE INDEX")
	deleteIndex()

	// index exists
	fmt.Println(":: EXISTS INDEXES ?")
	existsIndex("persons")
	existsIndex("bananas") // don't exist

	// index create with mapping
	fmt.Println(":: CREATE INDEX WITH MAPPING")
	createIndexWithMapping()

	// document create
	fmt.Println(":: CREATE DOCUMENT 1")
	createDocumentWithId("1")
	fmt.Println(":: CREATE DOCUMENT 2")
	createDocumentWithId("2")
	fmt.Println(":: CREATE DOCUMENT WITHOUT ID")
	generatedId := createDocumentWithoutId()

	// document update
	fmt.Println(":: UPDATE DOCUMENT 1")
	updateDocumentWithId("1")
	fmt.Println(":: UPDATE DOCUMENT 2")
	updateDocumentWithId("2")

	// document search
	// wait elastic to index the last update...
	fmt.Println(":: SEARCH DOCUMENT  WITH 'luis'")
	<-time.After(time.Second * 2)
	searchDocument("luis")

	// count index documents
	fmt.Println(":: COUNT DOCUMENTS ON INDEX WITH 'luis'")
	countOnIndex("luis")
	fmt.Println(":: COUNT DOCUMENTS ON DOCUMENT WITH 'luis'")
	countOnDocument("luis")

	// document delete
	fmt.Println(":: COUNT DOCUMENT WITH GENERATED ID")
	deleteDocumentWithId("1")
	deleteDocumentWithId("2")
	deleteDocumentWithId(generatedId)

	// bulk
	fmt.Println(":: BULK OPERATIONS")
	bulkCreate()
	bulkIndex()
	bulkDelete()

	// queue bulk create
	fmt.Println(":: QUEUE BULK CREATE")
	queueBulkCreate()

	// index delete
	fmt.Println(":: DELETE INDEX")
	deleteIndex()

}
