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
	s.Split(bufio.ScanWords)
}

func String() string {
	if !s.Scan() {
		log.Print("scan failed")
		os.Exit(1)
	}
	return s.Text()
}

func String2() (string, string) {
	return String(), String()
}

func String3() (string, string, string) {
	return String(), String(), String()
}

func Strings(n int) []string {
	s := make([]string, n)
	for i := range s {
		s[i] = String()
	}
	return s
}

func Int() int {
	i, err := strconv.Atoi(String())
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	return i
}

func Int2() (int, int) {
	return Int(), Int()
}

func Int3() (int, int, int) {
	return Int(), Int(), Int()
}

func Ints(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = Int()
	}
	return s
}
