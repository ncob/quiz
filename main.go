package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Kudos:
// https://tutorialedge.net/golang/reading-console-input-golang/

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Could not open file")
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Couldn't read file")
	}
	r := csv.NewReader(bytes.NewReader(b))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	correct := 0
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(records); i++ {
		fmt.Println(records[i][0])
		answer, _ := reader.ReadString('\n')
		answer = strings.Replace(answer, "\n", "", -1)

		if records[i][1] == answer {
			correct++
		}
	}

	fmt.Printf("%d out of %d correct", correct, len(records))
}
