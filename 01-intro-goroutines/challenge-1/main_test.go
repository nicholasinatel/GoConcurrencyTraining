package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	// Program prints to the Stdout
	// save standard output to the variable
	stdOut := os.Stdout

	// Read & Write from os.Pipe()
	os.Pipe()
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	// Hence this is a test, I can only add 1 waitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	// Call to testing function
	go updateMessage("this is a test!", &wg)
	wg.Wait()

	printMessage()

	// Close the writer pipe ignoring the error
	_ = writer.Close()

	// Read from the reader pipe end
	result, _ := io.ReadAll(reader)
	output := string(result)

	// return stdout to the default behaviour
	os.Stdout = stdOut

	// Compare for the real test
	if !strings.Contains(output, "this is a test!") {
		t.Errorf("expected something different")
	}
}
