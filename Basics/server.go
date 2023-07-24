package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func server() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// server2 emulates an echo and counters server
package main

var mu sync.Mutex
var count int

func server2() {
	http.HandleFunc("/",handler2)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8001",nil))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n",r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// this handler echos HTTP request
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w. "Header[%q] = %q\n",k,v)
	}
	fmt.Fprintf(w, "Host = %q\n",r.Host)
	fmt.Fprintf(w,"RemoteAddr = %q\n",r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w,"Form[%q] = %q\n",k,v)
	}
}

func main() {
	fmt.Println("Echo url request's Path component")
	server()
}
