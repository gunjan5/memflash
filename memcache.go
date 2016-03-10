package main

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
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
)

func main() {
	mc := memcache.New("192.168.99.100:11211")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	item, _ := mc.Get("foo")
	if item != nil {
		fmt.Println(string(item.Value))
	}

	// dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
	// DB_USER, DB_PASSWORD, DB_NAME, DB_HOST)

	// db, err := sql.Open("postgres", dbinfo)
	check(err)
	// defer db.Close()
}

func check(err error) {
	if err != nil {
		panic(err)
		log.Fatalf("ERROR: %s", err)

	}

}
