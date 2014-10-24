package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
	"unicode"

	"github.com/docopt/docopt.go"
)

var usage = `Usage: scramble [options] [<file>...]

Options:
  -h, --help     Show this screen and exit.
  -r, --random   Replace words with random characters.
  -V, --version  Show the version and exit.

If <file> is not specified, or if <file> is -, read from standard input.`

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

// OpenFiles opens the named files for reading and returns a slice of file
// descriptors.
// If there is an error opening any file, OpenFiles will return an empty slice.
func OpenFiles(filenames ...string) ([]*os.File, error) {
	files := []*os.File{}
	// Default to standard input.
	if len(filenames) == 0 {
		files = append(files, os.Stdin)
		return files, nil
	}
	if filenames[0] == "-" {
		files = append(files, os.Stdin)
		return files, nil
	}
	for _, fn := range filenames {
		f, err := os.Open(fn)
		if err != nil {
			return []*os.File{}, err
		}
		files = append(files, f)
	}
	return files, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	arguments, err := docopt.Parse(usage, nil, true, "scramble 0.1", false)
	if err != nil {
		fmt.Println(err)
	}

	// Construct a slice of file handles.
	filenames := arguments["<file>"].([]string)
	files, err := OpenFiles(filenames...)
	if err != nil {
		fmt.Println(err)
	}

	// Choose a function to apply to each file.
	random := arguments["--random"].(bool)
	var action func(string) string
	if random {
		action = Randomize
	} else {
		action = Scramble
	}

	// Apply function to each file.
	for _, f := range files {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(action(scanner.Text()))
		}
	}
}
