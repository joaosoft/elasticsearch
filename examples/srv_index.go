package main

import (
	"fmt"

	log "github.com/joaosoft/logger"
)

func createIndexWithMapping() {
	// you can define the configuration without having a configuration file
	//client1 := elastic.NewElastic(elastic.WithConfiguration(elastic.NewConfig("http://localhost:9200")))

	response, err := client.Index().Index("persons").Body([]byte(`
{
  "mappings": {
    "person": {
      "properties": {
        "age": {
          "type": "long"
        },
        "name": {
          "type": "text",
          "fields": {
            "keyword": {
              "type": "keyword",
              "ignore_above": 256
            }
          }
        }
      }
    }
  }
}
`)).Create()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated mapping for persons index ok: %t\n", response.Acknowledged)
	}
}

func deleteIndex() {

	response, err := client.Index().Index("persons").Delete()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleted persons index ok: %t\n", response.Acknowledged)
	}
}

func existsIndex(index string) {

	exists, err := client.Index().Index(index).Exists()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nexists index? %t\n", exists)
	}
}
