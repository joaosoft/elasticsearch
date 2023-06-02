package main

import (
	"github.com/joaosoft/elasticsearch"

	"fmt"

	"os"
)

func countOnIndex(name string) int64 {

	d1 := elasticsearch.CountTemplate{Data: map[string]interface{}{"name": name}}

	// index count
	dir, _ := os.Getwd()
	response, err := client.Search().
		Index("persons").
		Template(dir+"/examples/templates", "get.example.count.template", &d1, false).
		Count()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\ncount persons with name %s: %d\n", name, response.Count)
	}

	return response.Count
}

func countOnDocument(name string) int64 {

	d1 := elasticsearch.CountTemplate{Data: map[string]interface{}{"name": name}}

	// index count
	dir, _ := os.Getwd()
	response, err := client.Search().
		Index("persons").
		Type("person").
		Template(dir+"/examples/templates", "get.example.count.template", &d1, false).
		Count()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\ncount persons with name %s: %d\n", name, response.Count)
	}

	return response.Count
}
