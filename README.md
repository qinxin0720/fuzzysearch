# Golang orm fuzzysearch

## install

```sh
go get -u -v github.com/qinxin0720/fuzzysearch
```

## type

```go
type Student struct {
    ID          string `gorm:"column:id" xorm:"'id'"`
    Name        string `gorm:"column:name" xorm:"'name'"`
    Description string `gorm:"column:description" xorm:"'description'"`
}
```

## for xorm

```go
import fuzzysearch "github.com/qinxin0720/fuzzysearch/xorm"

list := make([]Student, 0)

cond, values, err := fuzzySearch.FuzzySearch("alice", "name", "description")
if err != nil {
    return err
}

db.Where(cond, values...).
    Find(&list)
```

## for gorm

```go
import fuzzysearch "github.com/qinxin0720/fuzzysearch/gorm"

list := make([]Student, 0)

db.Scopes(fuzzySearch.FuzzySearch("alice", "name", "description")).
    Find(&list)
```
