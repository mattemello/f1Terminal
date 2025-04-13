package errorsh

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

const path = "./tmp"

var f *os.File

func OpenFileLog() {
	if _, err := os.Open(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		AssertNilShutDown(err, "Error in the open of the log directory")
	}

	var err error
	f, err = os.OpenFile(path+"/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	// fileLog, err = os.OpenFile(path+"/log.txt", os.O_APPEND|os.O_RDONLY|os.O_CREATE, 0666)
	AssertNilShutDown(err, "Error in the open of the log file")

	log.SetOutput(f)
}

func ClearLogFile() {
	err := os.Truncate(path+"/log", 0)
	AssertNilShutDown(err, "Couldent clear the log file")
}

func CloseFile() {
	f.Close()
}

func AssertNilJson(err error, body []byte) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		if e, ok := err.(*json.SyntaxError); ok {
			AssertNilFile(errors.New("Sintax error"), fmt.Sprintf("syntax error at byte offset %d\n", e.Offset))
			// fmt.Printf("syntax error at byte offset %d\n", e.Offset)
		}
		AssertNilFile(err, fmt.Sprintf("file: %s, line: %d \n\n error in the unmarshal: %s \n\nbody:\n %s", file, line, err, string(body)))
		// fmt.Printf("error in the unmarshal: %s \n\nbody:\n %s", err, string(body))

		return true
	}

	return false
}

// This is for the minor error that don't need to shut down the program and write all in a file
// Return true if there is an err, false if there isn't.
func AssertNilFile(err error, txt string) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)

		log.Printf("file: %s, line: %d - %s: %s\n", file, line, txt, err)
		log.Printf("\n------------------------------------------------------------------------------------\n")

		return true
	}

	return false
}

// This is for the error that need to close the program and need to send a feedback to the user
func AssertNilTer(err error, txt string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("file: %s, line: %d - %s: %s", file, line, txt, err)
		AssertNilFile(err, txt)
		os.Exit(1)
	}
}

// This is for the error that need to close the program before the file is opened
func AssertNilShutDown(err error, txt string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("file: %s, line: %d - %s: %s", file, line, txt, err)
		os.Exit(1)
	}
}

func AssertNotAppening(ok bool, txt string) {
	if ok {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("file: %s, line: %d - %s\n", file, line, txt)
		log.Printf("file: %s, line: %d - %s\n", file, line, txt)
		os.Exit(1)
	}
}
