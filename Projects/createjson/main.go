package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"` //this will hide the password
	Tags     []string `json:"tags"`
}

//encoding json
func encodejson() {
	tmcourses := []course{
		{"Golang", 1000, "tm.com", "fuytf888", []string{"web", "go"}},
		{"Python", 100, "tmhhug.com", "kkjtf884", []string{"web", "python", "ml"}},
	}

	//package this data as json data
	finaljson, err := json.MarshalIndent(tmcourses, "", "\t") // or can just use marsal
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finaljson)
}

func decodejson() {
	samplejson := []byte(` {
	"coursename": "Python",
	"price": 100,
	"website": "tmhhug.com",
	"tags": [
			"web",
			"python",
			"ml"
	]
}`)

	checkvalid := json.Valid(samplejson)
	var mycourse course
	if checkvalid {

		json.Unmarshal(samplejson, &mycourse)
	} else {
		fmt.Println("Invalid json")
	}

	// fmt.Printf("%#v\n", mycourse)

	// for using a key value pair for data
	var myjsononlinedata map[string]interface{} //we used interface coz value apart from the key can be of anytype
	json.Unmarshal(samplejson, &myjsononlinedata)
	// fmt.Printf("hello \n%#v", myjsononlinedata)
	for k, v := range myjsononlinedata {
		fmt.Printf("The key is %v and Data is %v with the type %T\n", k, v, v)
	}

}

func main() {
	fmt.Println("welcome")
	encodejson()
	decodejson()

}
