package main

import (
	"log"
	"os"
)

func CheckLibFolder() (bool, error) {
	s, err := os.Stat("./lib")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return s.IsDir(), nil
}
