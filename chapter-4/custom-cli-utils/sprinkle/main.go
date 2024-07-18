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

	// transforms, err := getTransforms()

	// if err != nil {
	// 	fmt.Print(err)
	// 	return
	// }

	transforms := []string{
		"* app",
		"* site",
		"* time",
		"* hq",
		"get *",
		"go *",
		"lets * ",
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
	transformsDirty := strings.Split(string(data), "\n")
	transforms := []string{}

	for _, t := range transformsDirty {
		if t == "" {
			continue
		}
		transforms = append(transforms, t)
	}

	return transforms, nil
}
