package main

import (
	"fmt"
)

func bulkCreate() {
	bulk := client.Bulk()

	// document create with id
	id := "1"
	err := bulk.Index("persons").Type("person").Id(id).Body(person{
		Name: "joao",
		Age:  1,
	}).DoCreate()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nadding a new person with id %s\n", id)
	}

	id = "2"
	err = bulk.Index("persons").Type("person").Id(id).Body(person{
		Name: "tiago",
		Age:  2,
	}).DoCreate()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nadding a new person with id %s\n", id)
	}

	fmt.Println("executing bulk")
	_, err = bulk.Execute()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("success!")
	}
}

func bulkIndex() {
	bulk := client.Bulk()

	// document create with id
	id := "3"
	err := bulk.Index("persons").Type("person").Id(id).Body(person{
		Name: "joao",
		Age:  1,
	}).DoIndex()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nadding a new person with id %s\n", id)
	}

	id = "4"
	err = bulk.Index("persons").Type("person").Id(id).Body(person{
		Name: "tiago",
		Age:  2,
	}).DoCreate()

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\nadding a new person with id %s\n", id)
	}

	fmt.Println("executing bulk")
	_, err = bulk.Execute()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("success!")
	}
}

func bulkDelete() {
	bulk := client.Bulk()

	// document delete with id
	id := "1"
	err := bulk.Index("persons").Type("person").Id(id).DoDelete()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\ndeleting the person with id %s\n", id)
	}

	id = "2"
	err = bulk.Index("persons").Type("person").Id(id).DoDelete()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\ndeleting the person with id %s\n", id)
	}

	// document delete with id
	id = "3"
	err = bulk.Index("persons").Type("person").Id(id).DoDelete()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\ndeleting the person with id %s\n", id)
	}

	id = "4"
	err = bulk.Index("persons").Type("person").Id(id).DoDelete()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("\ndeleting the person with id %s\n", id)
	}

	fmt.Println("executing bulk")
	_, err = bulk.Execute()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("success!")
	}
}
