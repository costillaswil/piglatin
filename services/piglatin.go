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
func (p *PiglatinService) FormatText(originalText string) (map[string]interface{}, error) {

	var (
		vowelMap = map[string]bool{
			"a": true,
			"e": true,
			"i": true,
			"o": true,
			"u": true,
		}
		transText = ""
		altText   = ""
		initVowel = false
	)

	textArr := strings.Split(originalText, "")

	for key, val := range textArr {

		if _, ok := vowelMap[val]; ok && !initVowel {

			transText = arrangeResult(key, len(textArr), textArr)

			if _, ok := vowelMap[textArr[0]]; !ok {
				break
			}

			initVowel = true
		}

		if _, ok := vowelMap[val]; !ok && initVowel && key < len(textArr)-1 {
			if _, ok := vowelMap[textArr[key+1]]; ok {
				altText = arrangeResult(key+1, len(textArr), textArr)
				break
			}
		}
	}

	translationResult := model.TranslationResult{
		OriginalText:   originalText,
		TranslatedText: transText,
		AltTranslation: altText,
	}

	err := p.repository.NewText(&translationResult)

	if err != nil {
		return map[string]interface{}{}, err
	}

	return map[string]interface{}{
		"original_text":   originalText,
		"translated_text": transText,
		"alt_translation": altText,
	}, nil
}

func arrangeResult(vowelKey, textLength int, textSlice []string) string {
	const suffix = "ay"
	newTextArr := append(textSlice[vowelKey:textLength], textSlice[0:vowelKey]...)
	return strings.Join(newTextArr, "") + suffix
}
