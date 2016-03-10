package memflash

import (
	//"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	// "log"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

type DB struct {
	Mem, Mongo string
	Mc         *memcache.Client
}

func (d *DB) New() *memcache.Client {
	d.Mc = memcache.New(d.Mem)

	return d.Mc

}
