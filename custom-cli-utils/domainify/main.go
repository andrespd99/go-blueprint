package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var defaultTlds = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"
const tldAllowedChars = "abcdefghijklmnopqrstuvwxyz"

func main() {
	defaultTldValues := strings.Join(defaultTlds, ",")
	var tldsInput = flag.String("tlds", defaultTldValues, "Sets the list of TLDs to be used on the domain generation result")
	flag.Parse()

	tlds := parseTlds(*tldsInput)

	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." +
			tlds[rand.Intn(len(tlds))])
	}
}

func parseTlds(input string) []string {
	input = strings.ToLower(input)
	var inputClean []rune
	for _, r := range input {
		if !strings.ContainsRune(tldAllowedChars+",", r) {
			continue
		}
		inputClean = append(inputClean, r)
	}

	return strings.Split(string(inputClean), ",")
}
