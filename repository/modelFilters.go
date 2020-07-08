package repository

import (
	"github.com/gin-gonic/gin"
	"jea-api/models"
)


func UseFilters(filters []models.ModelFilter) Filter {
	return &APIFilter{ApplyFunc: func(ctx *gin.Context) []Options {
		var finalFilters = make(map[string][]interface{})
		for _, filter := range filters {
			key, values := filter.UseGin(ctx)
			if key == ""{
				continue
			}
			finalFilters[key] = values
		}
		return []Options{WithFields(finalFilters)}
	}}
}

