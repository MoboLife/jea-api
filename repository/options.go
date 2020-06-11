package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Options interface {
	Apply(db *gorm.DB) *gorm.DB
}

type DatabaseOptions struct {
	ApplyFunc func(db *gorm.DB) *gorm.DB
}

func (d *DatabaseOptions) Apply(db *gorm.DB) *gorm.DB {
	return d.ApplyFunc(db)
}

func WithPreloads(preloads ...string) Options {
	return &DatabaseOptions{
		ApplyFunc: func(db *gorm.DB) *gorm.DB {
			var database = db
			for _, preload := range preloads {
				database = database.Preload(preload)
			}
			return database
		},
	}
}

func WithLimit(limit int) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) *gorm.DB {
		var database = db
		database = database.Limit(limit)
		return database
	}}
}

func WithWhere(condition string, args ...interface{}) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) *gorm.DB {
		var database = db
		database = db.Where(condition, args)
		return database
	}}
}

func WithFilters(ctx *gin.Context, filters ...Filter) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) *gorm.DB {
		var database = db
		var options []Options
		for _, filter := range filters {
			options = append(options, filter.Apply(ctx)...)
		}
		for _, option := range options {
			database = option.Apply(database)
		}
		return database
	}}
}


func WithOffset(offset int) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) *gorm.DB {
		var database = db
		database = database.Offset(offset)
		return database
	}}
}

func WithOrder(order string) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) *gorm.DB {
		var database = db
		database = database.Order(order)
		return database
	}}
}