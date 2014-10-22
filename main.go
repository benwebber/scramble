package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
)

var lowercaseRegexp = regexp.MustCompile("[\\p{Ll}]")
var uppercaseRegexp = regexp.MustCompile("[\\p{Lu}]")
var uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVQXYZ"
var lowercaseChars = "abcdefghijklmnopqrstuvqxyz"

// Replace upper- or lowercase letters with random characters of the same case.
func ScrambleCharacters(text string) (ret string) {
	runes := []rune(text)
	for i, c := range text {
		if uppercaseRegexp.MatchString(string(c)) {
			runes[i] = rune(uppercaseChars[rand.Intn(len(uppercaseChars))])
		} else if lowercaseRegexp.MatchString(string(c)) {
			runes[i] = rune(lowercaseChars[rand.Intn(len(lowercaseChars))])
		}
	}
	ret = string(runes)
	return
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(ScrambleCharacters(scanner.Text()))
	}
}
