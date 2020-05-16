package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
)

var consoleMutex = &sync.Mutex{}

func fixUrl(url string) string {
	//Add http:// if it is not a well formed url
	result, err := regexp.MatchString("http(s)?://.*", url)

	if err != nil {
		log.Println(err)
		return url
	}

	if !result {
		return "http://" + url
	}
	return url
}

func request(url string, sem chan int) {
	//Check if the URL can be fixed
	url = fixUrl(url)

	//Launch the request
	resp, err := http.Get(url)

	if err != nil {
		//Indicate error and notify the channel
		fmt.Printf("Error requesting: '%s'.\n", url)
		<-sem
		return
	}

	//create the hash and write there the response
	hash := md5.New()
	resp.Write(hash)

	//Print it and notify the semaphore
	consoleMutex.Lock()
	fmt.Printf("%s %s\n", url, hex.EncodeToString(hash.Sum(nil)))
	consoleMutex.Unlock()
	<-sem
}

func main() {
	//Get the maximum number of parallel req. Default = 10.
	parallel := 10
	flag.IntVar(&parallel, "parallel", 10, "Number of allowed parallel requests")
	flag.Parse()

	if parallel == 0 {
		fmt.Println("parallel cannot be 0.")
		return
	}

	//Create the channel for parallel communication
	sem := make(chan int, parallel)

	//retrieve the arguments that are not the flag
	args := flag.Args()

	//Send the request every time there is space in the semaphore
	for _, req := range args {
		sem <- 1
		go request(req, sem)
	}

	//Check that the semaphore is empty, that means all are finished
	for i := 1; i <= parallel; i++ {
		sem <- i
	}
}
