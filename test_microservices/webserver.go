package main
import (
    "net/http"
    "log"
    "io/ioutil"
)
	
func main() {
    http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
	log.Println("Hello, World!!")
	d,_  := ioutil.ReadAll(r.body)
	log.Printf("Data %s\n",d)
    })
    http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	log.Println("Goobye Corona!!")
    })
    http.ListenAndServe(":9090", nil)
}
