package translator

import (
	"github.com/valyo95/gopher-translator/stringutil"
	"strings"
)

func TranslateWord(word string) string {
	// trim spaces
	word = strings.TrimSpace(word)

	if strings.Contains(word, "'") || stringutil.IsNumeric(word) {
		return word
	}

	if stringutil.StartsWithAVowel(word) {
		return "g" + word
	}

	if stringutil.StartsWithXR(word) {
		return "ge" + word
	}

	startsWithConsonantFollowedByQU, regex := stringutil.StartsWithConsonantFollowedByString(word, "qu")
	if startsWithConsonantFollowedByQU {
		return regex.ReplaceAllString(word, "${2}${1}ogo")
	}

	startsWithConsonant, regex := stringutil.StartsWithConsonant(word)
	if startsWithConsonant {
		return regex.ReplaceAllString(word, "${2}${1}ogo")
	}

	return word
}
