package main

import (
	"log"
	"os"
	"strings"
)

func createfolder(file string) {

	if _, err := os.Stat(file); !os.IsNotExist(err) {
		log.Println(file + " folder already exists")
	} else {
		os.Mkdir(file, 0700)
	}
}

func sortfile(file string, dir string) {

	currentdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	os.Rename(file, (currentdir+"/"+dir+"/")+file)
	log.Println(file + " sorted to proper directory.")
}

func directorycheck(file string) bool {
	fileInfo, err := os.Stat(file)
	if err != nil {
		// error handling
	}

	if fileInfo.IsDir() {
		return true

	} else {
		return false
	}
}

func main() {

	// If the file doesn't exist, create it or append to the file
	logfile, err := os.OpenFile("sortlogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logfile)

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		dircheck := directorycheck(e.Name())
		if (e.Name() == "sort.exe") || (e.Name() == "sort") || (e.Name() == "sortlogs.txt") || (e.Name() == "sort.go") {
			continue
		} else if dircheck == true {
			log.Println(e.Name() + " is a directory. Skipping.")
			continue
		} else {

			array := strings.Split(e.Name(), ".")
			length := len(array)
			if length > 1 {
				final := (array[(length - 1)])
				log.Println(e.Name() + " is a ." + final + " file")
				createfolder(final)
				sortfile(e.Name(), final)
			}
		}
	}
}
