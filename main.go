package main

import (
	"fmt"
	"os"
	"strconv"
	_container "tutorials-udemy/packages"
	_hello "tutorials-udemy/src"
)

func main() {
	_hello.SayHello("thareq")
	// _ := _database.GetDatabase()
	// args := os.Args
	// fmt.Println(args)
	i, _ := strconv.Atoi("-42")
	a, _ := strconv.Atoi("-42")

	fmt.Println(i, a)

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	res, err := strconv.ParseBool("false")
	fmt.Println(err)
	fmt.Println(res)

	_container.Container()

	fmt.Println(username, password)
}
