package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"jea-api/database"
	"reflect"
)

// Repository interface for search, create, update and delete data of models in database
type Repository interface {
	FindAll(options ...Options) (interface{}, error)
	Find(id int64, options ...Options) (interface{}, error)
	Create(entity interface{}, options ...Options) error
	Delete(id int64, options ...Options) error
	Update(entity interface{}, id int64, options ...Options) error
	Total([]Options) (int64, error)
}

// RepositoryContext context for Repository
type RepositoryContext struct {
	DB        *gorm.DB
	ModelType reflect.Type
	Model     interface{}
}

// FindAll findall items in database
func (r *RepositoryContext) FindAll(options ...Options) (interface{}, error) {
	var db = r.applyOptions(options)
	var items = reflect.New(reflect.SliceOf(r.ModelType)).Interface()
	err := db.Find(items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

// Total get total items of search
func (r *RepositoryContext) Total(modelFiltersOption []Options) (int64, error) {
	var db = r.applyOptions(modelFiltersOption)
	var total int64
	err := db.Model(r.Model).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil

}

// Find find specify registry
func (r *RepositoryContext) Find(id int64, options ...Options) (interface{}, error) {
	var db = r.applyOptions(options)
	var item = reflect.New(r.ModelType).Interface()
	err := db.First(item, id).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Create create item in database
func (r *RepositoryContext) Create(entity interface{}, options ...Options) error {
	var db = r.applyOptions(options)
	err := db.Create(entity).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete delete item in database
func (r *RepositoryContext) Delete(id int64, options ...Options) error {
	var db = r.applyOptions(options)
	err := db.Delete(r.Model, id).Error
	if err != nil {
		return err
	}
	return nil
}

// Update item in database
func (r *RepositoryContext) Update(entity interface{}, id int64, options ...Options) error {
	var db = r.applyOptions(options)
	err := db.Where("id = ?", id).Updates(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryContext) applyOptions(options []Options) *gorm.DB {
	var database = r.DB
	for _, option := range options {
		database = option.Apply(database)
	}
	return database
}

// NewRepository create repository
func NewRepository(model interface{}, db *gorm.DB) Repository {
	var modelType = reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	return &RepositoryContext{DB: db.Model(model), Model: model, ModelType: modelType}
}

// UseRepository middleware for setup repository in gin routers
func UseRepository(model interface{}, repository *Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		if repository == nil {
			*repository = NewRepository(model, database.GetDatabase(c))
		}
	}
}
