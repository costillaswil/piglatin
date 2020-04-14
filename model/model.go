package model

type TranslationResult struct {
	ID             int    `gorm:"column:id" json:"id"`
	OriginalText   string `gorm:"column:original_text" json:"original_text"`
	TranslatedText string `gorm:"column:trasnlated_text" json:"translated_text"`
	AltTranslation string `gorm:"column:alt_translation" json:"alt_translation"`
}
