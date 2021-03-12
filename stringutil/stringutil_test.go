package stringutil

import (
	"github.com/valyo95/gopher-translator/domain"
	"reflect"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"1", true},
		{"-10", true},
		{"1.25", true},
		{"1142341623416231311231231231", true},
		{"trombone", false},
		{"1.25.", false},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := IsNumeric(tt.str); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartsWithAVowel(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"age", true},
		{"even", true},
		{"interact", true},
		{"opportunity", true},
		{"understand", true},
		{"Age", true},
		{"Even", true},
		{"Interact", true},
		{"Opportunity", true},
		{"Understand", true},
		{"track", false},
		{"modem", false},
		{"Track", false},
		{"Modem", false},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			got := StartsWithAVowel(tt.word)
			if got != tt.want {
				t.Errorf("StartsWithAVowel() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_startsWithXR(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"xray", true},
		{"XRander", true},
		{"apple", false},
		{"test", false},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := StartsWithXR(tt.word); got != tt.want {
				t.Errorf("startsWithXR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_startsWithConsonant(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"abc", false},
		{"e", false},
		{"m", true},
		{"Apple", false},
		{"ENvelope", false},
		{"chair", true},
		{"people", true},
		{"trombone", true},
		{"schedule", true},
		{"SChedule", true},
		{".", false},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got, _ := StartsWithConsonant(tt.word); got != tt.want {
				t.Errorf("startsWithConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_startsWithConsonantFollowedByString(t *testing.T) {
	tests := []struct {
		word string
		str  string
		want bool
	}{
		{"abc", "", false},
		{"e", "qu", false},
		{"Apple", "", false},
		{"ENvelope", "", false},
		{"chair", "", true},
		{"people", "eo", true},
		{"trombone", "ombone", true},
		{"schedule", "che", true},
		{"SChedule", "", true},
		{"aqu", "", false},
		{"aeiqu", "eiqu", false},
		{"frequency", "quency", false},
		{"square", "qu", true},
		{"bcdfqu", "fqu", true},
		{"SQUARE", "QU", true},
		{"SQUARE", "U", true},
		{"SQUARE", "ARE", false},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got, _ := StartsWithConsonantFollowedByString(tt.word, tt.str); got != tt.want {
				t.Errorf("startsWithConsonant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SplitSentenceIntoWords(t *testing.T) {
	tests := []struct {
		sentence domain.SentenceRequest
		want     []string
	}{
		{domain.SentenceRequest{EnglishSentence: "Hi, how are you doing?"}, []string{"Hi", ",", "how", "are", "you", "doing", "?"}},
		{domain.SentenceRequest{EnglishSentence: "Let's eat, Grandma!"}, []string{"Let's", "eat", ",", "Grandma", "!"}},
		{domain.SentenceRequest{EnglishSentence: "I'm sorry; I love you"}, []string{"I'm", "sorry", ";", "I", "love", "you"}},
	}
	for _, tt := range tests {
		t.Run(tt.sentence.EnglishSentence, func(t *testing.T) {
			if got := SplitSentenceIntoWords(tt.sentence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitSentenceIntoWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimSuffix(t *testing.T) {
	type args struct {
		s      string
		suffix string
	}
	tests := []struct {
		args args
		want string
	}{
		{args{s: "test\n", suffix: "\n"}, "test"},
		{args{s: "tralala", suffix: "la"}, "trala"},
	}
	for _, tt := range tests {
		t.Run(tt.args.s, func(t *testing.T) {
			if got := TrimSuffix(tt.args.s, tt.args.suffix); got != tt.want {
				t.Errorf("TrimSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
