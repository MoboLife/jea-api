package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type ModelFilterType string

func (m ModelFilterType) GetSeparator(multiple bool) string {
	switch m {
	case Date:
		return "%s BETWEEN ? AND ?"
	case String:
		return "%s LIKE %?%"
	case Integer:
		if multiple{
			return "%s IN (?)"
		}
		return "%s = ?"
	case Equal:
		return "%s = ?"
	}
	return ""
}

var (
	Date ModelFilterType = "DATE"
	String	ModelFilterType = "STRING"
	Integer ModelFilterType = "INTEGER"
	Equal	ModelFilterType = "EQUAL"
)
type ModelFilter struct {
	Query		string
	Field		string
	Multiple	bool
	Type		ModelFilterType
}

func (m *ModelFilter) UseGin(ctx *gin.Context) (string, []interface{}) {
	queryValue := ctx.Query(m.Query)
	if queryValue == ""{
		return "", nil
	}
	var values = strings.Split(queryValue, ",")
	var returnValue []interface{}
	if m.Multiple {
		for _, v := range values {
			returnValue = append(returnValue, v)
		}
	}else{
		returnValue = append(returnValue, values)
	}
	return fmt.Sprintf(m.Type.GetSeparator(len(values) > 1), m.Field), returnValue
}


type Filterable interface {
	GetFilters() Filters
}

// Model basic model
type Model struct {
	ID        		int64      		`json:"id" gorm:"primary_key"`
	CreatedAt 		time.Time  		`json:"createdAt"`
	UpdatedAt 		*time.Time 		`json:"updatedAt,omitempty"`
}
