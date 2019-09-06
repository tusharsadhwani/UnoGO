package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func readInt(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	line := scanner.Text()
	line = strings.Trim(line, " \n")
	choice, err := strconv.Atoi(line)
	if err != nil {
		return -1
	}
	return choice
}
