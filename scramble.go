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
var uppercaseRegexp = regexp.MustCompile("[\\p{Lu}]")
var letterRegexp = regexp.MustCompile("\\p{L}+")
var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// ShuffleRunes shuffles a slice of runes and returns a new slice.
func ShuffleRunes(r []rune) []rune {
	runes := make([]rune, 0, len(r))
	for _, n := range rand.Perm(len(r)) {
		runes = append(runes, r[n])
	}
	return runes
}

// RandomRunes returns a random slice of runes of the given length.
// The set of runes used is equivalent to the set of lowercase ASCII letters.
func RandomRunes(n int) []rune {
	runes := make([]rune, n)
	for i, _ := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return runes
}

// Scramble scrambles the words in a Unicode string.
// If random is false, scramble words in place (e.g., lorem > merlo).
// If random is true, replace words with random ASCII letters.
// Scramble preserves punctuation and the positions of uppercase letters.
func Scramble(s string, random bool) string {
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
		runeIndex := wordStart
		if random {
			// Replace characters with random ASCII letters.
			for _, r := range RandomRunes(len(word)) {
				runes[runeIndex] = r
				runeIndex++
			}
		} else {
			// Shuffle the word, then replace the word in the rune slice with
			// its shuffled version.
			for _, r := range ShuffleRunes(word) {
				runes[runeIndex] = r
				runeIndex++
			}
		}
	}
	// Restore uppercase characters.
	for _, i := range uppercasePositions {
		runes[i] = unicode.ToUpper(runes[i])
	}
	return string(runes)
}

// OpenFiles opens the named files for reading and returns a slice of file
// descriptors.
// If there is an error opening any file, OpenFiles will return an empty slice.
func OpenFiles(filenames ...string) ([]*os.File, error) {
	files := []*os.File{}
	for _, fn := range filenames {
		if fn == "-" {
			files = append(files, os.Stdin)
			continue
		}
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
	arguments, _ := docopt.Parse(usage, nil, true, "scramble 0.1", true)
	// Construct a slice of file handles.
	filenames := arguments["<file>"].([]string)
	files, e := OpenFiles(filenames...)
	if pe, ok := e.(*os.PathError); ok {
		fmt.Fprintf(os.Stderr, "scramble: %v: %v\n", pe.Path, pe.Err)
		os.Exit(1)
	}
	// Default to standard input.
	if len(files) == 0 {
		files = append(files, os.Stdin)
	}
	// Scramble each file line-by-line.
	random := arguments["--random"].(bool)
	for _, f := range files {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Println(Scramble(scanner.Text(), random))
		}
	}
}
