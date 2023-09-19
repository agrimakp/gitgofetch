package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type githubUser struct {
	Login     string
	Name      string
	Followers int `json:"followers"`
	Node_id   string
}

func main() {
	var user githubUser
	arg := os.Args[1]

	url := "https://api.github.com/users/" + arg
	fmt.Println("Hello", arg, url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//todo i need to learn this

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status) //200 succ

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &user)

	if err != nil {
		panic(err)
	}

	fmt.Println("Name:  " + user.Name)
	fmt.Println("Login:  " + user.Login)
	fmt.Println("Followers :  ", user.Followers)
}
