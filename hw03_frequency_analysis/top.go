package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

/* Safe initialization of global variables holding compiled regular expressions.
A regular expression is a sequence of characters that specifies a search pattern in text.
MustCompile panics if the expression cannot be parsed
The character class \s matches a space, tab, new line, carriage return or form feed, and + says one or more of those.*/
var re = regexp.MustCompile(`\s+`)

func Top10(inputText string) []string {
	result := make([]string, 0, 10)

	// If inputText is empty return the result
	if inputText == "" {
		return result
	}

	// Replace each whitespace substring with a single space character.
	inputText = re.ReplaceAllString(inputText, " ")
	// Split inputText into all substrings separated by " ". As result there is a slice of strings (in var top).
	top := strings.Split(inputText, " ")
	// Sort the slice of strings in increasing order.
	sort.Strings(top)

	// Create and fill a map with string in the key and int in the value
	rating := make(map[string]int, 300) // allocate memory for 300 keys in advance (we have 276 words in inputText)
	for _, val := range top {
		rating[val]++
	}

	// Create a sorted slice from map values.
	sortVal := make([]int, 0, 300)
	for _, val := range rating {
		sortVal = append(sortVal, val)
	}
	sort.Ints(sortVal)

	// Create a slice of top 10 values.
	rate10 := make([]int, 0, 10)
	for k := len(sortVal) - 1; k >= 0; k-- {
		if len(rate10) < 10 {
			rate10 = append(rate10, sortVal[k])
		}
	}

	// Create a slice of keys sorted lexicographically.
	keys := make([]string, 0, 300)
	for k := range rating {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Create the result slice of top 10 the most popular words in the input text.
	var prev int
	for _, val := range rate10 {
		if prev != val {
			for _, v := range keys {
				if rating[v] == val {
					result = append(result, v)
					if len(result) == 10 {
						break
					}
				}
				prev = val
			}
		} else {
			continue
		}
		if len(result) == 10 {
			break
		}
	}
	return result
}
