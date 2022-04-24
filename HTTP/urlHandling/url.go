package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursname=reactjs&paymentid=3er4tt"

func main() {
	fmt.Println("Welcome to handling urls", myUrl)

	//parsing

	res, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Hostname())
	fmt.Println(res.Path)
	fmt.Println(res.RawQuery)

	qparams := res.Query()

	fmt.Printf("qparams is of type %T\n", qparams)
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("params :", val)
	}

	// constract url

	partsOfURL := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
	}
	fmt.Println(partsOfURL)

}
