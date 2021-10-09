package main

import "fmt"

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type Post struct {
	Id               int
	Caption          string
	Image_URL        string
	Posted_Timestamp string
}

func main() {
	fmt.Println("hi")
}
