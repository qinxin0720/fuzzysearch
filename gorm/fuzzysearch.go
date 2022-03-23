package gorm

import (
	"strings"

	"gorm.io/gorm"
)

func FuzzySearch(keywords, key string, moreKeys ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		condition := make([]string, 0)
		keywords = strings.ReplaceAll(keywords, "|", "||")
		keywords = strings.ReplaceAll(keywords, "_", "|_")
		keywords = strings.ReplaceAll(keywords, "%", "|%")
		keywords = strings.ReplaceAll(keywords, "'", "|'")
		searchText := "%" + strings.ToUpper(keywords) + "%"
		condition = append(condition, "upper("+key+`) LIKE ? ESCAPE '|'`)
		for _, v := range moreKeys {
			condition = append(condition, "upper("+v+`) LIKE ? ESCAPE '|'`)
		}
		values := make([]any, len(moreKeys)+1)
		for i := range values {
			values[i] = searchText
		}
		return db.Where(strings.Join(condition, " OR "), values...)
	}
}
