package core

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const arrow = "-->"

var (
	cyan    = color.New(color.FgCyan)
	green   = color.New(color.FgGreen)
	red     = color.New(color.FgRed)
	blue    = color.New(color.FgBlue)
)

// WriteTaskHeader writes the task header message followed by a new line.
func WriteTaskHeader(taskName string) {
	cyan.Printf("bonclay: %s task\n\n", taskName)
}

// WriteTaskFooter writes the task footer message followed by a new line.
func WriteTaskFooter(taskName string, wasSuccessful bool) {
	if wasSuccessful {
		green.Printf("\n===> %s Successful\n\n", strings.Title(taskName))
	} else {
		red.Printf("\nSome errors occurred during %s:\n", taskName)
	}
}

// WriteTaskSuccess writes a success response for a src/dst pair, where src/dst
// is either a file or a directory.
func WriteTaskSuccess(src, dst string) {
	taskResponse(src, dst, true)
}

// WriteTaskFailure writes a failure response for a src/dst pair, where src/dst
// is either a file or a directory.
func WriteTaskFailure(src, dst string) {
	taskResponse(src, dst, false)
}

func taskResponse(src, dst string, wasSuccessful bool) {
	c := green
	if !wasSuccessful {
		c = red
	}

	blue.Printf("%s ", src)
	c.Printf("%s ", arrow)
	blue.Printf("%s\n", dst)
}

// WriteTaskErrors writes the errors, if any occurred.
//
// Duplicates are removed.
func WriteTaskErrors(errors []string) {
	if len(errors) == 0 {
		return
	}

	var uniqueErrors []string
	var errExists = make(map[string]bool)
	for _, v := range errors {
		if _, exists := errExists[v]; !exists {
			uniqueErrors = append(uniqueErrors, v)
			errExists[v] = true
		}
	}

	for i, v := range uniqueErrors {
		if i == (len(uniqueErrors) - 1) {
			fmt.Printf("\t%s\n\n", v)
		} else {
			fmt.Printf("\t%s\n", v)
		}
	}
}
