package main

import (
	"github.com/bradfitz/gomemcache/memcache"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//io.WriteString(w, "Hello world!")
	mc := memcache.New("192.168.99.100:11211")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	item, _ := mc.Get("foo")
	if item != nil {
		io.WriteString(w, string(item.Value))
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}
