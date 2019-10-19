package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const imgPath string = "PATH OF IMAGES STORAGE DIRECTORY"
const csvPath string = "PATH OF CSV FILES FROM EXPORT TOOLS"
const numberOfWorker int = 10
const urlIndex int = 12
const nameIndex int = 32
const reportRate = 100 // report progress every 100 download
var r, _ = regexp.Compile("[0-9]+$")

type data struct {
	url  string
	name string
}

func main() {
	input := make(chan data, 10)
	output := make(chan bool, 10)
	var counter int = 0 // use to name downloaded image files
	var queue []data

	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(csvFile)

	/* read csv file */
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		/* get rid of lichen from fungi dataset */
		if !strings.Contains(line[nameIndex], "lichen") {
			queue = append(queue, data{
				url:  line[urlIndex],
				name: line[nameIndex],
			})
		}
	}

	queue = queue[1:] // remove title row

	/* init worker */
	for i := 0; i < numberOfWorker; i++ {
		url := queue[0]
		queue = queue[1:]

		counter++
		input <- url
		go worker(input, output)
	}

	for i := 0; i < len(queue); i++ {
		data := queue[i]

		<-output
		input <- data
		go worker(input, output)

		if i%reportRate == 0 {
			fmt.Println(i)
		}
	}
}

func worker(input chan data, done chan bool) {
	data := <-input
	url := data.url
	name := data.name

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	path := imgPath + name + "/"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	index := r.FindStringSubmatch(url)

	file, err := os.Create(path + strings.Join(index, "") + ".jpg")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	done <- true
}
