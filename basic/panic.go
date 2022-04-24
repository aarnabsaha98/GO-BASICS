/*
Panic > handle exception
	even if the func is panic it will execute the defer statements and then panic
Recover

*/

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("hello Go !!"))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}
