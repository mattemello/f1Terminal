package errorsh

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
)

func OpenFileLog() *os.File {
	if _, err := os.Open("/tmp/f1Terminal"); os.IsNotExist(err) {
		err := os.Mkdir("/tmp/f1Terminal", 0755)
		AssertNilShutDown(err, "Error in the open of the log directory")
	}

	fileLog, err := os.OpenFile("/tmp/f1Terminal/logfile", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	AssertNilShutDown(err, "Error in the open of the log file")

	return fileLog
}

func AssertNilJson(err error, body []byte) {
	if err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			fmt.Printf("syntax error at byte offset %d\n", e.Offset)
		}
		fmt.Printf("error in the unmarshal: %s \n\nbody:\n %s", err, string(body))

		os.Exit(1)
	}
}

// This is for the minor error that don't need to shut down the program and write all in a file
func AssertNilFile(err error, txt string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("file: %s, line: %d - %s: %s", file, line, txt, err)
	}
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
