package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Index string
	Data  string
}

func main() {
	session, err := mgo.Dial("192.168.99.100:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// 	jsondata := `{
	// "index": "1",
	// "index_start_at": "56",
	// "integer": "34",
	// "float": "17.2187",
	// "name": "Maxine",
	// "surname": "Chandler",
	// "fullname": "Harvey O",
	// "email": "kimberly@cooke.sb",
	// "bool": "true"
	// }`

	c := session.DB("test").C("people")
	// 	err = c.Insert(&Person{"1", jsondata})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	result := Person{}
	err = c.Find(bson.M{"index": "1"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", result.Data)
}
