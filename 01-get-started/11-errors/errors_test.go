package _1_errors

import (
	"log"
	"os"
	"testing"
)

func TestErrors(t *testing.T) {
	_, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	}
	// do something with the open *File f
}
