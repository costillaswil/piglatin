package db

import (
	"pigLatin/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type PiglatinRepo struct {
	db *gorm.DB
}

func NewDB(databaseURI string) (*PiglatinRepo, error) {

	db, err := gorm.Open("sqlite3", databaseURI)

	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	db.LogMode(true)

	return &PiglatinRepo{
		db,
	}, nil
}

func (pr *PiglatinRepo) NewText(translationRes *model.TranslationResult) error {

	if ok := pr.db.HasTable(&model.TranslationResult{}); !ok {
		pr.db.AutoMigrate(&model.TranslationResult{})
	}

	return pr.db.Create(translationRes).Error
}
