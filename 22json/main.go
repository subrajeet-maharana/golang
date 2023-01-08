package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //The json:"coursename" is the alias which is used as the key while creating the json
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //- inside the json alias makes sure that the password is not shown in the json
	Tags     []string `json:"tags,omitempty"` //omitempty property makes sure that the key value pair is not shown if its value is nil
}

func main() {
	fmt.Println("Welcome to JSON in golang")

	//making slice using the course struct

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learncodeonline.com", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "Learncodeonline.com", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "Learncodeonline.com", "efg123", nil},
	}

	//Packaging this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //MarchalIndent() takes 3 arguments i.e 1. data that is to be packages into JSON 	2. prefx (what should be before every line of the json) 	3. Indent (how the data inside the json is separated)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}
