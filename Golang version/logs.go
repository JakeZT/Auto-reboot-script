package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func checkFile(Filename string) bool {
	var exist = true
	if _, err := os.Stat(Filename); os.IsNotExist(err) {
		exist = false
		if err != nil {
			fmt.Println("Log File not exist, it will be created later.")
		}
	}
	return exist
}

func writeToLogFile(str string) {
	var f1 *os.File
	var err1 error
	formatedStr := str + "\n" + time.Now().Format(time.RFC3339) + "\n"
	Filenames := "./logs.log"

	if checkFile(Filenames) {
		f1, err1 = os.OpenFile(Filenames, os.O_APPEND|os.O_WRONLY, 0666)
		if err1 != nil {
			fmt.Println("File exist, opening now.")
		}

	} else {
		f1, err1 = os.Create(Filenames)
		if err1 != nil {
			fmt.Println("Failed to create a log file")
		}
		fmt.Println("Successfully created a log file")
	}
	_, err1 = io.WriteString(f1, formatedStr)
	if err1 != nil {
		fmt.Println(err1)
	}
	defer f1.Close()
}
