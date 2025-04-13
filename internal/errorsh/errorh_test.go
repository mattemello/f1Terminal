package errorsh

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestOpenFile(t *testing.T) {
	ClearLogFile()
	OpenFileLog()

	AssertNilFile(errors.New("test"), "This is the first test")
	file, err := os.Open(path + "/log")
	if err != nil {
		fmt.Println("ops, this is not working")
	}

	by, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("ops, this is not working2")
	}
	if strings.Contains(string(by), "file: /home/matteo/code/f1Terminal/internal/errorsh/errorsh_test.go, line: 16 - This is the first test: test\n") {
		t.Fatal("The content of the file is not what you expected")
	}

	CloseFile()
}
