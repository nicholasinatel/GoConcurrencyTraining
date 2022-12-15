package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/**
	** os.Pipe()
	* Used to create a synchronous, in-memory pipe.
	* Returns a pair of connected file descriptors, which can be used to communicate between 2 go routines
	* One goroutine can write data to the pipe,
	* Second goroutine can read data from the pipe.
	** Useful way to pass data between goroutines without using global variables or shared memory
	 */
	// 1st create the pipe and the file descriptors
	read, write, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	/**
	* Next, create a goroutine that will write data to the pipe.
	* The write file descriptor is passed as an argument to the goroutine.
	**/
	go func(w *os.File) {
		// Write some data to the pipe
		w.Write([]byte("Hello from the entrance of the pipe"))
		// Close the write end of the pipe to signal to the reader the end of the transmission
		w.Close()
	}(write)

	// Now, create a slice of byte cappable of receiving the information
	data := make([]byte, 100)
	// In the Main Go Routine:
	// Read the data which was written to the pipe by the other go routine
	n, err := read.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data[:n]))
	// Close the read end of the pipe
	read.Close()

}
