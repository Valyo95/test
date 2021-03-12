package translator

import "testing"

func TestTranslateWord(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{"apple", "gapple"},
		{"xray", "gexray"},
		{"APple", "gAPple"},
		{"XRray", "geXRray"},
		{"chair", "airchogo"},
		{"CHair", "airCHogo"},
		{"mambo", "ambomogo"},
		{"flambe", "ambeflogo"},
		{"square", "aresquogo"},
		{"SQUare", "areSQUogo"},
		{"brmquare", "arebrmquogo"},
		{"flambe", "ambeflogo"},
		{"m", "mogo"},
		{"mquy", "y" + "mqu" + "ogo"},
		{"msdsdquare", "are" + "msdsdqu" + "ogo"},
		{"don't", "don't"},
		{"shouldn't", "shouldn't"},
		{"on-the-move", "gon-the-move"},
		{"well-defined", "ell-definedwogo"},
		{"insta360", "ginsta360"},
		{"grinsta360", "insta360grogo"},
		{"doing", "oingdogo"},
		{".", "."},
		{"?", "?"},
		{"!!", "!!"},
		{"doing?", "oing?dogo"},
		{"sure_", "ure_sogo"},
		{"1453665679", "1453665679"},
	}
	for _, tt := range tests {
		t.Run(tt.word, func(t *testing.T) {
			if got := TranslateWord(tt.word); got != tt.want {
				t.Errorf("TranslateWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
