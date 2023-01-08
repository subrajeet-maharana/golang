package main

import (
	"fmt"
	"net/url"
)

const myUrl = "http://lco.dev:3000/learn?coursename=golang&paymentid=ksmdlskdlfg"

func main() {
	fmt.Println("Welcome to URLs in golang")
	fmt.Println(myUrl)

	//Parsing the url and storing the result
	result, err := url.Parse(myUrl)
	catchingNilErr(err)

	fmt.Println(result)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//Taking all the query parameters in the url
	qparams := result.Query()
	fmt.Printf("The type of query parameters are %T\n", qparams) //Result will be url.Valuess which means it is a key-value pair so we can use all the operations related to the key value pairs
	fmt.Println(qparams["coursename"])

	//printing all the query parameters
	//Here we don't need index
	for _, value := range qparams {
		fmt.Println("Param is: ", value)
	}

	//Making a url from the constituents
	//To make the url we have to use the url package. We MUST pass the reference (&) here as we want to send th actual url not a copy of the url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=Subrajeet",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}

func catchingNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
