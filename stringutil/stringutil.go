package stringutil

import (
	"github.com/valyo95/gopher-translator/domain"
	"regexp"
	"strconv"
	"strings"
)

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func StartsWithAVowel(word string) bool {
	matchString, _ := regexp.MatchString("(?i)^[aeiou]", word)
	return matchString
}

func StartsWithXR(word string) bool {
	matchString, _ := regexp.MatchString("(?i)^xr*", word)
	return matchString
}

/*
	Every word starting with a consonant is followed by the empty ("")
	string, cause every string is is followed by an empty string
	since (str == str + "").
*/
func StartsWithConsonant(word string) (bool, *regexp.Regexp) {
	return StartsWithConsonantFollowedByString(word, "")
}

/*
	This method is more generic than the simple startsWithConsonantFollowedByQU
	and can be extended in the future to get any word that starts with a consonant
	and is followed by a given string.
*/
func StartsWithConsonantFollowedByString(word string, str string) (bool, *regexp.Regexp) {
	re := regexp.MustCompile(`(?i)(^[^aeiou\W]+` + str + ")(.*)")
	return re.MatchString(word), re
}

func SplitSentenceIntoWords(sentenceRequest domain.SentenceRequest) []string {
	sentenceRequest.EnglishSentence = regexp.MustCompile(`(\w+)([,.!?;:])`).ReplaceAllString(sentenceRequest.EnglishSentence, "${1} ${2}")
	words := strings.Fields(sentenceRequest.EnglishSentence)
	return words
}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
