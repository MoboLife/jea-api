package database

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// Postgres connection drive
	_ "github.com/lib/pq"

	"jea-api/environment"

	log "github.com/sirupsen/logrus"
)

// ConnectionInfo database connection info
type ConnectionInfo struct {
	Host     string
	Port     string
	User     string
	Database string
	Password string
}

// NewDatabase connect with database
func NewDatabase(info ConnectionInfo) *gorm.DB {
	gorm.DefaultTableNameHandler = environment.TableNameHandler
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", info.Host, info.Port, info.User, info.Database, info.Password))
	if err != nil {
		log.Fatal("Error in connect with database. Error", err.Error())
	}
	return db
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
