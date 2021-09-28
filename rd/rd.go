package rd

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var s *bufio.Scanner

func init() {
	s = bufio.NewScanner(os.Stdin)
}

func String() string {
	if !s.Scan() {
		log.Print("scan failed")
		os.Exit(1)
	}
	return s.Text()
}

func Int() int {
	i, err := strconv.Atoi(String())
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	return i
}
