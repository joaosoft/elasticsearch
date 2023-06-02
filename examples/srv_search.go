package main

import (
	"github.com/joaosoft/elasticsearch"

	"fmt"

	"os"
)

func searchDocument(name string) {
	var data []person

	d1 := elasticsearch.SearchTemplate{Data: map[string]interface{}{"name": name}}

	// document search
	dir, _ := os.Getwd()
	_, err := client.Search().
		Index("persons").
		Type("person").
		Object(&data).
		Template(dir+"/examples/templates", "get.example.search.template", &d1, false).
		Search()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nsearch person by name:%s %+v\n", name, data)
	}
}
