package helper

import (
	"fmt"
	"log"
	"os"
)

type errorLog struct {
}

func (e errorLog) Write(p []byte) (n int, err error) {
	fmt.Println("Error: " + string(p))
	file, _ := os.OpenFile("error.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	file.WriteString(string(p))
	// Close the file when the surrounding function exists
	defer file.Close()
	return n, err
}

// Logger is a helpper method to print out a more useful error message
var Logger = log.New(errorLog{}, "", log.Lmicroseconds|log.Lshortfile)
