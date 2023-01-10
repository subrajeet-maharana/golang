package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course -- file
type Course struct {
	CourseId   string `json:"courseid"`
	CourseName string `json:"coursename"`
	//we can toggle to receive json params by settings CoursePrice int `json:"-"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	// return c.CourseName == "" && c.CourseId == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API -- The Beast")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "C++", CoursePrice: 299, Author: &Author{Fullname: "Bjarne Straustrop", Website: "cplusplus.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "C", CoursePrice: 399, Author: &Author{Fullname: "Dennis Ritchie", Website: "c.com"}})
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	//listening
	log.Fatal(http.ListenAndServe(":5000", r))
}

// controllers -file
// serve homeroute
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API in Golang!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All Courses are ")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	//grab id from request
	params := mux.Vars(r)
	//loop through the courses,find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	//what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}
	//TODO: Check if title is duplicate or not
	//if yes, loop -->title matches -->course.CourseName --> json

	//generate unique id,string
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")
	//grab request
	params := mux.Vars(r)
	//loop --> id --> remove --> add with myId
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// loop --> id --> remove
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			//confirm or deny response
			break
		}
	}
}
