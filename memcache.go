package main

import (
	"fmt"

	//"github.com/bradfitz/gomemcache/memcache"
	"github.com/gunjan5/memflash/memflash"
	// "database/sql"
	//_ "github.com/lib/pq"
	"log"
	//"time"
)

const (
	DB_USER     = "admin"
	DB_PASSWORD = "mypassword"
	DB_NAME     = "test"
	DB_HOST     = "192.168.99.100"
	MEM_IP      = "192.168.99.100"
	MEM_PORT    = ":11211"
	MONGO_IP    = "192.168.99.100"
	MONGO_PORT  = ":27017"
)

func main() {

	//mc := memcache.New("192.168.99.100:11211")
	//fmt.Printf("mc: %T\n", mc)
	mydb := memflash.DB{MEM_IP + MEM_PORT, "", nil}

	ref := mydb.New()

	// 	ref.Set(&memcache.Item{Key: "foo", Value: []byte(`{
	// "index": "1",
	// "index_start_at": "56",
	// "integer": "34",
	// "float": "17.2187",
	// "name": "Maxine",
	// "surname": "Chandler",
	// "fullname": "Harvey O",
	// "email": "kimberly@cooke.sb",
	// "bool": "true"
	// }`)})

	item, _ := ref.Get("0")
	if item != nil {
		fmt.Println("Result:", string(item.Value))
	}

}

func check(err error) {
	if err != nil {
		panic(err)
		log.Fatalf("ERROR: %s", err)

	}

}
