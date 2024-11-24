package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// create a structure for storing
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool
}

func getmethod() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	response, error := http.Get(myURL)

	if error != nil {
		fmt.Println("error while getting getting URL")
	}
	//make sure the required resourece used iis closed
	defer response.Body.Close()

	//we can add a status check

	if response.StatusCode != http.StatusOK {
		fmt.Println("the response wasn't correct", response.Status)
	}

	//now we tcreate an object where we want to store the value received from response
	var todo Todo
	err := json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding response")
		return
	}
	fmt.Println("Todo Id:", todo.ID)
	fmt.Println("Todo Id:", todo.Title)
	fmt.Println("Todo Id:", todo.Completed)
}

func postmethod() {
	myURL := "https://jsonplaceholder.typicode.com/todos"
	//create data for sedning to the DB
	todo := Todo{
		UserID:    1,
		ID:        2302,
		Title:     "Learn GO",
		Completed: true,
	}
	//comver data into json format as we need to dealwith web
	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("the data does not get cpnverted")
		return
	}
	//we pass data as a string so we convert t into a string
	jsonStr := string(jsonData)

	//create an io reader
	jsonReader := strings.NewReader(jsonStr)

	//send the post response--Reader is the body passed and we need to convret the string data to reader
	resp, errr := http.Post(myURL, "application/json", jsonReader)

	if errr != nil {
		fmt.Println("hehehehhe error")
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
	fmt.Println(resp.StatusCode)
}

func updatemethod() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1"
	//put is basically update in API
	todo := Todo{
		UserID:    1,
		ID:        23,
		Title:     "Learn",
		Completed: false,
	}
	//create data and conver into JOSN format by marshalling
	jsondata, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error marshaling json")
		return
	}
	jsonStr := string(jsondata)

	//create an io reader
	jsonReader := strings.NewReader(jsonStr)
	//update is basically put request
	//http doe not have put method so we need to create our own request
	//req, error := http.NewRequest(http.MethodPut, myURL, bytes.NewBuffer(jsondata))
	req, error := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if error != nil {
		fmt.Println("req failed")
		return
	}
	// when we have created our request then we need to speciffy the content type as whenever we send req we need to provide content type
	req.Header.Set("Content-Type", "application/json")
	// when ever connection is created wee do it between client and server so in this case we need to create a new client
	client := http.Client{}
	resp, er := client.Do(req)

	if er != nil {
		fmt.Println("error while sending request")
	}

	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	fmt.Println("status code", resp.Status)

}

func Deletemethod() {
	myURL := "https://jsonplaceholder.typicode.com/todos/1"
	//Create delete request
	req, err := http.NewRequest(http.MethodDelete, myURL, nil)
	if err != nil {
		fmt.Println("req failed")
		return
	}
	client := http.Client{}
	resp, errr := client.Do(req)
	if errr != nil {
		fmt.Println("response failed")
		return
	}
	fmt.Println("Response status", resp.Status)
	defer resp.Body.Close()
}
func main() {
	fmt.Println("Hi working on CRUD opereations")
	//getmethod()
	//postmethod()
	//updatemethod()
	Deletemethod()
}
