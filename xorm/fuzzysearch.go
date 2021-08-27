package xorm

import (
	"fmt"
	"strings"

	"xorm.io/builder"
)

func FuzzySearch(keywords, key string, moreKeys ...string) (string, []interface{}, error) {
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
	return builder.ToSQL(builder.Expr(strings.Join(condition, " OR "), values...))
}
