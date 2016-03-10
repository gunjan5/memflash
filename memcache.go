package memcache


import (
	"fmt"
	"log"
	"github.com/bradfitz/gomemcache/memcache"
)

type Ref struct {
	memC *memcache.Client
	//mgoC *

}

func New(memcacheIP, mongoIP string) *Ref {
	memcache.New(memcacheIP)
	
	
}