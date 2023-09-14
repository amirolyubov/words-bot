package messages

import (
	"fmt"
	"strings"
	"words-bot/dictionary"
)

func CardCaption(word dictionary.Word) string {
	head := fmt.Sprintf(`\[%s]

`, word.Transcription)
	meaning := ""
	for i, mean := range word.Meaning {
		meaning = meaning + fmt.Sprintf(`%s. *%s.* %s
`,
			fmt.Sprint(i+1), mean.PartOfSpeech, mean.Explanation) + fmt.Sprintf(`_%#v_
`, mean.Example)
	}

	translate := fmt.Sprintf("\n_translation_\nru: %s\nfr: %s\n\n", word.Translations.Ru, word.Translations.Fr)

	synonyms := ""
	if len(word.Synonyms) > 0 {
		synonyms = fmt.Sprintf(`_synonyms_
%v`,
			strings.Join(word.Synonyms, ", "))
	}

	message := head + meaning + translate + synonyms
	return message
}

func QuizCaption(word dictionary.Word) string {
	head := fmt.Sprintf(`\[%s]

`, word.Transcription)
	meaning := ""
	for i, mean := range word.Meaning {
		meaning = meaning + fmt.Sprintf(`%s. *%s.* %s
`,
			fmt.Sprint(i+1), mean.PartOfSpeech, mean.Explanation) + fmt.Sprintf(`_%#v_
`, mean.Example)
	}

	synonyms := ""
	if len(word.Synonyms) > 0 {
		synonyms = fmt.Sprintf(`
_synonyms_
%v`,
			strings.Join(word.Synonyms, ", "))
	}

	message := head + meaning + synonyms
	return message
}
