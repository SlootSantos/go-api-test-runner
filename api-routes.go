package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type api []route

type route struct {
	name        string
	link        string
	description string
}

func (a api) test() {
	fmt.Println(a)
}

func (a api) startTest() {
	for _, x := range a {

		resp, err := http.Get("https://josh-be-hbimvzvkkk.now.sh" + x.link)
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Println(string(body))
	}
}

func newAPI() api {
	return api{
		{
			name:        "users",
			link:        "/users",
			description: "I dont really know..",
		},
		{
			name:        "index",
			link:        "/",
			description: " INDEX",
		}}
}
