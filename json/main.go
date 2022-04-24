package main

import (
	"encoding/json"
	"fmt"
)

// struct first character lower case will be used for local
// under the back-strics we can manipulate or say customize the struct variable behaviour
type course struct {
	Name     string   `json:"course-name"`
	Price    int      `json:"course-price"`
	Password string   `json:"-"`
	Platform string   `json:"website"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("hey ! creare valid json")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	courseArray := []course{
		{"Go PlayGround", 199, "testcourse@123", "www.learncode.com", []string{"backend", "web"}},
		{"python PlayGround", 199, "testcourse@123", "www.learncode.com", []string{"backend", "web", "data-science"}},
		{"Javascript PlayGround", 199, "testcourse@123", "www.learncode.com", []string{"Fullstack", "web"}},
		{"Java PlayGround", 199, "testcourse@123", "www.learncode.com", nil},
	}

	//package this data as JSON data

	finalJson, err := json.MarshalIndent(courseArray, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s/n ", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
	"course-name": "Javascript PlayGround",
	"course-price": 199,
	"website": "www.learncode.com",
	"tags": ["Fullstack","web"]
	}
	`)

	var myCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("json is valid")
		json.Unmarshal(jsonDataFromWeb, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("Json not valid")
	}

	// some caes where we just want to add data to key value pair

	var myCourseOutline map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myCourseOutline)

	fmt.Printf("%#v\n", myCourseOutline)

	for k, v := range myCourseOutline {
		fmt.Printf("key is %v and value is %v and type is  : %T\n", k, v, v)
	}

}
