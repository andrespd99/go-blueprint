package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var ErrFileNotFound = errors.New("error: file not found")

const otherWord = "*"

func main() {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	transforms, err := getTransforms()
	if err != nil {
		fmt.Print(err)
		return
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}

}

func getTransforms() ([]string, error) {
	data, err := os.ReadFile("transforms.txt")
	if err != nil {
		return []string{}, ErrFileNotFound
	}
	transforms := string(data)

	return strings.Split(transforms, "\n"), nil
}
