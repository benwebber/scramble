package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
	"unicode"
)

var lowercaseRegexp = regexp.MustCompile("[\\p{Ll}]")
var uppercaseRegexp = regexp.MustCompile("[\\p{Lu}]")
var letterRegexp = regexp.MustCompile("\\p{L}+")
var uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVQXYZ"
var lowercaseChars = "abcdefghijklmnopqrstuvqxyz"

// Shuffle a slice of runes.
func ShuffleRunes(r []rune) (ret []rune) {
	ret = make([]rune, 0, len(r))
	for _, n := range rand.Perm(len(r)) {
		ret = append(ret, r[n])
	}
	return
}

// Replace upper- or lowercase letters with random characters of the same case.
func Randomize(text string) (ret string) {
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

func Scramble(s string) (ret string) {
	// Use FindAllStringIndex() to locate and preserve the positions of
	// uppercase letters.
	uppercasePositions := []int{}
	for _, i := range uppercaseRegexp.FindAllStringIndex(s, -1) {
		uppercasePositions = append(uppercasePositions, i[0])
	}

	// Convert the string to lowercase runes for processing.
	runes := []rune(s)
	for i, _ := range runes {
		runes[i] = unicode.ToLower(runes[i])
	}

	// Use FindAllStringIndex() to establish word boundaries by looking for
	// continuous sequences of letters.
	for _, i := range letterRegexp.FindAllStringIndex(s, -1) {
		wordStart := i[0]
		wordEnd := i[1]
		word := runes[wordStart:wordEnd]
		// ShuffleRunes the word, then replace the word in the rune slice with
		// its shuffled version.
		runeIndex := wordStart
		for _, r := range ShuffleRunes(word) {
			runes[runeIndex] = r
			runeIndex++
		}
	}

	// Restore uppercase characters.
	for _, i := range uppercasePositions {
		runes[i] = unicode.ToUpper(runes[i])
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
		fmt.Println(Scramble(scanner.Text()))
	}
}
