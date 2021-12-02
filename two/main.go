package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	contentString := string(content)
	commandString := strings.Split(contentString, "\n")
	x := 0
	y := 0
	for _, commandTuple := range commandString {
		commandTuple := strings.Split(commandTuple, " ")
		command := commandTuple[0]
		moves, err := strconv.Atoi(commandTuple[1])
		if err != nil {
			log.Fatal(err)
		}
		if command == "forward" {
			x += moves
		} else if command == "down" {
			y += moves
		} else if command == "up" {
			y -= moves
		}
	}
	fmt.Printf("%d \n", x*y)
}
