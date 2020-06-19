package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Options base options for apply mutable data in database connection
type Options interface {
	Apply(db *gorm.DB) *gorm.DB
}

// DatabaseOptions database options builder
type DatabaseOptions struct {
	ApplyFunc func(db *gorm.DB) *gorm.DB
}

// Apply apply database customization
func (d *DatabaseOptions) Apply(db *gorm.DB) *gorm.DB {
	return d.ApplyFunc(db)
}

// WithPreloads add preload in database connection
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

// WithLimit apply limit for database query
func WithLimit(limit int) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) *gorm.DB {
		var database = db
		database = database.Limit(limit)
		return database
	}}
}

// WithWhere apply where for database query
func WithWhere(condition string, args ...interface{}) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) *gorm.DB {
		var database = db
		database = db.Where(condition, args)
		return database
	}}
}

// WithFilters apply filters in database query
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

// WithOffset apply offset in database query
func WithOffset(offset int) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) *gorm.DB {
		var database = db
		database = database.Offset(offset)
		return database
	}}
}

// WithOrder apply ordering in database query
func WithOrder(order string) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) *gorm.DB {
		var database = db
		database = database.Order(order)
		return database
	}}
}
