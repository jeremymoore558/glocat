package main

/*NOTES:
260413:
	Got file io working. Need to add argn argc type handling into main, and
	a way to check if the file exists already.

	Got a way to accept a file name from the command line, now need to verify file exists,
	and if not, just print the arguments as the string.

260412:
	Got proper rainbow set up, over integers,
	and construction of the right ansi escape sequences.

	Still have to implement reading a piece of text from stdin.
	Idea of the logic:
		glocat fkdslf.file
		First, check if there is a file in the current directory by that name.
		If not, just output the input text through the glocat processing.
*/

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var test_str string
	// If multiple args recieved, process those. If not, read from std in.
	if len(os.Args) >= 2 {
		// Check if file exists. If not, cat all args to one string.
		if _, err := os.Stat(os.Args[1]); err == nil {
			input := os.Args[1]
			test_str = read_file(input)
		} else {
			test_str = concatenate_args()
		}
	} else {
		test_str = get_std_in()
	}
	print_colored(test_str)
	fmt.Printf("\x1b[0m")
}

func isPiped() bool {
	info, _ := os.Stdin.Stat()
	return (info.Mode() & os.ModeCharDevice) == 0
}

func get_std_in() string {
	if !isPiped() {
		return ""
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func concatenate_args() string {
	var args string
	args = ""
	for i := 0; i < len(os.Args); i++ {
		args += os.Args[i] + " "
	}
	return args
}

func read_file(filename string) string {
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	buf := make([]byte, 1024)
	for {
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
	}
	return string(buf)
}
