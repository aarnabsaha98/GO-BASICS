package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for identity file
type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Authon      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// DB
var courses []Course

//helper
func (c *Course) IsEmpty() bool {

	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

}

// controllers

// server home route
func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> welcome to go</h1>"))
}

func getAllCouses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("GEt one Course")
	w.Header().Set("Content-Type", "application/json")

	//get ID from Request
	params := mux.Vars(r)

	fmt.Println("the request params : ", params)

	// loop

	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course.Authon.Fullname)
			return
		}
	}
	json.NewEncoder(w).Encode("No Couse found ")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one Course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data ")
	}

	// what about {}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	// 1  generate unique Id , string (random)
	// 2  append couse into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one Course")
	w.Header().Set("Content-Type", "application/json")

	// get the update ID

	params := mux.Vars(r)

	// loop through  id value  remove  and add the value with my ID

	for index, course := range courses {
		if course.CourseID == params["id"] {
			//get the perticular course
			courses = append(courses[:index], courses[index+1:]...)
			fmt.Println(courses)
			// take copy of Course struct
			var course Course
			// decode the request body
			_ = json.NewDecoder(r.Body).Decode(&course)
			// update the perticular update
			course.CourseID = params["id"]
			// put it back in bag
			courses = append(courses, course)
			//update the course
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// Todo : id not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

		}
	}
}
