package main

import (
	"fmt"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/gunjan5/memflash/memflash"
	// "database/sql"
	//_ "github.com/lib/pq"
	"log"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	in:=""
	memCh := make(chan bool)
	mongoCh := make(chan bool)

	//mc := memcache.New("192.168.99.100:11211")
	//fmt.Printf("mc: %T\n", mc)
	mydb := memflash.DB{MEM_IP + MEM_PORT, "", nil}

	ref := mydb.New()

	session, err := mgo.Dial(MONGO_IP + MONGO_PORT)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	result := Person{}

	start := time.Now()


	for i := 0; i < 10; i++ {
		start1 := time.Now()


		go func(i int) {

			item, err := ref.Get(strconv.Itoa(i))
		if err == memcache.ErrCacheMiss {
			//fmt.Println("CACHE MISS!!!!!")
			err = c.Find(bson.M{"index": strconv.Itoa(i)}).One(&result)
		//fmt.Println(err)
		check(err)

		//log.Println("Mongo Result:", result.Data)
		//fmt.Println("MONGO RES ~~~~~~~~~~~~~~")
		mongoCh <- true
		}
		if item != nil {
			//log.Println("*******************  Mem Result:", string(item.Value))
			//fmt.Println("HIT")

		}
		memCh<- true
		}(i)

		// go func(i int) {
		// 	err = c.Find(bson.M{"index": strconv.Itoa(i)}).One(&result)
		// //fmt.Println(err)
		// check(err)

		// //log.Println("Mongo Result:", result.Data)
		// //fmt.Println("MONGO RES ~~~~~~~~~~~~~~")
		// mongoCh <- true

		// }(i)

		select{
		case <-memCh:
			fmt.Println("MEMCACHE WINS!!")
			fmt.Printf("* time %s \n ", time.Since(start1))
		case <-mongoCh:
			fmt.Println("MONGO wins")
			fmt.Printf("* time %s \n ", time.Since(start1))

		}

	}

	fmt.Printf("****** Executed in time %s \n ", time.Since(start))

	fmt.Scanf("%s", &in)

}

func check(err error) {
	if err != nil {
		//panic(err)
		log.Fatalf("ERROR: %s", err)

	}

}
