package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Function reads and returns first line from file
//Second argument 1 - temp, 2 - fan
// Returns `1` if file read error
// `2` if conversion issue
// `3` if wrong second argument
func readLine(fn string, tf uint8) (int) {
	f, err := os.Open(fn)
	if err != nil {
		return 1
	} else {
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	
	value, err := strconv.Atoi(line)
	if err != nil {
		return 2
	}

	switch tf {
		case 1:
			return value / 1000
		case 2:
			return value
		default:
			return 3	
		}
	}
}

func main() {
	s := readLine("ex.txt", 1)
	fmt.Printf("%d\n", s)
}
