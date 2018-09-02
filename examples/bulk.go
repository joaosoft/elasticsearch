package main

import (
	"fmt"

	log "github.com/joaosoft/logger"
)

func bulkCreate() {
	bulk := client.Bulk()

	// document create with id
	id := "1"
	err := bulk.Index("persons").Type("person").Id(id).Body(person{
		Name: "joao",
		Age:  1,
	}).Create()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nadding a new person with id %s\n", id)
	}

	id = "2"
	err = bulk.Index("persons").Type("person").Id(id).Body(person{
		Name: "tiago",
		Age:  2,
	}).Create()

	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\nadding a new person with id %s\n", id)
	}

	fmt.Println("executing bulk")
	result, err := bulk.Execute()
	if err != nil {
		log.Error(err)
		fmt.Printf("%+v", result)
	} else {
		fmt.Printf("success with %+v", result)
	}
}

func bulkDelete() {
	bulk := client.Bulk()

	// document delete with id
	id := "1"
	err := bulk.Index("persons").Type("person").Id(id).Delete()
	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleting the person with id %s\n", id)
	}

	id = "2"
	err = bulk.Index("persons").Type("person").Id(id).Delete()
	if err != nil {
		log.Error(err)
	} else {
		fmt.Printf("\ndeleting the person with id %s\n", id)
	}

	fmt.Println("executing bulk")
	result, err := bulk.Execute()
	if err != nil {
		log.Error(err)
		fmt.Printf("%+v", result)
	} else {
		fmt.Printf("success with %+v", result)
	}
}
