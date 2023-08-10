package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

/// Goroutines

func getReponse01(url string) {
	fmt.Println("step1 : ", url)
	res, error := http.Get(url)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("step2 : ", url)
	defer res.Body.Close()

	fmt.Println("step3 : ", url)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("step4 : ", len(body))

}

func main01() {
	go getReponse01("https://www.golangprograms.com")
	go getReponse01("https://coderwall.com")
	go getReponse01("https://stackoverflow.com")
	time.Sleep(10 * time.Second)
}

// Waiting for Goroutines to Finish Execution
var wg sync.WaitGroup

func getReponse02(url string) {

	defer wg.Done()

	fmt.Println("step1 : ", url)
	res, error := http.Get(url)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("step2 : ", url)
	defer res.Body.Close()

	fmt.Println("step3 : ", url)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("step4 : ", url, len(body))

}

func main() {
	wg.Add(3)

	go getReponse02("https://www.golangprograms.com")
	go getReponse02("https://coderwall.com")
	go getReponse02("https://stackoverflow.com")

	wg.Wait()
	fmt.Println("Terminating Program")
}
