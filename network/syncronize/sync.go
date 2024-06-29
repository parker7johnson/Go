package main

import (
	"bufio"
	"fmt"
	"os"
)

//make a program that listens to user input and makes the user that just typed wait until the other user responded

type User struct {
	message    chan string
	finishType chan bool
	respond    chan bool
	err        chan string
	number     int
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	getInput()
}

func getInput() {
	user1 := User{make(chan string), make(chan bool), make(chan bool), make(chan string), 1}
	user2 := User{make(chan string), make(chan bool), make(chan bool), make(chan string), 2}
	activeUser := user1
	nextUser := user2

	handleInput(activeUser, nextUser)
}

func handleInput(activeUser, nextUser User) {
	for {
		go getMessage(activeUser)

		select {
		case err := <-activeUser.err:
			fmt.Printf("Sorry there was an error while User %d was responding: %s", activeUser.number, err)
		case msg := <-activeUser.message:
			fmt.Printf("User %d said: %s\n", activeUser.number, msg)
		}

		go respond(nextUser)
		select {
		case err := <-nextUser.err:
			fmt.Printf("Sorry there was an error while User %d was responding: %s", nextUser.number, err)
		case msg := <-nextUser.message:
			fmt.Printf("User %d responded with: %s\n", nextUser.number, msg)
		}
		activeUser, nextUser = nextUser, activeUser

	}
}

func getMessage(user User) {

	println("What would you like to say?")
	line, err := reader.ReadString('\n')

	if err != nil {
		print(err)
		user.err <- "err"
	}

	user.message <- line
}

func respond(nextUser User) {

	println("How would you like to respond?")

	line, err := reader.ReadString('\n')

	if err != nil {
		print(err)
		nextUser.err <- "err"
	}

	nextUser.message <- line
}
