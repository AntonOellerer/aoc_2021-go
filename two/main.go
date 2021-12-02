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
	depth := 0
	aim := 0
	for _, commandTuple := range commandString {
		commandTuple := strings.Split(commandTuple, " ")
		command := commandTuple[0]
		units, err := strconv.Atoi(commandTuple[1])
		if err != nil {
			log.Fatal(err)
		}
		if command == "forward" {
			x += units
			depth += aim * units
		} else if command == "down" {
			aim += units
		} else if command == "up" {
			aim -= units
		}
	}
	fmt.Printf("%d \n", x*depth)
}
