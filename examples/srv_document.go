package main

import (
	"fmt"

	"strconv"

	log "github.com/joaosoft/logger"
)

func createDocumentWithId(id string) {

	// document create with id
	age, _ := strconv.Atoi(id)
	response, err := client.Document().Index("persons").Type("person").Id(id).Body(person{
		Name: "joao",
		Age:  age + 20,
	}).Create()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated a new person with id %s\n", response.ID)
	}
}

func createDocumentWithoutId() string {

	// document create without id
	response, err := client.Document().Index("persons").Type("person").Body(person{
		Name: "joao",
		Age:  30,
	}).Create()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ncreated a new person with id %s\n", response.ID)
	}

	return response.ID
}

func updateDocumentWithId(id string) {

	// document update with id
	age, _ := strconv.Atoi(id)
	response, err := client.Document().Index("persons").Type("person").Id(id).Body(person{
		Name: "luis",
		Age:  age + 20,
	}).Update()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nupdated person with id %s\n", response.ID)
	}
}

func deleteDocumentWithId(id string) {

	response, err := client.Document().Index("persons").Type("person").Id(id).Delete()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleted person with id %s\n", response.ID)
	}
}
