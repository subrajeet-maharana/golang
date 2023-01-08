package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Welcome to web request verbs")
	fmt.Println("*********** GET Request ***********")
	PerformGetRequest()
	fmt.Println("*********** POST Request ***********")
	PerformPostJsonRequest()
	fmt.Println("*********** POSTFORM Request ***********")
	performPostFormRequest()

}

// Function to performe get request (If we want it to be public then the first letter be capital)
func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	checkingNilErr(err)

	//It is our responsibility to close the conection at the end so we have to end the connection with defer
	defer response.Body.Close()

	fmt.Println("Status code is", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	//This is the simple way of reading the content got in the response
	// content, err := ioutil.ReadAll(response.Body)	//Reading the response using the ioutil package
	// checkingNilErr(err)

	// fmt.Println(string(content))

	//We have better way to read the content in the response, by using the Strings package	('string' is a dattype but 'strings' is a package)
	var responseString strings.Builder //Making a string builder using the strings package
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content) //Using write method we are writing the content inside the responseString variable. Now we can do whatever we want with this variable in which the content has been written. We can print it in a string form also
	fmt.Println("Byte count is", byteCount)

	fmt.Println(responseString.String()) //Printing the messege in the responseString as a string

}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	//Creating a http payload (another name for data that to be sent)
	requestBody := strings.NewReader(`
		{
			"Name" : "Subrajeet",
			"Age" : 22,
			"Country" : "India"
			
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	checkingNilErr(err)

	//We MUST close the connection at the end so we are closing it with defer
	defer response.Body.Close()

	//Reading the response using the strings package
	var responseBody strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseBody.Write(content)
	fmt.Println("Bytecount is:", byteCount)
	fmt.Println(responseBody.String())

}

func performPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	//Form data
	data := url.Values{}
	data.Add("FirstName", "Subrajeet")
	data.Add("LastName", "Maharana")
	data.Add("Email", "subrajeet.info@gmail.com")

	response, err := http.PostForm(myUrl, data)
	checkingNilErr(err)

	//Closing the connection
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkingNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
