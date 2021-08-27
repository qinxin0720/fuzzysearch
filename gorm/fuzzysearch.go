package gorm

import (
	"fmt"
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
		condition = append(condition, fmt.Sprintf(`upper(%s) LIKE ? ESCAPE '|'`, key))
		for _, v := range moreKeys {
			condition = append(condition, fmt.Sprintf(`upper(%s) LIKE ? ESCAPE '|'`, v))
		}
		values := make([]interface{}, len(moreKeys)+1)
		for i := range values {
			values[i] = searchText
		}
		return db.Where(strings.Join(condition, " OR "), values...)
	}
}
