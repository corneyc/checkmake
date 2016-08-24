// Package parser implements all the parser functionality for Makefiles
// this specific file holds the functionality for the scanner
package parser

import (
	"bufio"
	"fmt"
	"os"
)

// MakefileScanner is a wrapping struct around bufio.Scanner which provides
// extra functionality like the current line number
type MakefileScanner struct {
	Scanner    *bufio.Scanner
	LineNumber int
	FileHandle *os.File
}

// Scan is a thin wrapper around the bufio.Scanner Scan() function
func (s *MakefileScanner) Scan() bool {
	s.LineNumber++
	return s.Scanner.Scan()
}

// Close closes all open handles the scanner has
func (s *MakefileScanner) Close() {
	s.FileHandle.Close()
}

// Text is a thin wrapper around bufio.Scanner Text()
func (s *MakefileScanner) Text() string {
	return s.Scanner.Text()
}

// NewMakefileScanner returns a MakefileScanner struct for parsing a Makefile
func NewMakefileScanner(filepath string) (*MakefileScanner, error) {
	ret := &MakefileScanner{}
	var fileOpenErr error
	ret.FileHandle, fileOpenErr = os.Open(filepath)
	if fileOpenErr != nil {
		return ret, fmt.Errorf("Error opening the provided filepath '%s'", filepath)
	}
	ret.Scanner = bufio.NewScanner(ret.FileHandle)
	ret.Scanner.Split(bufio.ScanLines)

	return ret, nil
}