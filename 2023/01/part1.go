package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	sum := 0
	for _, line := range text {
		chars := []string{}
		for _, char := range line {
			if unicode.IsDigit(char) {
				chars = append(chars, string(char))
			}
		}
		first_and_last := chars[0] + chars[len(chars)-1]
		n, _ := strconv.Atoi(first_and_last)
		sum += n
	}
	fmt.Println(sum)
}
