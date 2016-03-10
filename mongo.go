package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("192.168.99.100:27017")
	if err != nil {
		panic(err)
	}
	// mongoDialInfo := &mgo.DialInfo{
	//                Addrs:    []string{MongoDBHosts},
	//                Timeout:  60 * time.Second,
	//                Database: MongoDatabase,
	//                Username: MongoUserName,
	//                Password: MongoPassword,
	//        }

	//        // Create a session which maintains a pool of socket connections
	//        // to our MongoDB.
	//        session, err := mgo.DialWithInfo(mongoDialInfo)
	//        if err != nil {
	//                log.Fatalf("CreateSession: %s\n", err)
	//        }

	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Gunjan", "408 901 1234"},
		&Person{"Obama", "129 123 4631"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Obama"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Phone:", result.Phone)
}
