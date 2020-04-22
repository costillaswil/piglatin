package service

import (
	"pigLatin/db"
	"pigLatin/model"
	"strings"
)

type PiglatinService struct {
	repository *db.PiglatinRepo
}

func NewPiglatinService(repository *db.PiglatinRepo) *PiglatinService {
	return &PiglatinService{
		repository: repository,
	}
}

//Format text translates given text to a Piglatin translated format.
func (p *PiglatinService) FormatText(originalText string) (model.TranslationResult, error) {

	var (
		vowelMap = map[string]bool{
			"a": true,
			"e": true,
			"i": true,
			"o": true,
			"u": true,
		}
		transText []string
		altText   []string
		initVowel = false
	)

	for wordCounter, word := range strings.Split(originalText, " ") {

		textArr := strings.Split(word, "")
		initVowel = false
		transText = append(transText, word)
		altText = append(altText, word)

		for key, char := range textArr {

			if _, ok := vowelMap[char]; ok && !initVowel {

				transText[wordCounter] = arrangeResult(key, len(textArr), textArr)
				altText[wordCounter] = arrangeResult(key, len(textArr), textArr)
				if _, ok := vowelMap[textArr[0]]; !ok || len(textArr) == 1 {
					break
				}

				initVowel = true
			}

			if _, ok := vowelMap[char]; !ok && initVowel && key < len(textArr)-1 {

				if len(textArr)-2 == key {
					break
				}

				if _, ok := vowelMap[textArr[key+1]]; ok {
					altText[wordCounter] = arrangeResult(key+1, len(textArr), textArr)
					break
				}
			}

		}
	}

	translationResult := model.TranslationResult{
		OriginalText:   originalText,
		TranslatedText: strings.Join(transText, " "),
		AltTranslation: strings.Join(altText, " "),
	}

	err := p.repository.NewText(&translationResult)

	if err != nil {
		return translationResult, err
	}

	return translationResult, nil
}

func arrangeResult(vowelKey, textLength int, textSlice []string) string {
	const suffix = "@ay"

	if textLength > 1 {
		newTextArr := append(textSlice[vowelKey:textLength], textSlice[0:vowelKey]...)
		return strings.Replace(suffix, "@", strings.Join(newTextArr, ""), -1)
	}

	return strings.Replace(suffix, "@", textSlice[0], -1)
}
