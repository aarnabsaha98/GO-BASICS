/*
Defer > to delayed the execution , actually it is called when the main func executes
         then it start looking for defer and executes it (LIFO)
*/

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/2")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
