package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	const uploadedDir = "uploaded/"
	const imagesDir = "images/"

	err := os.RemoveAll(uploadedDir)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Mkdir(uploadedDir, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}

	files, err := ioutil.ReadDir(imagesDir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		originalImg, err := os.Open(imagesDir + file.Name())
		newImg, err := os.Create(uploadedDir + file.Name())

		if err != nil {
			log.Fatalln(err)
		}
		bytesWritten, err := io.Copy(newImg, originalImg)
		fmt.Printf("Bytes Written: %d\n", bytesWritten)

		if err != nil {
			log.Fatalln(err)
		}
	}
}
