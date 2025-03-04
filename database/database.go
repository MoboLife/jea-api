package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// MySQL connection drive
	_ "github.com/go-sql-driver/mysql"

	// Postgres connection drive
	_ "github.com/lib/pq"

	"jea-api/environment"
)

// NewDatabase connect with database
func NewDatabase(info ConnectionInfo, schemaResolver bool) (*gorm.DB, error) {
	if schemaResolver {
		gorm.DefaultTableNameHandler = environment.TableNameHandler
	}
	if info.URL != "" {
		db, err := gorm.Open(info.Driver, info.URL)
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	db, err := gorm.Open(info.Driver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", info.Host, info.Port, info.User, info.Database, info.Password))
	if err != nil {
		return nil, err
	}
	return db, nil
}

// UseDatabase gin middleware for setup database connection
func UseDatabase(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}

// GetDatabase method for get database connection of gin context
func GetDatabase(ctx *gin.Context) *gorm.DB {
	db, exists := ctx.Get("db")
	database, ok := db.(*gorm.DB)
	if ok && exists {
		return database
	}
	return nil
}
