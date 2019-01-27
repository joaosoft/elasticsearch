package main

import (
	"elastic"

	"fmt"

	"os"

	log "github.com/joaosoft/logger"
)

func countOnIndex(name string) int64 {

	d1 := elastic.CountTemplate{Data: map[string]interface{}{"name": name}}

	// index count
	dir, _ := os.Getwd()
	response, err := client.Search().
		Index("persons").
		Template(dir+"/examples/templates", "get.example.count.template", &d1, false).
		Count()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncount persons with name %s: %d\n", name, response.Count)
	}

	return response.Count
}

func countOnDocument(name string) int64 {

	d1 := elastic.CountTemplate{Data: map[string]interface{}{"name": name}}

	// index count
	dir, _ := os.Getwd()
	response, err := client.Search().
		Index("persons").
		Type("person").
		Template(dir+"/examples/templates", "get.example.count.template", &d1, false).
		Count()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncount persons with name %s: %d\n", name, response.Count)
	}

	return response.Count
}
