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
	splitString := strings.Split(contentString, "\n")
	window := []int{0, 0, 0}
	counter := 0
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range splitString {
		newDepth, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if window[0] > 0 {
			oldDepthSum := window[0] + window[1] + window[2]
			newDepthSum := window[1] + window[2] + newDepth
			if oldDepthSum < newDepthSum {
				counter += 1
			}
		}
		window[0] = window[1]
		window[1] = window[2]
		window[2] = newDepth
	}
	fmt.Printf("%d\n", counter)
}
