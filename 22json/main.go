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
	EncodeJson()
	DecodeJson()

	//making slice using the course struct

}
func EncodeJson() {
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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}

}
