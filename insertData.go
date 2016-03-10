package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gunjan5/memflash/memflash"
	"gopkg.in/mgo.v2"
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

	mydb := memflash.DB{MEM_IP + MEM_PORT, "", nil}

	ref := mydb.New()

	ref.Set(&memcache.Item{Key: "foo", Value: []byte(`{
"index": "1",
"index_start_at": "56",
"integer": "34",
"float": "17.2187",
"name": "Maxine",
"surname": "Chandler",
"fullname": "Harvey O",
"email": "kimberly@cooke.sb",
"bool": "true"
}`)})

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < 10000; i++ {

		jsondata := fmt.Sprintf(
			`{
	"index": "%d",
	"index_start_at": "%d",
	"integer": "34",
	"float": "17.2187",
	"name": "%s",
	"surname": "Chandler",
	"fullname": "Harvey O",
	"email": "kimberly@cooke.sb",
	"bool": "true"
	}`, i, r1.Intn(100), rune(65+r1.Intn(26)))

		fmt.Println(jsondata)

		c := session.DB("test").C("people")
		err = c.Insert(&Person{strconv.Itoa(i), jsondata})
		if err != nil {
			log.Fatal(err)
		}

		if i%2 == 0 {

			ref.Set(&memcache.Item{Key: strconv.Itoa(i), Value: []byte(jsondata)})

		}

	}

}
