package main

import (
	"log"
	"os"
)

func CreateConstant(name string) {
	if libExist, err := CheckLibFolder(); err == nil && libExist {
		checkConstantsFolder()
		createConstantFile(name)
	}
}

func createConstantFile(name string) {
	constantCode := []byte(
		`class ` + name + `Const {
  const ` + name + `Const._();

  // example static const main = 'example-variable';
}
`)

	if err := os.WriteFile("./lib/constants/"+name+"_const.dart", constantCode, 0644); err != nil {
		log.Fatal(err)
	}
}

func checkConstantsFolder() {
	_, err := os.Stat("./lib/constants")
	if err != nil {
		if err := os.Mkdir("./lib/constants", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}
