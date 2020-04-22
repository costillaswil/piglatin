package service

import (
	"pigLatin/db"
	"pigLatin/model"
	"regexp"
	"strings"
	"unicode"
)

type PiglatinService struct {
	repository *db.PiglatinRepo
}

func NewPiglatinService(repository *db.PiglatinRepo) *PiglatinService {
	return &PiglatinService{
		repository: repository,
	}
}

/*
	takes word or group of words.

	if group of words is supplied explode the string via white spaces,
	after exploding it checks if it contains special characters.

	words with special characters are not translated and automatically concatenated
	to the output string.

	check for panctuations, if a word contains multiple punctuations, collect word to
	translate till the next punctuation is raised.

	returns a struct contains the original text and translated text

*/

//Format text translates given text to a Piglatin translated format.
func (p *PiglatinService) FormatText(originalText string) (model.TranslationResult, error) {

	var (
		punctuations = "!?.,-"
		transText    []string
	)

	for _, word := range strings.Split(originalText, " ") {

		toTranslate := ""
		if ok, _ := hasSpecialChar(word); !ok {
			for key, val := range word {

				if !strings.Contains(punctuations, string(val)) {
					toTranslate += string(val)
				} else {
					if toTranslate != "" {
						transText = append(transText, translate(toTranslate))
					}
					transText = append(transText, string(val))
					toTranslate = ""
				}

				if key == len(word)-1 && !strings.Contains(punctuations, string(val)) {
					transText = append(transText, translate(toTranslate))
				}

			}
		} else {
			transText = append(transText, word)
		}

	}

	translationResult := model.TranslationResult{
		OriginalText:   originalText,
		TranslatedText: strings.Join(transText, " "),
	}

	err := p.repository.NewText(&translationResult)

	if err != nil {
		return translationResult, err
	}

	return translationResult, nil
}

/*
	check if initial character of the word string is vowel.
	if the character is vowel concatenate word with vowel suffix
	otherwise get get the index of the char before the next vowel.

	If the inital char of the word to be translated is capital, convert
	output to capitalized.

*/
func translate(word string) (returnSTR string) {

	var (
		vowel       = "aeiou"
		suffix      = "ay"
		vowelSuffix = "way"
		counter     = 0
		wordRune    = []rune(word)
		first       = word[0:1]
	)

	if strings.Contains(vowel, strings.ToLower(first)) {
		returnSTR = word + vowelSuffix
	} else {

		for !strings.Contains(vowel, string(word[counter])) {
			counter++
			if counter > len(word)-1 {
				counter = 0
				break
			}
		}

		returnSTR = word[counter:] + word[0:counter] + suffix
	}

	if unicode.IsUpper(rune(wordRune[0])) {
		return strings.Title(returnSTR)
	}

	return returnSTR
}

/*
	check if the word supplied contains special character
	return true if the condition is satisfied else false.

*/

func hasSpecialChar(word string) (bool, error) {
	return regexp.MatchString(`[^a-zA-Z'!?,. -]`, word)
}
